function solution(clothes) {
    let count = 1;
    const table = {}
    for (const i of clothes) {
        if(!table[i[1]]) table[i[1]] = 1
        table[i[1]] = table[i[1]] + 1
    }
    for (const i in table) {
        count *= table[i]
    }
    return count -1
}




