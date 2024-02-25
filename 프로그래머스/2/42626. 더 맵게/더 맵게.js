function solution(scoville, K) {
    let count = 0,
        idx;
    const h = [...scoville.sort((a,b)=>a-b)]
    
    const bubbleUp = () => {
        idx = h.length-1
        let pIdx = Math.floor((idx - 1) / 2)
        
        while(h[pIdx] > h[idx]) {
            //swap
            [h[pIdx], h[idx]] = [h[idx], h[pIdx]]
            idx = pIdx;
            pIdx = Math.floor((idx - 1) / 2)
        }
    }
    const bubbleDown = () => {
        idx = 0
        let LChildIdx = idx * 2 + 1 
        let RChildIdx = idx * 2 + 2
        
        while(LChildIdx < h.length) {
            let smallerIdx = idx
            if(h[LChildIdx] && h[LChildIdx] < h[smallerIdx]) smallerIdx = LChildIdx
            if(h[RChildIdx] && h[RChildIdx] < h[smallerIdx]) smallerIdx = RChildIdx
            if(smallerIdx === idx) break
            
            //swap
            [h[idx], h[smallerIdx]] = [h[smallerIdx], h[idx]]
            idx = smallerIdx
            LChildIdx = idx * 2 + 1
            RChildIdx = idx * 2 + 2
        }
    }
    
    const push = (value) => {
        h.push(value)
        bubbleUp()
    }
    const del = () => {
        if(h.length < 2) return h.pop()
        const root = h[0]
        h[0] = h.pop()
        bubbleDown();
        return root
    }

    for(let i=0; i<scoville.length; i++) {
        if(h[0] < K) {
            const first = del()
            const second = del()

            push(first+(second*2))
            count++
        } else return count
    }  
    return -1
}