function solution(k, dungeons) {
    let answer = 0;
    const count = (arr) => {
        let curK = k
        const n = arr.reduce((acc, cur) => {
            if(curK >= cur[0]) {
                curK -= cur[1]
                return acc+=1
            } else return acc
        }, 0)
        return n
    }
    const permute = (arr, depth = 0) => {
        if(depth === arr.length) {
            const cnt = count(arr)
            if(answer < cnt) answer = cnt 
        }
        for(let i=depth; i<arr.length; i++) {
            [arr[depth], arr[i]] = [arr[i], arr[depth]];
            permute(arr, depth + 1);
            [arr[depth], arr[i]] = [arr[i], arr[depth]];   
        }
    }
    permute(dungeons)
    return answer;
}