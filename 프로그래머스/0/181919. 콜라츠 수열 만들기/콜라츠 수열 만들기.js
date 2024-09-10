function solution(n) {
    let answer = [];
    const a = (num) => {
        answer.push(num)
        if(num===1) return
        let newNum = num%2==0 ? num/2 : num*3 + 1
        a(newNum)
    }
    a(n)
    return answer;
}