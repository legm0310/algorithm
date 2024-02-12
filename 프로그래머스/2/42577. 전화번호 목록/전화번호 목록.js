function solution(phone_book) {
    const obj = {}
    for (let i=0 ;i<phone_book.length; i++) {
        obj[phone_book[i]] = true
    }
    for (let i=0 ;i<phone_book.length; i++) {
        const num = phone_book[i];
        for (let j=0 ;j<phone_book[i].length; j++) {
            if(obj[num.slice(0, j)]) return false
        }
    }
    return true
}

