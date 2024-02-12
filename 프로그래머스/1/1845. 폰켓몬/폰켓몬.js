function solution(nums) {
    return [...new Set(nums)].length < nums.length/2 ? [...new Set(nums)].length : nums.length/2
}