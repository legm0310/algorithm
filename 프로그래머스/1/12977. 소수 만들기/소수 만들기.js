function solution(nums) {
    let answer = 0;
    const isPrime = (num) => {
        if (num===1) return false
        for(let i=2 ; i<=Math.sqrt(num); i++) {
            if(num%i==0) return false
        } 
        return true
    }
    const getNums = (arr, curCombination = [], startIdx = 0) => {
        if(curCombination.length === 3) {
            const sum = curCombination.reduce((acc,cur)=>acc+=cur, 0)
            if(isPrime(sum)) answer++
            return
        }
        for(let i=startIdx; i<arr.length; i++) {
            getNums(arr, [...curCombination, arr[i]], i+1)
        }
    }
    getNums(nums)

    return answer;
}