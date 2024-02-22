function solution(s) {
    return s.split("").sort((a,b)=>b.charCodeAt(b)-a.charCodeAt(a)).join("");
}