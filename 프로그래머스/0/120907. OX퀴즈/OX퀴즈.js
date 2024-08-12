function solution(quiz) {
    return quiz.map((v) => {
        let arr = v.split(" ")
        if(arr[1] ==="+") {
            return (+arr[0]+ +arr[2] === +arr[4] ? "O":"X")
        } if(arr[1] ==="-"){
            return (+arr[0]- +arr[2] === +arr[4] ? "O":"X")
        }
    })  
}


