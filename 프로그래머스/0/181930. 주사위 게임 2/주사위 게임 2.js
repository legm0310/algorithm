function solution(a, b, c) {
    const num = new Set(arguments)
    const count = [...num].length===1 ? 3 : [...num].length===2 ? 2 : 1
    let sum = 1
    for(let i = 1 ; i <= count ; i++) {
        sum*=(a**i)+(b**i)+(c**i)
    }
    return sum
}