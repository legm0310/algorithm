const input = require("fs").readFileSync("/dev/stdin", "utf8");

for (let i = 0; i < 2 * input - 1; i++) {
  const abs = Math.abs(input - (i + 1));
  const blank = " ".repeat(abs);
  const star = "*".repeat(input * 2 - 1 - abs * 2);
  const output = `${blank}${star}`;
  console.log(output);
}
