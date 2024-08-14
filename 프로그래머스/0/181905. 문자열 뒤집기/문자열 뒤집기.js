function solution(my_string, s, e) {
    const arr = my_string.slice(s, e+1)
    return my_string.replace(arr, arr.split("").reverse().join(""));
}