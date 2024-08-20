function solution(denum1, num1, denum2, num2) {
    //기약분수 전 분모 분자
    let answer = [];
    let moth = num1*num2;
    let son = denum1*num2 + denum2*num1;
    let gcd = 1;
    //최대공약수 구하기
    for(let i = 2; i<moth+son ;i++) {
        if(son%i===0 && moth%i===0) {
            gcd = i;
        }
    }
    answer.push(son/gcd)
    answer.push(moth/gcd)
    return answer;
}