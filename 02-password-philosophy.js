const fs = require("fs");

function getNumberOfValidPWs(validationFn) {
  const input = fs.readFileSync("./inputs/02-input.txt", "utf8");
  let count = 0;
  input.split("\n").forEach(line => {
    if (!line) {
      return;
    }
    const [rule, middle, password] = line.split(" ");
    const [firstNum, secondNum] = rule.split("-");
    const mandatoryChar = middle[0];
    if (validationFn(Number(firstNum), Number(secondNum), mandatoryChar, password)) {
      count++
    }
  });
  return count;
}

function isValidPasswordPart1(min, max, mandatoryChar, password) {
  let mandatoryCharCount = 0;
  for (i = 0; i < password.length; i++) {
    if (password[i] === mandatoryChar) {
      mandatoryCharCount++;
      if (mandatoryCharCount > max) {
        return false;
      }
    }
  }
  return mandatoryCharCount >= min;
}

function isValidPasswordPart2(firstPos, secondPos, mandatoryChar, password) {
  if (password[firstPos - 1] === mandatoryChar) {
    return password[secondPos - 1] !== mandatoryChar;
  }

  return password[secondPos - 1] === mandatoryChar;
}

const firstCount = getNumberOfValidPWs(isValidPasswordPart1);
const secondCount = getNumberOfValidPWs(isValidPasswordPart2);
console.log(`📝 Part 1: ${firstCount}, Part 2: ${secondCount}`);
