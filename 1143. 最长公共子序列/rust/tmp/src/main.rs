fn main() {
    println!("{}",Solution::longest_common_subsequence("def".to_string(),"abv".to_string()));
}

struct Solution{}

/**
 * dp[i][j]是以text1[i],text2[j]结尾的最长公共子序列
 * dp[i][j]有两种情况,text1[i]==(!=)text2[j]
 * ==: dp[i][j]=dp[i-1][j-1]+1;
 * !=: dp[i][j]=max(dp[i-1][j],dp[i][j-1])即可能来自i或j的前进
 */
use std::cmp::max;
impl Solution {
    pub fn longest_common_subsequence(text1: String, text2: String) -> i32 {
        // +1 避免需要单独初始化dp[0][j]和dp[i][0]
        let mut dp:Vec<Vec<i32>>=vec![vec![0;text2.len()+1];text1.len()+1];
        let str1=text1.as_bytes();
        let str2=text2.as_bytes();

        for i in 1..=text1.len(){
            for j in 1..=text2.len(){
                if str1[i-1]==str2[j-1]{
                    dp[i][j]=dp[i-1][j-1]+1;
                }else{
                    dp[i][j]=max(dp[i-1][j],dp[i][j-1]);
                }
            }
        }

        return dp[text1.len()][text2.len()];
    }
}
