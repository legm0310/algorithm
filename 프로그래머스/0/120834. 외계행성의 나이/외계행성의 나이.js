// function solution(age) {
//     let ages = {};
//     for(let i=0 ; i<=9 ;i++) {
//         ages[i] = String.fromCodePoint(97+i);
//     }
//     return [(ages[age.toString()[0]]), 
//             (ages[age.toString()[1]]), 
//             (ages[age.toString()[2]]),
//             (ages[age.toString()[3]])].join('') ;
// }
function solution(age) {
  let str = "";
  let ages = age.toString().split("");
  for (let n of ages) {
    str += String.fromCodePoint(+n + 97);
  }
  return str;
}