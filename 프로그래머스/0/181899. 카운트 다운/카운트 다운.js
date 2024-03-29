function solution(start, end_num) {
    let result = []
    for(let i=start; i>=end_num; i--) {
        result.push(i)
    }
    return result
}