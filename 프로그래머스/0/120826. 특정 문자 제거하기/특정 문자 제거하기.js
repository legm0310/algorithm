function solution(my_string, letter) {
    let answer = "";
    for (let char of my_string) {
        if(char !== letter) { 
            answer = answer + char
        }
    }
    return answer;
}