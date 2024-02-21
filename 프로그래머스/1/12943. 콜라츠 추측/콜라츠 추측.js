function solution(num) {
    var answer = 0;
    let count = 0;
    
    function collatz(n) {
        if(n===1) return answer
        answer++
        if(count===500) return -1
        count++    
        if (n%2===0) return collatz(n/2)
        if (n%2!==0) return collatz((3*n)+1) 
    }
    return num===1 ? 0 : collatz(num)
}
