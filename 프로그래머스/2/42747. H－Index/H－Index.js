function solution(citations) {
    let h = 0
    citations.sort((a,b) => b-a).forEach((v,i) => {
        if(i+1<=v) h += 1
    })
    return h
}
