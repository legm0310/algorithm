function solution(myString) {
    return myString.split("").map((v, i)=>"l".charCodeAt(0) > myString.charCodeAt(i) ? "l" : v).join("")
     
}