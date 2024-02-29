function solution(s) {
    let n = 0
    return s.split("").map((v, i)=> {
        if(v===" ") {
            n = 0
            return " "
        }
        else if((n) % 2 !==0) {
            n++
            return v.toLowerCase()
        }
        else {
            n++
            return v.toUpperCase()
        }
    }).join("")
}