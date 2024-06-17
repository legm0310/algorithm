function solution(polynomial) {
    const a = polynomial.split(" + ").reduce((acc, cur)=> {
        if(isFinite(+cur)=== true) { acc.num = (acc.num || 0) + +cur}
        else {
            let str = cur.replace(/[^0-9]/g, '')
            acc.x = (cur==='x'? (acc.x||0)+1 : (acc.x||0) + +str)
        }
        return acc
    }, {})
    a.x = (a.x ? (a.x===1 ? a.x='x' : a.x+'x'):null)
    return (a.x && a.num ? a.x+' + '+a.num : a.x ? a.x : a.num ? String(a.num) : null )
} 