function solution(s) {
    return s.split("").every((v)=>!isNaN(v)) &&
        (s.length===4 || s.length===6);
}