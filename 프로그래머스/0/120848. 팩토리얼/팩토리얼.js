function solution(n) {
    let fact = 1;
    for(let i = 2; i<= 10; i++){
        fact *= i 
        if(fact===n) return i;
        if(fact>n) return i-1;
    }
}