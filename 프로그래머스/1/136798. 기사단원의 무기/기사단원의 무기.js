function solution(number, limit, power) {
    let answer = 0
    for (let i=1 ; i<=number ; i++) {
        let num = 0
        for( let j=1 ; j<=Math.sqrt(i) ; j++) {
            if(i%j===0) {
                if(j**2===i)num+=1
                else num+=2
            }
        }
        if(num <= limit) answer += num
        else answer += power
    }
    return answer;
}