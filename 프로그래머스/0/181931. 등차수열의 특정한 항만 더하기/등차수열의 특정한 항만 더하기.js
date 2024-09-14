function solution(a, d, included) {
    let answer = 0
    included.reduce((acc,cur,idx) => {
        if(included[idx]) answer+=acc
        return acc+=d
    }, a)
    return answer;
}