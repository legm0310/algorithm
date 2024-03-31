function solution(babbling) {
    const words = ["aya", "ye", "woo", "ma"]
    return babbling.filter((v) => {
        let temp = v
        words.forEach((w, i) => {
            if(temp.includes(w.repeat(2))) return
            temp = temp.replaceAll(w,`${i}`)
        })
        for (let j=0; j<words.length; j++) {
            temp = temp.replaceAll(`${j}`,"")
        }
        return temp.length===0
    }).length
}