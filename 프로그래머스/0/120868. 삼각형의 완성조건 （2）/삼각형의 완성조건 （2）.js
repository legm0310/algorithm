function solution(sides) {
    sides.sort((a, b)=> b-a);
    return (sides[0]!==sides[1] ? sides[0]-(sides[0]-sides[1]) : 0) + 
        (sides[0]!==sides[1] ? ((sides[0]+sides[1]) - sides[0]-1) : sides[0]+sides[1]-1)
}
                 
