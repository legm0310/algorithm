class MinHeap {
    constructor() {
        this.heap = []
    }
    
    getlength() {
        return this.heap.length
    }
    
    swap(idx, idx2) {
        [this.heap[idx], this.heap[idx2]] = [this.heap[idx2], this.heap[idx]]
        return
    }
    
    push(value) {
        this.heap.push(value)
        if(this.heap.length === 1) return
        let idx = this.heap.length - 1
        let pIdx = Math.floor((idx-1) / 2)
        
        while(this.heap[idx] && this.heap[pIdx] &&
              this.heap[idx][1] < this.heap[pIdx][1]) {
            this.swap(idx, pIdx)
            idx = pIdx
            pIdx = Math.floor((idx-1) / 2)
        }
    }
    
    del() {
        if (this.heap.length === 1) return this.heap.pop();
        const root = this.heap[0]
        this.heap[0] = this.heap.pop()
        
        let idx = 0
        let leftChildIdx = idx * 2 + 1
        let rightChildIdx = idx * 2 + 2 
        
        while(true) {
            let smallerIdx = idx
            if(this.heap[leftChildIdx] && 
               this.heap[leftChildIdx][1] < this.heap[smallerIdx][1]) smallerIdx = leftChildIdx 
            if(this.heap[rightChildIdx] && 
               this.heap[rightChildIdx][1] < this.heap[smallerIdx][1]) smallerIdx = rightChildIdx 
            if(smallerIdx === idx) break
            
            this.swap(idx, smallerIdx)
            idx = smallerIdx
            leftChildIdx = idx * 2 + 1
            rightChildIdx = idx * 2 + 2 
        }
        return root
    }
}

function solution(jobs) { 
    jobs.sort((a,b) => a[0]===b[0] ? a[1]-b[1] : a[0]-b[0])
    let answer = 0
    let time = 0
    let i = 0
    const jobsQueue = new MinHeap()
    
    while(i < jobs.length || jobsQueue.getlength() > 0) {
        while(i < jobs.length && jobs[i][0] <= time) {
            jobsQueue.push(jobs[i])
            i++
        }
        if(jobsQueue.getlength() === 0 && i < jobs.length) {
            time = jobs[i][0]  
        } 
        if (jobsQueue.getlength() > 0) {
            const poppedJob = jobsQueue.del()
            time += poppedJob[1]
            answer += time - poppedJob[0]
        } 
    }
    return Math.floor(answer/jobs.length)
}