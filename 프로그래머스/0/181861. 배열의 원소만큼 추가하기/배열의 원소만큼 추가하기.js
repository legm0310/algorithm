function solution(arr) {
    return arr.reduce((result, v)=> {
        for(let i=0; i<v; i++){
            result.push(v)
        }
        return result
    }, [])
}