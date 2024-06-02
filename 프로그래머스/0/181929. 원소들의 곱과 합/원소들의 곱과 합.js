function solution(num_list) {
    let sum = num_list.reduce((acc, cur)=>{
        acc[0]*=cur
        acc[1]+=cur
        return acc
    } ,[1, 0])
    return sum[0]<sum[1]**2 ? 1 : 0;
}