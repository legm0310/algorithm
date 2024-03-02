function solution(s) {
    let arr = s.split("").reverse()
    return arr.map((v, i)=> {
        return arr.slice(i).findIndex((e,idx)=>e===v&&idx!==0)
    }).reverse()
}
