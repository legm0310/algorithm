using System;
using System.Collections.Generic;
using System.Linq;

public class Solution {
    public long solution(long n) {
        List<int> list = new List<int>();

        while (n > 0)
        {
            list.Add((int)(n%10));
            n /= 10;
        }
        
        list.Sort((a, b) => b.CompareTo(a));
        return long.Parse(string.Join("", list));
    }
}