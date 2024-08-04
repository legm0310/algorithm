function solution(num, k) {
    let result = +String(num).split('').map((v, i) => {
        return +v===k ? i+1 :''
    }).join('').slice(0,1)
    return result ? result : -1
}