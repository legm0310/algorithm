function solution(num_list) { 
    var answer = [];
    let i = 0
    while(i < num_list.length) {
        answer[i] = num_list[(num_list.length-1)-i];
        i++;
    }
    return answer;
}