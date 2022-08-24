import { getInput } from "./helpers";

function computeNbOfIncrease() {
  return getInput()
    .split("\n")
    .map((_, i, arr) =>
      i === 0 || i === arr.length - 2 ? 0 : +arr[i + 2] - +arr[i - 1]
    )
    .filter((nb) => nb > 0).length;
}

console.log(computeNbOfIncrease());
