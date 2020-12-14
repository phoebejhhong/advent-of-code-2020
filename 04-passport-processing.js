const fs = require("fs");


function processPassport(stringPassport) {
  const map = {};
  stringPassport.split(" ").forEach(entry => {
    if (!entry) {
      return;
    }
    const [field, value] = entry.split(":");
    map[field] = value;
  });
  return map
}

const MANDATORY_FIELDS = [
  "byr",
  "iyr",
  "eyr",
  "hgt",
  "hcl",
  "ecl",
  "pid",
];

function isValidYear(year, min, max) {
  return /^\d{4}$/.test(year) && year >= min && year <= max;
}

function isValidHeight(value) {
  if (value.length < 3) {
    return false;
  }
  const unit = value.substring(value.length - 2);
  const height = value.substring(0, value.length - 2)

  if (unit === "cm") {
    return height >= 150 && height <= 193;
  } else if (unit === "in") {
    return height >= 59 && height <= 76;
  }
  return false;
}

function isValueValid (key, value) {
  switch (key) {
    case "byr":
      return isValidYear(value, 1920, 2002);
    case "iyr":
      return isValidYear(value, 2010, 2020);
    case "eyr":
      return isValidYear(value, 2020, 2030);
    case "hgt":
      return isValidHeight(value);
    case "hcl":
      return /^#[abcdef\d]{6}$/.test(value);
    case "ecl":
      return ["amb", "blu", "brn", "gry", "grn", "hzl", "oth"].indexOf(value) != -1;
    case "pid":
    return /^\d{9}$/.test(value);
    default:
      return false;
  }
}

function isPassportValid(passport) {
  for (i = 0; i < MANDATORY_FIELDS.length; i++) {
    const field = MANDATORY_FIELDS[i]
    const value = passport[field];
    if (!value || !isValueValid(field, value)) {
      return false;
    }
  }
  return true;
}

function processPassports() {
  const input = fs.readFileSync("./inputs/04-input.txt", "utf8");
  let currentPassport = "";
  let validCount = 0;

  input.split("\n").forEach(line => {
    if (!line) { // empty line means current passport content is finished
      const passport = processPassport(currentPassport);
      if (isPassportValid(passport)) {
        validCount++;
      }
      currentPassport = "";
    } else {
      currentPassport += (" " + line);
    }
  });

  return validCount;
}

const validCount = processPassports();
console.log(`📝 ${validCount} valid passports.`)
