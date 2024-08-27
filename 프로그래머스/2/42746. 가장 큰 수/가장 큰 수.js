function compare(a, b) {
    const strA = String(a)
    const strB = String(b)
    return strA+strB > strB+strA ? -1: 1
}

function solution(numbers) {
    const result = numbers.sort((a,b) => compare(a,b)).join("")
    return !!(+result) ? result : "0"
}
                 
                 