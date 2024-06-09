function solution(chicken) {
    let c = 0
    let cur = chicken
    while(cur > 1) {
        c+=(cur/10)
        cur=(cur/10)
    }
    return Math.floor(c)
}