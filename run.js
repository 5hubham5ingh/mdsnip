#!/usr/bin/js

const options = [
  "Release a new version",
  "Start a dev server",
  "Re-release last version",
];

const choice = enquire.choose(options);

if (choice === options[1]) {
  $`npx serve .`;
  std.exit(0);
}

if (choice === options[2]) {
  const lastVersion = exec("git describe --tags --abbrev=0");
  if (!enquire.confirm(`Re-release version "${lastVersion}"?`)) std.exit(0);
  $`git tag -d ${lastVersion}`.log();
  $`git push origin --delete ${lastVersion}`.log();
  $`git push origin ${lastVersion}`.log();
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
  if (!scriptArgs.includes("-y") && !enquire.confirm("Run" + `"${cmd}"`)) break;
  os.exec(cmd.split(" "));
}
