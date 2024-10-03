function solution(N, stages) {
    let user = stages.length;
    const arr = [];
    const table = stages.reduce((acc, cur) => {
        if(cur>N) return acc
        acc[cur] = (acc[cur] || 0) + 1
        return acc;
    }, {})
    
    for(let i = 1; i <= N; i++) {
        arr.push(i)
        if(table[i]) {
            let tmp = table[i] / user
            user -= table[i]
            table[i] = tmp
        } else table[i] = 0
    }
    
    return arr.sort((a, b) => table[b] - table[a])
}