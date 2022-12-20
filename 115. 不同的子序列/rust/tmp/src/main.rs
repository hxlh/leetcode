use std::{iter::TrustedLen, slice::SliceIndex};

fn main() {
    println!("Hello, world!");
}

struct Solution{}
/**
 * dp[i][j]为以s[i-1]和t[j-1]结尾的 
 * 当两者不相等时：（继承上一个的数据）
 * dp[i][j]=dp[i-1][j]
 * 相等时：
 * dp[i][j]=dp[i-1][j-1]+dp[i-1][j]
 * （dp[i-1][j-1]，i前面的字符继承的情况）
 *  即s[i-1]==t[j-1]的最后的数据，如下面B->B相等的最后数据3
 * （dp[i-1][j]，i于上一次的情况）
 * 例如：
 *      A   B   C
 * A    1   0   0
 * B    1   1   0
 * B    1   2   0
 * B    1   3   0
 * C    1   3   3
 * 
 */
impl Solution {
    pub fn num_distinct(s: String, t: String) -> i32 {
        if s.len()==0{
            return 0;
        }
        let mut dp:Vec<Vec<i32>>=vec![vec![0;t.len()+1];s.len()+1];
        let ss=s.as_bytes();
        let tt=t.as_bytes();
        
        // init
        for i in 0..=s.len(){
            dp[i][0]=1;
        }

        for i in 1..=s.len(){
            for j in 1..=t.len(){
                if ss[i-1]==tt[j-1]{
                    dp[i][j]=dp[i-1][j]+dp[i-1][j-1];
                }else{
                    dp[i][j]=dp[i-1][j];
                }
            }
        }

        return dp[s.len()][t.len()];
    }
}

