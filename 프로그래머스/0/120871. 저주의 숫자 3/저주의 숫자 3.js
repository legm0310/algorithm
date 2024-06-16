function solution(n) {
    let answer = 0
    for(let i = 1; i <= n ; i++) {
        while( answer%3==0 || Math.floor(answer%10) ==3 || Math.floor(answer/10)==3 || Math.floor(answer/10)==13) {answer++}
        answer++
    }    
    return answer-1
}
