function solution(priorities, location) {
    let count = 0
    const table = []
    let first = Math.max(...priorities)

    priorities.forEach((v, i) => {
        table.push(i)
    })
    
    while(priorities.length !== 0) {
        if (priorities[0] === first) {
            count++
            priorities.shift()
            if(table.shift() === location) return count
            first =Math.max(...priorities)
        } 
        if (priorities[0] < first){
            priorities.push(priorities.shift())
            table.push(table.shift())
        }
    }
    return count
}