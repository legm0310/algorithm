function solution(dots) {
    let a = dots.map(v=> Math.sqrt((dots[0][0]-v[0])**2 + (dots[1][1]-v[1])**2)).sort((a,b)=>a-b);
    return a[1]*a[2];
}