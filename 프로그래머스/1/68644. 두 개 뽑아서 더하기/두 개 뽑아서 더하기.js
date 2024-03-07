function solution(numbers) {
    const result = new Set()
    for(let i=0; i<numbers.length; i++) {
        for(let j=0; j<numbers.length; j++) {
            if(i!==j) result.add(numbers[i]+numbers[j])
        }  
    }
    return [...result].sort((a,b)=>a-b)
}