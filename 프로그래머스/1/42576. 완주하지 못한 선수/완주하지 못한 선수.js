function solution(participant, completion) {
    let unfinish;
    participant.sort()
    completion.sort()
    
    for (let i=0; i<=participant.length; i++) {
        if(participant[i]!==completion[i]) {
            unfinish=participant[i]
            break;
        }
    }
    return unfinish
}