function solution(s){
    let count = 0
    const stack = []
    s.split("").forEach((v)=>{
        if(v==="(") {
            count++
            stack.push(1)
        }
        else {
            count--
            stack.pop()
        }
    })
    return stack.length===0 && count===0 ? true : false;
}