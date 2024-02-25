function solution(n, m) {
    var answer = [1, 1];
    for(let i=2; i<=Math.min(n, m); i++) {
        if(n%i===0 && m%i===0) answer[0] = i
    }
    while(1) {
        if(answer[1]%n===0 && answer[1]%m===0) break
        answer[1] ++
    }
    return answer;
}