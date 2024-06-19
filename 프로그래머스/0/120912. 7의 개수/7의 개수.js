function solution(array) {
    return array.join('').split('').reduce((acc, cur) => {
         return +cur===7 ? acc+1 : acc
    }, 0)
}