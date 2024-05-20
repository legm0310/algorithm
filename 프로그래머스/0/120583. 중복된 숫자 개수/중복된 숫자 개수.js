function solution(array, n) {
    return array.reduce((acc, cur)=> {
        cur==n ? acc++ : acc
        return acc
    }, 0)
}