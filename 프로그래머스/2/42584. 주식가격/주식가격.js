function solution(prices) {
    const qu = [];

    for (let i = 0; i < prices.length; i++) {
        qu.push(0);
        for (let j = i+1; j < prices.length; j++) {
            if (prices[i] > prices[j]) {
                qu[i] = j - i;
                break;
            } else {
                qu[i]++;
            }
        }
    }
    return qu;
}