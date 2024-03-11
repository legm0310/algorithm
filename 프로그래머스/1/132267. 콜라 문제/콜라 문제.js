function solution(a, b, n) {
    var answer = 0;
    while(n/a>=1) {
        answer += b*Math.floor(n/a)
        n = n-(a*Math.floor(n/a)) + (b*Math.floor(n/a))
    }
    return answer;
}

