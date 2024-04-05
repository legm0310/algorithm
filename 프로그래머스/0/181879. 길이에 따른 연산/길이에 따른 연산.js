function solution(num_list) {
    let answer = num_list.reduce((acc, cur)=> {
        num_list.length<11 ? acc[0]*=cur : acc[1]+=cur
        return acc
    } ,[1,0])
    
    return answer[1]===0 ? answer[0] : answer[1]
}