fn main() {
    println!("{}",Solution::is_subsequence("abc".to_string(),"ahbgdc".to_string()));
}

struct Solution{}
/**
 * 可转换为求最大子序列长度，并且长度要求==s的长度
 */
impl Solution {
    pub fn is_subsequence(s: String, t: String) -> bool {
        let mut dp:Vec<Vec<i32>>=vec![vec![0;t.len()+1];s.len()+1];
        let ss=s.as_bytes();
        let tt=t.as_bytes();

        for i in 1..=s.len(){
            for j in 1..=t.len(){
                if ss[i-1]==tt[j-1]{
                    dp[i][j]=dp[i-1][j-1]+1;

                }else{
                    dp[i][j]=std::cmp::max(dp[i-1][j], dp[i][j-1]);
                }
            }
        }

        if dp[s.len()][t.len()] as usize==s.len(){
            return true;
        }
        return false;
    }
}