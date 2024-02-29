function solution(number) {
    let answer = 0;
    for(let i=0; i<number.length; i++) {
        for(let j=i; j<number.length; j++) {
            for(let k=j; k<number.length; k++) {
                if(i===j || i===k || j===k) continue
                if(number[i]+number[j]+number[k]===0) answer++ 
            }
        }
    }
    return answer;
}