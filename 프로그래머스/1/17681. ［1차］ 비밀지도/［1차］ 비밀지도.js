function solution(n, arr1, arr2) {
    var answer = [];
    for (let i=0; i<n; i++) {
        let binary = (arr1[i] | arr2[i]).toString(2)
        if(binary.length<n) answer.push("0".repeat(n-binary.length)+binary)
        else answer.push(binary)
        answer[i] = answer[i].replaceAll("1", "#")
        answer[i] = answer[i].replaceAll("0", " ")
    }
    return answer;
}