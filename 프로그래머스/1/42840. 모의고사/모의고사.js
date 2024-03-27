function solution(answers) {
    var answer = []
    const students = [
        [1, 2, 3, 4, 5],
        [2, 1, 2, 3, 2, 4, 2, 5],
        [3, 3, 1, 1, 2, 2, 4, 4, 5, 5]
    ]
    const counts = students.reduce((a, c) => {
        let count = answers.filter((v, i) => v === c[i%c.length]).length
        a.push(count)
        return a
    }, [])
    const max = Math.max(...counts)
    for(let i=0; i<students.length; i++) {
        if (counts[i] === max) answer.push(i+1)
    }
    return answer
}