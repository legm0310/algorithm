const line = require("fs").readFileSync("/dev/stdin", "utf8");
let input = line.trim()
const isPalindrome = input === input.split("").reverse().join("") ? 1 : 0;
console.log(+isPalindrome);