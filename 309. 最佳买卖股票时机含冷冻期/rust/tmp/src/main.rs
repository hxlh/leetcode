fn main() {
    println!("{}", Solution::max_profit(vec![1]));
}
struct Solution {}

use std::cmp::max;
/**
 * 加了冷冻期后可分为四个状态
 * 1.当天持有股票（当天买入或之前持有或来自冷冻期）
 * 2.当天不持有股票（来自之前冷冻期之后或不操作）
 * 3.当天不持有股票（当天卖出）
 * 4.当天是冷冻期 （上一次卖出）
 *
 * 递推公式：
 * dp[i][0]=max(dp[i-1][0],max(dp[i-1][3],dp[i-1][1])-prices[i])
 * dp[i][1]=max(dp[i-1][1],dp[i-1][3])
 * dp[i][2]=dp[i-1][0]+prices[i]
 * dp[i][3]=dp[i-1][2]
*/
impl Solution {
    pub fn max_profit(prices: Vec<i32>) -> i32 {
        
        let mut dp: Vec<Vec<i32>> = vec![vec![0; 4]; prices.len()];
        dp[0][0] = -prices[0];

        for i in 1..prices.len() {
            dp[i][0] = max(dp[i - 1][0], max(dp[i - 1][1], dp[i - 1][3]) - prices[i]);
            dp[i][1] = max(dp[i - 1][1], dp[i - 1][3]);
            dp[i][2] = dp[i - 1][0] + prices[i];
            dp[i][3] = dp[i - 1][2];
        }
        // 返回值，取不持有股票的最大情况
       return max(dp[prices.len()-1][3],max(dp[prices.len()-1][1],dp[prices.len()-1][2]))
    }
}
