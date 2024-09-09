function solution(a, b) {
    return parseInt([a, b].join("")) >= parseInt([b, a].join("")) ? 
        parseInt([a, b].join("")) : parseInt([b, a].join(""))
}