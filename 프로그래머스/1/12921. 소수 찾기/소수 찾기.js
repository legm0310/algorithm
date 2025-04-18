function solution(n) {
    let arr = [false, false, ...Array(n-1).fill(true)]
    for(let i=2; i*i <= n; i++) {
        if(arr[i]) {
            for(let j = i*i; j <= n; j+=i) {
                arr[j] = false
            }
        }
    }
    return arr.filter(e=>e).length
}
