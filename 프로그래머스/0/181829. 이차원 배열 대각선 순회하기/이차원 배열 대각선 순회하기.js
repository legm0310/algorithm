function solution(board, k) {
    return board.reduce((acc, cur, i) => {
        for(let j=0; j<cur.length; j++) {
            if(i+j<=k) acc+=cur[j]
        }
        return acc
    }, 0);
}