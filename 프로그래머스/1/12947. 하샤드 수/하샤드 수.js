function solution(x) {
    return x % (String(x).split("").reduce((a,c) => a+=parseInt(c), 0)) === 0 ? true : false
}