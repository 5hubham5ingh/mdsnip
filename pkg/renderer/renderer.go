package renderer

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	extast "github.com/yuin/goldmark/extension/ast"
	"github.com/yuin/goldmark/text"
	"golang.org/x/term"
)

// ANSI codes (8-bit black and white only)
const (
	ansiReset     = "\x1b[0m"
	ansiBold      = "\x1b[1m"
	ansiBoldOff   = "\x1b[22m"
	ansiItalic    = "\x1b[3m"
	ansiItOff     = "\x1b[23m"
	ansiUL        = "\x1b[4m"
	ansiULOff     = "\x1b[24m"
	ansiStrike    = "\x1b[9m"
	ansiStrikeOff = "\x1b[29m"
	ansiDim       = "\x1b[2m"
	ansiDimOff    = "\x1b[22m"
	ansiCodeBg    = "\x1b[47;30m" // black text, white bg
)

// Box drawing
const (
	bTL = "\u256d"
	bTR = "\u256e"
	bBL = "\u2570"
	bBR = "\u256f"
	bH  = "\u2500"
	bV  = "\u2502"
	bTT = "\u252c"
	bBT = "\u2534"
	bLT = "\u251c"
	bRT = "\u2524"
	bX  = "\u253c"
)

const qBar = "\u258c" // ▌

var hSizes = [7]int{0, 7, 6, 5, 4, 3, 2}

type renderer struct {
	src   []byte
	width int
	qd    int // blockquote depth
	ld    int // list depth
}

// Render parses markdown and prints styled output to the terminal.
func Render(content []byte) {
	if len(content) == 0 {
		return
	}
	w, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil || w <= 0 {
		w = 80
	}
	md := goldmark.New(goldmark.WithExtensions(extension.Table, extension.Strikethrough))
	doc := md.Parser().Parse(text.NewReader(content))
	r := &renderer{src: content, width: w}
	out := strings.TrimRight(r.block(doc), "\n") + "\n"
	fmt.Print(out)
}

// Error reports render errors to stderr.
func Error(err error) {
	fmt.Fprintf(os.Stderr, "Render Error: %v\n", err)
}

// ── Block dispatcher ──

func (r *renderer) block(n ast.Node) string {
	switch n.Kind() {
	case ast.KindDocument:
		return r.children(n)
	case ast.KindHeading:
		return r.heading(n.(*ast.Heading))
	case ast.KindParagraph:
		return r.para(n.(*ast.Paragraph))
	case ast.KindTextBlock:
		return r.inlineAll(n) + "\n"
	case ast.KindFencedCodeBlock, ast.KindCodeBlock:
		return r.codeBlock(n)
	case ast.KindBlockquote:
		return r.blockquote(n.(*ast.Blockquote))
	case ast.KindList:
		return r.list(n.(*ast.List))
	case ast.KindListItem:
		return ""
	case ast.KindThematicBreak:
		return strings.Repeat(bH, r.avail()) + "\n\n"
	case ast.KindHTMLBlock:
		return ""
	}
	if n.Kind() == extast.KindTable {
		return r.table(n.(*extast.Table))
	}
	if n.Type() == ast.TypeBlock {
		return r.children(n)
	}
	return ""
}

func (r *renderer) children(n ast.Node) string {
	var b strings.Builder
	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		b.WriteString(r.block(c))
	}
	return b.String()
}

// ── Inline dispatcher ──

func (r *renderer) inline(n ast.Node) string {
	switch n.Kind() {
	case ast.KindText:
		t := n.(*ast.Text)
		s := string(t.Segment.Value(r.src))
		if t.HardLineBreak() {
			s += "\n"
		} else if t.SoftLineBreak() {
			s += " "
		}
		return s
	case ast.KindString:
		return string(n.(*ast.String).Value)
	case ast.KindEmphasis:
		e := n.(*ast.Emphasis)
		c := r.inlineAll(n)
		if e.Level == 2 {
			return ansiBold + c + ansiBoldOff
		}
		return ansiItalic + c + ansiItOff
	case ast.KindCodeSpan:
		return ansiCodeBg + " " + r.plainText(n) + " " + ansiReset
	case ast.KindLink:
		l := n.(*ast.Link)
		dest := string(l.Destination)
		content := r.inlineAll(n)
		// OSC 8: \x1b]8;;URL\x1b\TEXT\x1b]8;;\x1b\
		return "\x1b]8;;" + dest + "\x1b\\" + ansiUL + content + ansiULOff + "\x1b]8;;\x1b\\" + ansiDim + " (" + dest + ")" + ansiDimOff
	case ast.KindAutoLink:
		u := string(n.(*ast.AutoLink).URL(r.src))
		return "\x1b]8;;" + u + "\x1b\\" + ansiUL + u + ansiULOff + "\x1b]8;;\x1b\\"
	case ast.KindImage:
		return "[" + r.plainText(n) + "](" + string(n.(*ast.Image).Destination) + ")"
	case ast.KindRawHTML:
		return ""
	}
	if n.Kind() == extast.KindStrikethrough {
		return ansiStrike + r.inlineAll(n) + ansiStrikeOff
	}
	return r.inlineAll(n)
}

