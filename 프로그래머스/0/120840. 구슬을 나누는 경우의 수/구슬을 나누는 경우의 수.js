function solution(balls, share) {
    function fact(num) {
        if (BigInt(num)===1n) return BigInt(1);
        return BigInt(num)*fact(BigInt(num)-BigInt(1));    
    }
    if(balls===share){
        return 1;
    }    
    return (fact(balls)/(fact((balls-share))*fact(share))); 
}

    // let fact = (num) =>  BigInt(num)=== 1n ? BigInt(1) : BigInt(num)*fact(BigInt(num)-BigInt(1))   
    // if(balls===share){
    //     return 1;
    // }    
    // return BigInt(fact(balls)/(fact((balls-share))*fact(share))); 
