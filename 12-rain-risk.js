const fs = require("fs");


function directionToCoord(direction, number, currentCoord) {
  switch (direction) {
    case "N":
      return [
        currentCoord[0],
        currentCoord[1] + number,
      ];
    case "S":
      return [
        currentCoord[0],
        currentCoord[1] - number,
      ];
    case "E":
      return [
        currentCoord[0] + number,
        currentCoord[1],
      ];
    case "W":
      return [
        currentCoord[0] - number,
        currentCoord[1],
      ];
  }
}

const clockwiseDirections = ["N", "E", "S", "W"];
function rotateToDirection(rotate, degree, currentDirection) {
  const turns = degree / 90;
  const currentIndex = clockwiseDirections.findIndex(d => d === currentDirection);
  let newIndex = 0;
  switch (rotate) {
    case "L":
      newIndex = currentIndex - turns
      if (newIndex < 0) {
        newIndex += 4;
      }
      break;
    case "R":
      newIndex = currentIndex + turns;
      if (newIndex > 3) {
        newIndex -= 4;
      }
    break;
  }
  return clockwiseDirections[newIndex];
}

function rotateWayPoint(rotate, degree, currentWaypoint) {
  switch (rotate) {
    case "R":
      switch (degree) {
        case 90:
          return [
            currentWaypoint[1],
            -currentWaypoint[0],
          ];
        case 180:
          return [
            -currentWaypoint[0],
            -currentWaypoint[1],
          ];
        case 270:
          return [
            -currentWaypoint[1],
            currentWaypoint[0],
          ];
      }
    case "L":
    switch (degree) {
      case 90:
        return [
          -currentWaypoint[1],
          currentWaypoint[0],
        ];
      case 180:
        return [
          -currentWaypoint[0],
          -currentWaypoint[1],
        ];
      case 270:
        return [
          currentWaypoint[1],
          -currentWaypoint[0],
        ];
    }
  }
}
function moveShip(ship, instruction) {
  const [ command, ...rest ] = instruction;
  const number = Number(rest.join(""));
  switch (command) {
    case "N":
    case "S":
    case "E":
    case "W":
      ship.coord = directionToCoord(command, number, ship.coord)
      break;
    case "F":
      ship.coord = directionToCoord(ship.direction, number, ship.coord)
      break;
    case "L":
    case "R":
      ship.direction = rotateToDirection(command, number, ship.direction);
      break;
  }
}

function moveShip2(ship, instruction) {
  const [ command, ...rest ] = instruction;
  const number = Number(rest.join(""));
  switch (command) {
    case "N":
    case "S":
    case "E":
    case "W":
      ship.waypoint = directionToCoord(command, number, ship.waypoint)
      break;
    case "F":
      ship.coord = [
        ship.coord[0] + ship.waypoint[0] * number,
        ship.coord[1] + ship.waypoint[1] * number,
      ]
      break;
    case "L":
    case "R":
      ship.waypoint = rotateWayPoint(command, number, ship.waypoint)
      break;
  }
}

function main() {
  const input = fs.readFileSync("./inputs/12-input.txt", "utf8");
  const instructions = input.split("\n");
  instructions.pop();

  const ship = {
    direction: "E",
    coord: [0, 0],
  }
  const ship2 = {
    coord: [0, 0],
    waypoint: [10, 1], // relative to ship
  }
  instructions.forEach(instruction => {
    moveShip(ship, instruction);
    moveShip2(ship2, instruction);
  });
  console.log(`📝 Manhattan distance traveled: Part 1- ${Math.abs(ship.coord[0]) + Math.abs(ship.coord[1])}; Part 1- ${Math.abs(ship2.coord[0]) + Math.abs(ship2.coord[1])}`);
}
main();
