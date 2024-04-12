function solution(num_list) {
    let sum = num_list.reduce((acc, cur, i)=>{
        i%2===0 ? acc[0]+=cur : acc[1]+=cur
        return acc
    }, [0, 0])
    return sum[0]>=sum[1] ? sum[0] : sum[1]
}