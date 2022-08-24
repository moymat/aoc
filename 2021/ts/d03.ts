import { getInput } from "./helpers";

enum MostCommon {
  Zero,
  One,
  Equal,
}

function whichIsMostCommon(inputs: string[], idx: number): MostCommon {
  const ones = inputs.reduce(
    (sum, line) => (line[idx] === "1" ? sum + 1 : sum),
    0
  );
  const zeros = inputs.length - ones;
  return ones > zeros
    ? MostCommon.One
    : ones < zeros
    ? MostCommon.Zero
    : MostCommon.Equal;
}

function getRates() {
  const lines = getInput().split(/\s+/g);
  const firstLineArr = lines[0].split("");

  const gammaRate = firstLineArr
    .map((_, i) =>
      whichIsMostCommon(lines, i) !== MostCommon.Zero ? "1" : "0"
    )
    .join("");
  const epsilonRate = gammaRate
    .split("")
    .map((bit) => (bit === "0" ? "1" : "0"))
    .join("");

  const oxygenRate = firstLineArr.reduce(
    (list, _, i) =>
      list.length === 1
        ? list
        : whichIsMostCommon(list, i) !== MostCommon.Zero
        ? list.filter((line) => line[i] === "1")
        : list.filter((line) => line[i] === "0"),
    [...lines]
  );

  const co2Rate = firstLineArr.reduce(
    (list, _, i) =>
      list.length === 1
        ? list
        : whichIsMostCommon(list, i) === MostCommon.Zero
        ? list.filter((line) => line[i] === "1")
        : list.filter((line) => line[i] === "0"),
    [...lines]
  );

  console.log(parseInt(oxygenRate[0], 2), parseInt(co2Rate[0], 2));

  return [
    parseInt(gammaRate, 2) * parseInt(epsilonRate, 2),
    parseInt(oxygenRate[0], 2) * parseInt(co2Rate[0], 2),
  ];
}

console.log(getRates());
