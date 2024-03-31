function solution(rsp) {
    return rsp.toString().split("").map(x => +x===0 ? 5 : (+x===2 ? 0 : (+x===5 ? 2 : null))).join('');
}