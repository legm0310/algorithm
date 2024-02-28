function solution(d, budget) {
    let answer = 0
    d.sort((a,b)=>a-b).forEach(v => {
        if(budget-v >= 0) {
            answer++
            budget-=v
        }
        else return
    })
    return answer;
}