function solution(my_string) {
    return my_string.split('').filter(x => isFinite(x)===true).map(v => +v).sort((a, b) => a - b);
}