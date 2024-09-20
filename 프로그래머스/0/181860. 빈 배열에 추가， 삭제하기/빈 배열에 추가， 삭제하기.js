function solution(arr, flag) {
    return arr.reduce((acc, cur, idx) => {
        if(flag[idx]) {
            return acc = acc.concat(new Array(cur*2).fill(cur))
        } else {
            acc.splice(acc.length - cur, cur)
            return acc
        }
    }, [])
}