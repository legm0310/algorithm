function solution(id_pw, db) {   
    let a = db.filter(el => el[0]===id_pw[0]).flat()
    return a.length===0?"fail":a[1]!==id_pw[1]?"wrong pw":"login"
}  
