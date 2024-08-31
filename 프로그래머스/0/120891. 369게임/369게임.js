function solution(order) {
    let result = 0;
    order.toString().split("").map(v=> {if(+v%3===0&& +v!==0) result++})
    return result;

}