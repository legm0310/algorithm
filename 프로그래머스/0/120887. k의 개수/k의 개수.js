function solution(i,j,k) {
    let str = "";
    for (let a=i ; a <=j ; a++) {
        str += a
    }
    return str.split("").reduce((a,b)=> {
        if(+b===+k) a+=1
        return a
    },0)
}