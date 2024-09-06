function solution(emergency) {
  let result = [];
  let arr = [...emergency];
    
  emergency.sort(function compare(a, b) {
    return b - a;
  });
    
  for (let i = 0; i < emergency.length; i++) {
    result[arr.indexOf(emergency[i])] = i + 1;
  }
  return result;
}