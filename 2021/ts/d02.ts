import { getInput } from "./helpers";

interface Position {
  x: number;
  y: number;
  aim: number;
}

function getMoveInput(input: string, aim: number): Position {
  const [direction, amount] = input.split(" ");

  if (direction === "forward") {
    return { x: +amount, y: aim * +amount, aim };
  }
  if (direction === "down") {
    return { x: 0, y: 0, aim: aim + +amount };
  }
  return { x: 0, y: 0, aim: aim + -amount };
}

function getPosition(): number {
  const finalPosition = getInput()
    .split("\n")
    .reduce<Position>(
      (final, input) => {
        const position = getMoveInput(input.trim(), final.aim);
        return {
          x: final.x + position.x,
          y: final.y + position.y,
          aim: position.aim,
        };
      },
      { x: 0, y: 0, aim: 0 }
    );

  return finalPosition.x * finalPosition.y;
}

console.log(getPosition());
