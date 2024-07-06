function solution(numLog) {
    var answer = '';
    for(let i=1; i<numLog.length; i++) {
        if(numLog[i]-numLog[i-1]>0)
            answer += (Math.abs(numLog[i]-numLog[i-1])===1 ? "w" : "d")
        else answer += (Math.abs(numLog[i]-numLog[i-1])===1 ? "s" : "a")
    }
    return answer;
}