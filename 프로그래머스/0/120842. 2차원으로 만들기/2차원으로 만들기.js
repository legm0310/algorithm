function solution(num_list, n) {
    let answer = [];
    while(num_list.length>0){
        answer.push(num_list.slice(0,n))
        num_list.splice(0,n)
    }
    return answer
}