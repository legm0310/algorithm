function solution(dot) {
    return Math.sign(dot[0]*dot[1])===1 ? (Math.sign(dot[0])===1 ?  1 :  3) : (Math.sign(dot[0])===1 ?  4 : 2)
}