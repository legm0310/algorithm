function solution(my_string) {
    return my_string.split("").filter(v => (v!=='a'&& v!=='e'&& v!=='i'&& v!=='o'&& v!=='u')).join(""); 
}