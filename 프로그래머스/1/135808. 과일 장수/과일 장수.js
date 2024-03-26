function solution(k, m, score) {
    score.sort((a,b) => b-a).slice(0,m)
    return score.reduce((a, c, i) => {
        const box = score.slice(i*m, i*m+m)
        if (box.length === m) a += m*box.pop()
        return a
    }, 0)
}