func (r *renderer) inlineAll(n ast.Node) string {
	var b strings.Builder
	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		b.WriteString(r.inline(c))
	}
	return b.String()
}

// ── Block renderers ──

func (r *renderer) heading(n *ast.Heading) string {
	t := r.plainText(n)

	if isKitty() {
		sz := 1
		if n.Level >= 1 && n.Level <= 6 {
			sz = hSizes[n.Level]
		}

		return fmt.Sprintf("\x1b]66;s=%d;%s\x07%s", sz, t, strings.Repeat("\n", sz))
	}

	// Fallback for non-kitty terminals
	level := n.Level
	if level < 1 {
		level = 1
	}
	if level > 6 {
		level = 6
	}
	prefix := strings.Repeat("#", level)
	return ansiBold + prefix + " " + t + ansiBoldOff + "\n\n"
}

func isKitty() bool {
	return os.Getenv("TERM") == "xterm-kitty" || os.Getenv("KITTY_WINDOW_ID") != ""
}

func (r *renderer) para(n *ast.Paragraph) string {
	return wrapText(r.inlineAll(n), r.avail()) + "\n\n"
}

func (r *renderer) codeBlock(n ast.Node) string {
	var lines []string
	l := n.Lines()
	for i := 0; i < l.Len(); i++ {
		seg := l.At(i)
		s := string(seg.Value(r.src))
		lines = append(lines, strings.TrimRight(s, "\n\r"))
	}
	// Remove trailing empty lines
	for len(lines) > 0 && strings.TrimSpace(lines[len(lines)-1]) == "" {
		lines = lines[:len(lines)-1]
	}
	if len(lines) == 0 {
		return ""
	}
	return renderBox(lines, r.avail()) + "\n"
}

func (r *renderer) blockquote(n *ast.Blockquote) string {
	r.qd++
	content := strings.TrimRight(r.children(n), "\n")
	r.qd--
	if content == "" {
		return ""
	}
	var b strings.Builder
	for _, line := range strings.Split(content, "\n") {
		b.WriteString(qBar + " " + line + "\n")
	}
	return b.String() + "\n"
}

func (r *renderer) list(n *ast.List) string {
	r.ld++
	var b strings.Builder
	idx := 1
	if n.Start > 0 {
		idx = n.Start
	}
	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		if item, ok := c.(*ast.ListItem); ok {
			b.WriteString(r.listItem(item, n.IsOrdered(), idx))
			idx++
		}
	}
	r.ld--
	if r.ld == 0 {
		b.WriteString("\n")
	}
	return b.String()
}

func (r *renderer) listItem(n *ast.ListItem, ordered bool, idx int) string {
	marker := "• "
	if ordered {
		marker = fmt.Sprintf("%d. ", idx)
	}
	mw := visibleLen(marker)

	var content strings.Builder
	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		if c.Kind() == ast.KindParagraph {
			content.WriteString(wrapText(r.inlineAll(c), r.avail()-mw))
		} else {
			content.WriteString(r.block(c))
		}
	}

	lines := strings.Split(strings.TrimRight(content.String(), "\n"), "\n")
	indent := strings.Repeat(" ", mw)
	var b strings.Builder
	for i, line := range lines {
		if i == 0 {
			b.WriteString(marker + line + "\n")
		} else {
			b.WriteString(indent + line + "\n")
		}
	}
	return b.String()
}

