fn main() {
    println!("{}",Solution::max_profit(vec![
        7,1,5,3,6,4
    ]));
}

struct Solution{}

/**
 * dp[i][0]表示：第i天持有股票的现金数
 * dp[i][1]表示：第i天不持有股票的现金数
 * dp[i][0]有两个来源：
 * （1）来自前一天（保持现状）
 * （2）当天买入后剩余现金
 * dp[i][0]=max(dp[i-1][0],-prices[i])
 * dp[i][1]同样：
 * （1）之前的现金（保持现状）
 * （2）当天卖出的现金
 * dp[i][1]=max(dp[i-1][1],prices[i]+dp[i-1][0])
 * 遍历顺序：因为依赖之前的数据，所以从小到大
 * 
 */
impl Solution {
    pub fn max_profit(prices: Vec<i32>) -> i32 {
        let mut dp:Vec<Vec<i32>>=vec![vec![0;2];prices.len()];
        dp[0][0]=-prices[0];
        dp[0][1]=0;

        for i in 1..prices.len(){
            dp[i][0]=std::cmp::max(dp[i-1][0],-prices[i]);
            dp[i][1]=std::cmp::max(dp[i-1][1],prices[i]+dp[i-1][0])
        }

        return dp[dp.len()-1][1];
    }
}