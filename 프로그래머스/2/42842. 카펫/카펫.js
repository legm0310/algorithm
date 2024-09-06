function solution(brown, yellow) {
    let width
    let height
    for(let i=0; i<brown; i++) {
        for(let j=0; j<brown; j++) {
            if(i>j) continue
            if((i)*(j)==yellow+brown && (2*i+2*j)-4 == brown) {
                width = j
                height = i
                break
            }
        }
    }
    return [width, height]
}
