function solution(n, m, section) {
    let cnt = 1
    let wall = section[0]
    for (let i = 0 ; i <= section.length; i++) {
        if (section[i] - wall > m - 1) {
            cnt++
            wall = section[i]
        }
    }
    return cnt
}