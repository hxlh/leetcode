fn main() {
    println!("{}",Solution::max_profit(vec![
        1
    ],2));
}
struct Solution{}
// 状态转移方程多考虑手续费
/**
 * dp[i][0]持有
 * dp[i][1]不持有
 * 
 * dp[i][0]=max(dp[i-1][0],dp[i-1][1]-prices[i])
 * dp[i][1]=max(dp[i-1][1],dp[i-1][0]+prices[i]-fee)
 */
use std::cmp::max;
impl Solution {
    pub fn max_profit(prices: Vec<i32>, fee: i32) -> i32 {
        let mut dp:Vec<Vec<i32>>=vec![vec![0;2];prices.len()];
        dp[0][0]=-prices[0];
        dp[0][1]=0;

        for i in 1..prices.len(){
            dp[i][0]=max(dp[i-1][0],dp[i-1][1]-prices[i]);
            dp[i][1]=max(dp[i-1][1],dp[i-1][0]+prices[i]-fee);
        }
        
        return dp[prices.len()-1][1];
    }
}