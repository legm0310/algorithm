function solution(numbers) {
    let num = ["zero","one","two","three","four","five", "six","seven","eight","nine"]
    let obj = num.reduce((ac, curr, i) => {
        ac[curr] = i
        return ac
    },{})
    return +numbers.replace(/zero|one|two|three|four|five|six|seven|eight|nine/g, (match)=> {return obj[match]})
}