func (r *renderer) table(n *extast.Table) string {
	var rows [][]string
	for row := n.FirstChild(); row != nil; row = row.NextSibling() {
		var cells []string
		for cell := row.FirstChild(); cell != nil; cell = cell.NextSibling() {
			cells = append(cells, r.inlineAll(cell))
		}
		rows = append(rows, cells)
	}
	if len(rows) == 0 {
		return ""
	}

	// Column count & widths
	nc := 0
	for _, row := range rows {
		if len(row) > nc {
			nc = len(row)
		}
	}
	cw := make([]int, nc)
	for _, row := range rows {
		for i, c := range row {
			if w := visibleLen(c); w > cw[i] {
				cw[i] = w
			}
		}
	}
	for i := range cw {
		if cw[i] < 3 {
			cw[i] = 3
		}
	}

	var b strings.Builder
	// Top border
	b.WriteString(bTL)
	for i, w := range cw {
		b.WriteString(strings.Repeat(bH, w+2))
		if i < nc-1 {
			b.WriteString(bTT)
		}
	}
	b.WriteString(bTR + "\n")

	for ri, row := range rows {
		b.WriteString(bV)
		for i := 0; i < nc; i++ {
			cell := ""
			if i < len(row) {
				cell = row[i]
			}
			pad := cw[i] - visibleLen(cell)
			b.WriteString(" " + cell + strings.Repeat(" ", pad) + " " + bV)
		}
		b.WriteString("\n")
		// Header separator
		if ri == 0 && len(rows) > 1 {
			b.WriteString(bLT)
			for i, w := range cw {
				b.WriteString(strings.Repeat(bH, w+2))
				if i < nc-1 {
					b.WriteString(bX)
				}
			}
			b.WriteString(bRT + "\n")
		}
	}

	// Bottom border
	b.WriteString(bBL)
	for i, w := range cw {
		b.WriteString(strings.Repeat(bH, w+2))
		if i < nc-1 {
			b.WriteString(bBT)
		}
	}
	b.WriteString(bBR + "\n\n")
	return b.String()
}

// ── Helpers ──

func (r *renderer) avail() int {
	w := r.width - (r.qd * 2) - (r.ld * 2)
	if w < 20 {
		w = 20
	}
	return w
}

func (r *renderer) plainText(n ast.Node) string {
	var b bytes.Buffer
	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		switch c.Kind() {
		case ast.KindText:
			b.Write(c.(*ast.Text).Segment.Value(r.src))
		case ast.KindString:
			b.Write(c.(*ast.String).Value)
		default:
			b.WriteString(r.plainText(c))
		}
	}
	return b.String()
}

func renderBox(lines []string, maxW int) string {
	// Find max content width
	contentW := 0
	for _, l := range lines {
		if w := visibleLen(l); w > contentW {
			contentW = w
		}
	}
	innerW := contentW + 2 // 1 space padding each side
	if innerW > maxW-2 {
		innerW = maxW - 2
	}
	if innerW < 4 {
		innerW = 4
	}

	var b strings.Builder
	b.WriteString(bTL + strings.Repeat(bH, innerW) + bTR + "\n")
	for _, l := range lines {
		pad := innerW - 2 - visibleLen(l)
		if pad < 0 {
			pad = 0
		}
		b.WriteString(bV + " " + l + strings.Repeat(" ", pad) + " " + bV + "\n")
	}
	b.WriteString(bBL + strings.Repeat(bH, innerW) + bBR + "\n")
	return b.String()
}

func wrapText(s string, width int) string {
	if width <= 0 {
		return s
	}
	words := strings.Fields(s)
	if len(words) == 0 {
		return ""
	}
	var lines []string
	cur := words[0]
	curW := visibleLen(cur)
	for _, w := range words[1:] {
		ww := visibleLen(w)
		if curW+1+ww > width {
			lines = append(lines, cur)
			cur = w
			curW = ww
		} else {
			cur += " " + w
			curW += 1 + ww
		}
	}
	lines = append(lines, cur)
	return strings.Join(lines, "\n")
}

func visibleLen(s string) int {
	n := 0
	i := 0
	for i < len(s) {
		if s[i] == '\x1b' {
			i++
			if i < len(s) && s[i] == '[' {
				i++
				for i < len(s) && !((s[i] >= 'A' && s[i] <= 'Z') || (s[i] >= 'a' && s[i] <= 'z')) {
					i++
				}
				if i < len(s) {
					i++
				}
			} else if i < len(s) && s[i] == ']' {
				i++
				for i < len(s) && s[i] != '\x07' {
					if s[i] == '\x1b' && i+1 < len(s) && s[i+1] == '\\' {
						i += 2
						break
					}
					i++
				}
				if i < len(s) && s[i] == '\x07' {
					i++
				}
			} else if i < len(s) {
				i++
			}
			continue
		}
		_, sz := utf8.DecodeRuneInString(s[i:])
		n++
		i += sz
	}
	return n
}
