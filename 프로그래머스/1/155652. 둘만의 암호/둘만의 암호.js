function solution(s, skip, index) {
    return s.split("").map((v)=>{
        let s = v
        for(let i=0 ; i<index ; i++) {
            s = String.fromCharCode( ((s.charCodeAt()+1 - 97) % 26 + 97) )
            skip.indexOf(s) !== -1 && i--
        }
        return s
    }).join("")
}

