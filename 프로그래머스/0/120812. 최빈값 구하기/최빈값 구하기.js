function solution(array) {
  let count;
  let countKeys;
  let countValue;
  let dup;
  let maxValueIndex;
  let maxValue;

  if (array.length === 1) {
    return +array[0];
  }

  count = array.reduce((accu, curr) => {
    accu[curr] = (accu[curr] || 0) + 1;
    return accu;
  }, {});
  countValue = Object.values(count); //count value
  countKeys = Object.keys(count);
  maxValueIndex = countValue.indexOf(Math.max(...countValue));
  if (Math.max(...countValue) === 1) return -1;

  dup = countValue.reduce((accu, curr) => {
    accu[curr] = (accu[curr] || 0) + 1;
    return accu;
  }, {});

  maxValue = Object.keys(dup)[Object.keys(dup).length - 1];
  if (dup[maxValue] > 1) {
    return -1;
  }
  return +countKeys[maxValueIndex];
}