import * as fs from "fs";

export function getInput() {
  const [_, dayArg, fileArg] = process.argv;
  const day = dayArg.match(/d(\d+)\.ts$/)[1];
  const file = ["main", "test"].includes(fileArg) ? fileArg : "main";

  console.log("Running day " + day + " with " + file + " inputs");

  return fs
    .readFileSync(`${__dirname.replace("ts", "inputs")}/d${day}/${file}.txt`)
    .toString();
}
