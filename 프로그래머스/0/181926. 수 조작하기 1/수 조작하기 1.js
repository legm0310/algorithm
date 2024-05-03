function solution(n, control) {
    const wsda = {
        "w": 1,
        "s": -1,
        "d": 10,
        "a": -10
    }
    return control.split("").reduce((acc,cur)=>acc+wsda[cur], n)
}