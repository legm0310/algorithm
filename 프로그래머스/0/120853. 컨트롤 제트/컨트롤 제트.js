function solution(s) {
    let arr = [];
    s.split(" ").map((v) => {isFinite(+v) === true ? arr.push(+v) : arr.pop()});
    return arr.reduce((a,b) => {return a+b},0)
} 




                      


