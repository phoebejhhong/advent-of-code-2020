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

const MANEATORY_FIELDS = [
  "byr",
  "iyr",
  "eyr",
  "hgt",
  "hcl",
  "ecl",
  "pid",
];

function isPassportValid(passport) {
  for (i = 0; i < MANEATORY_FIELDS.length; i++) {
    if (!passport[MANEATORY_FIELDS[i]]) {
      return false;
    }
  }
  return true;
}

function processPassports() {
  const input = fs.readFileSync("./04-input.txt", "utf8");
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
