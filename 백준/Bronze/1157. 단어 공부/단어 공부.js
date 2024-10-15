const line = require("fs").readFileSync("/dev/stdin", "utf8");
let input = line.trim();

const arr = input.toLowerCase().split("");
const table = arr.reduce((acc, cur) => {
  acc[cur] = (acc[cur] || 0) + 1;
  return acc;
}, {});
const sorted = Object.keys(table).sort((a, b) => table[b] - table[a]);

const result = table[sorted[0]] === table[sorted[1]] ? "?" : sorted[0].toUpperCase();
console.log(result);

