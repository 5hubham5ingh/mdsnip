#!/usr/bin/js

const options = ["Release a new version", "Start a dev server"];

const what = enquire.choose(options);

if (what === options[1]) {
  $`npx serve .`;
  std.exit(0);
}

print("Last version:");
$`git describe --tags --abbrev=0`.log();

const version = enquire.ask("New version");

if (
  !enquire.confirm(
    `Begin building version "${version}"?`,
  )
) std.exit(1);

if (enquire.confirm("Compile for release?")) {
  if (os.exec("bash build.sh".split(" "))) std.exit(1);
}

const cmds = [
  "git add index.html static/* TODO.md",
  `git commit -m ${version}`,
  "git push origin main",
  `git tag ${version}`,
  `git push origin ${version}`,
];

for (cmd of cmds) {
  if (!enquire.confirm("Run" + `"${cmd}"`)) break;
  os.exec(cmd.split(" "));
}
