const line = require("fs").readFileSync("/dev/stdin", "utf8");
let input = line.trim().split(" ");

if (+input[0] > +input[1]) console.log(">");
else if (+input[0] == +input[1]) console.log("==");
else console.log("<");
