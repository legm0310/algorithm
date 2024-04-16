// 아래 정답
function solution(babbling) {
    return babbling.reduce((accu,cur)=>{
        if(cur.split(/aya|ye|woo|ma/g).join('')==='') accu+=1
        return accu
    }, 0)
}


// function solution(babbling) {
//     let a = ["aya", "ye", "woo", "ma"];
//     return babbling.reduce((accu,cur)=>{
//         let str = cur;
//         a.forEach(el => str = str.replace(el, ''))
//         if(str==='') accu+=1
//         return accu
//     }, 0)
// }