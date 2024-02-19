function solution(progresses, speeds) {
    var answer = [];
    let queue = [];
    
    progresses.forEach((progress, index) => {
        let remain = 100-progress
        queue.push(Math.ceil(remain/speeds[index]))
    })

    for(let i=0; i<progresses.length; i++){
        let deployIdx = 0
        for(const [index, value] of queue.entries()) {
            if(queue[0] < value) {
                answer.push(index)
                deployIdx = index
                break
            }
        }
        if(Math.max(...queue)===queue[0]) {
            answer.push(queue.length)
            break
        }
        queue.splice(0, deployIdx)
    }
    return answer

}
