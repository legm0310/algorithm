function solution(lines) {
    const m = Array.from({length:200},()=>new Set())
    lines.map((e,idx) => {
        for(let i = e[0] ; i<e[1] ; i++) {
            m[i+100].add(idx)
        }
    })
    return m.reduce((acc,cur)=> {
        if([...cur].length > 1) acc++
        return acc
    },0)
}
// function solution(lines) {
//     function i(l1, l2) {
//         console.log(l1,l2)
//         if(l1[0] <= l2[0] && l1[1] >= l2[1]) {
//             console.log('case1')
//             return Math.abs(l2[1]) - Math.abs(l2[0])
//         }
//         else if(l1[0] >= l2[0] && l1[1] <= l2[1]) {
//             console.log("case2")
//             return Math.abs(l1[1]) - Math.abs(l1[0])                                         
//         }
//         // else if(l1[0] < l2[1]) return Math.abs(l1[1]) - Math.abs(l1[0]) 
//         // else if(l1[1] < l2[0]) return Math.abs(l1[1]) - Math.abs(l1[0]) 
//         else return 0
//     }
    
//     return i(lines[0], lines[1])+i(lines[1], lines[2]) +i(lines[0], lines[2])
// }

//배열을 만들어서 1열로 쭉 세운 후 겹치는 부분의 갯수를 샌다.

// function solution(lines) {
//     let answer = 2
//     let arr = [...Array.from({length:lines[0][1]-lines[0][0]+1},(v,i)=>+lines[0].slice(0,1)+i),
//            ...Array.from({length:lines[1][1]-lines[1][0]+1},(v,i)=>+lines[1].slice(0,1)+i),
//            ...Array.from({length:lines[2][1]-lines[2][0]+1},(v,i)=>+lines[2].slice(0,1)+i)].sort((a,b)=>a-b)
//     console.log(arr)
//     return arr.filter((x, i, a)=> (x===a[i-1]))
// }


// function solution(lines) {
    
//     return lines.reduce((acc, cur, idx, arr) => {
//         let a = lines.slice(idx+1, idx+2).reduce((acc, val) => [ ...acc, ...val ], [])
//         console.log(a[0] ,cur[1])
//         if(arr[idx+1]&&a[0]<cur[1]) {
//             // console.log("sa")
//             console.log(a[0]-cur[1])
//             return acc += Math.abs(a[0]-cur[1] )
//         } 
//         else return acc
//     }, 0)
// }
