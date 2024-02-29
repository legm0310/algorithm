function solution(s, n) {
    return s.split("").map((v, i)=> {
        const charCode = s.charCodeAt(i)
        let pushedCode;
        if(v===" ") return " "
        else {
            if(charCode >= 65 && charCode <= 90) {
                pushedCode = charCode+n > 90 ? (charCode+n)%90 + 64 : charCode+n
            } else {
                pushedCode = charCode+n > 122 ? (charCode+n)%122 + 96 : charCode+n
            }
            console.log(pushedCode)
            return String.fromCharCode(pushedCode)
        }
    }).join("")
}