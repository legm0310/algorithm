function solution(name, yearning, photo) {
    const table = name.reduce((a, c, i)=>{
        a[c]=yearning[i]
        return a
    },{})
    return photo.map((v)=> {
        return v.reduce((a,c)=>table[c] ? a+table[c] : a, 0)
    })
}