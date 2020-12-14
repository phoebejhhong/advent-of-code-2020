const fs = require("fs");

function findPair() {
  const map = {};
  const input = fs.readFileSync("./inputs/01-input.txt", "utf8");
  input.split("\n").forEach(line => {
    if (line) {
      map[Number(line)] = true
    }
  });

  let pair;
  Object.keys(map).find(num => {
    const theOtherNum = 2020 - num;
    if (map[theOtherNum]) {
      pair = [Number(num), theOtherNum];
      return true;
    }
    return false;
  });
  return pair;
}

function findTreeNumbers() {
  const map = {};
  const numbers = [];
  const input = fs.readFileSync("./inputs/01-input.txt", "utf8");
  input.split("\n").forEach(line => {
    if (line) {
      const number = Number(line);
      numbers.push(number);
      map[number] = true;
    }
  });

  for(i = 0; i < numbers.length - 1; i++) {
    const firstNum = numbers[i];
    for (j = i + 1 ; j < numbers.length; j++) {
      const secondNum = numbers[j];
      const pairSum = firstNum + secondNum;
      const thirdNum = 2020 - pairSum;
      if (map[thirdNum]) {
        return [firstNum, secondNum, thirdNum];
      }
    }
  }
}

const pair = findPair();
const threeNumbers = findTreeNumbers();
console.log(`📝 Part 1: ${pair.join(" & ")} mutiply to ${pair[0] * pair[1]}. \n📝 Part 2: ${threeNumbers[0]} & ${threeNumbers[1]} & ${threeNumbers[2]} multiply to ${threeNumbers[0] * threeNumbers[1] * threeNumbers[2]}.`);
