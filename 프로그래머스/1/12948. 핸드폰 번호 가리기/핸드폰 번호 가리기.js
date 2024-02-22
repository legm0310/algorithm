function solution(phone_number) {
    return phone_number.split("").map((v, i)=> {
        if(i<phone_number.length-4) return "*"
        else return v
    }).join("");
}