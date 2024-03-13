function solution(k, score) {
    const top = [];
    return score.map((v, i) => {
        if(i+1 > k) {
            if(v > Math.min(...top)) top.splice(top.indexOf(Math.min(...top)), 1, v);
        }
        else top.push(v)
        return Math.min(...top)
    })
}