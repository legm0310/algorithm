function solution(price) {
    if (price < 100000) {
        return Math.floor(price)
    } if (100000 <= price && price < 300000) {
        return Math.floor(price - price*(1/20))
    } if (300000 <= price && price < 500000) {
        return Math.floor(price - price*(1/10))
    } else {
        return Math.floor(price - price*(1/5))
    }
}