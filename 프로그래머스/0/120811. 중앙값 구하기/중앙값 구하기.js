function solution(array) {
  return array.sort((a, b) => {
    return a - b;
  })[Math.floor(array.length / 2)];
}