fn main() {
   println!("{}",Solution::max_profit(vec![
    7
   ]))
}

struct Solution{}

/**
 * 本题在于可以多次买卖
 * dp[i][0]持有股票
 * dp[i][1]不持有
 * 
 * dp[i][0]发生变化，因为可以用之前的利润来买下一次的股票
 * dp[i][0]=max(dp[i-1][0],dp[i-1][1]-prices[i])
 */
impl Solution {
    pub fn max_profit(prices: Vec<i32>) -> i32 {
        let mut dp:Vec<Vec<i32>>=vec![vec![0;2];prices.len()];
        dp[0][0]=-prices[0];
        dp[0][1]=0;

        for i in 1..prices.len(){
            dp[i][0]=std::cmp::max(dp[i-1][0], dp[i-1][1]-prices[i]);
            dp[i][1]=std::cmp::max(dp[i-1][1],dp[i-1][0]+prices[i])
        }

        return dp[prices.len()-1][1];
    }
}