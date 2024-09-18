function solution(str1, str2) {
    let answer = ''
    for(let i=0; i<str1.length+str2.length; i++) {
        if(i%2===0) answer += str1[i/2]
        else answer += str2[Math.floor(i/2)]
    }
    return answer
}