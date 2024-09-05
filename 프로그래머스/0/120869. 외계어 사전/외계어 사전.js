function solution(spell, dic) {
    return (dic.filter((e) => spell.every((v)=>e.includes(v))).length>0 ? 1 : 2)
}
