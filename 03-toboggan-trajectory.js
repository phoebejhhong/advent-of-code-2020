const fs = require("fs");

const TREE = "#";

function getMark(rows, pos) {
  const numColumn = rows[0].length;
  const calculatedX = pos[0] % numColumn;
  return rows[pos[1]][calculatedX];
}

function traverseAndGetTreeCount(xStep, yStep) {
  const input = fs.readFileSync("./03-input.txt", "utf8");
  const rows = input.split("\n");
  if (!rows[rows.length - 1]) {
    rows.pop(); // ignore last newline
  }

  const numRow = rows.length;
  let treeCount = 0;
  let pos = [0, 0];

  const numberOfTravels = (numRow - 1) / yStep;
  for (let i = 0; i < numberOfTravels; i++) {
    pos = [pos[0] + xStep, pos[1] + yStep];
    const destination = getMark(rows, pos);
    if (destination === TREE) {
      treeCount++
    }
  }
  return treeCount;
}

let multiply = 1;
[[1, 1], [3, 1], [5, 1], [7, 1], [1,2]].forEach(rule => {
  const treeCount = traverseAndGetTreeCount(rule[0], rule[1]);
  console.log(`When traveling right ${rule[0]}, down ${rule[1]}, you would encounter ${treeCount} tree(s).`)
  multiply *= treeCount;
});

console.log(`📝 They multipy to ${multiply}.`)
