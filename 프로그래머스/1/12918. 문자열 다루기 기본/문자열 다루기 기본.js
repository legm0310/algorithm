function solution(s) {
    return s.split("").filter(v=>!isNaN(+v)).length === s.length && 
        (s.length===4 || s.length===6)     
}