function solution(a, b) {
    let sum = 0
    let num = a>=b ? b : a,
        other = num===b ? a : b
    while(num<=other) {
        sum+=num
        num++
    }
    return sum
}