function solution(cards1, cards2, goal) {
    const goal2 = [...goal]
    for(let i=0; i<Math.max(cards1.length, cards2.length); i++) {
        if(cards1[i]) goal2.splice(goal2.indexOf(cards1[i]), 1)
        if(cards2[i]) goal.splice(goal.indexOf(cards2[i]), 1)
    }
    if (cards1.join().includes(goal.join()) && cards2.join().includes(goal2.join()) || goal.length===0) return "Yes"
    else return "No"
}