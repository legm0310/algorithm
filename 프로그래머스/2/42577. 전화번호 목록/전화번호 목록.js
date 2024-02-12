function solution(phone_book) {
    const hashMap = {}
    
    for (const num of phone_book) {
        hashMap[num]= true
    }
    
    for (const num of phone_book) {
        for (let i=0 ; i<num.length ; i++) {
            if(hashMap[num.substring(0, i)]) return false
        }
    } return true
}
