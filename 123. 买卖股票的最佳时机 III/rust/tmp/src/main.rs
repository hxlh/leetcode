fn main() {
    println!("{}",Solution::max_profit(vec![
        3,3,5,0,0,3,1,4
    ]));
}

struct Solution{}

/**
 * 本题要求最多购买两次，因此有五种状态
 * 无操作
 * 第一次买入
 * 第一次卖出
 * 第二次买入
 * 第二次卖出
 * dp[i][0]=dp[i-1][0]
 * dp[i][1]=max(dp[i-1][1],dp[i-1][0]-prices[i])
 * dp[i][2]=max(dp[i-1][2],dp[i-1][1]+prices[i])
 * dp[i][3]=max(dp[i-1][3],dp[i-1][2]-prices[i])
 * dp[i][4]=max(dp[i-1][4],dp[i-1][3]+prices[i])
 */
impl Solution {
    pub fn max_profit(prices: Vec<i32>) -> i32 {
        let mut dp:Vec<Vec<i32>>=vec![vec![0;5];prices.len()];
        // init
        dp[0][0]=0;
        dp[0][1]=-prices[0];
        dp[0][2]=0;// 当天买入当天卖出，收入为0
        dp[0][3]=-prices[0];
        dp[0][4]=0;

        for i in 1..prices.len(){
            dp[i][0]=dp[i-1][0];
            dp[i][1]=std::cmp::max(dp[i-1][1],dp[i-1][0]-prices[i]);
            dp[i][2]=std::cmp::max(dp[i-1][2],dp[i-1][1]+prices[i]);
            dp[i][3]=std::cmp::max(dp[i-1][3],dp[i-1][2]-prices[i]);
            dp[i][4]=std::cmp::max(dp[i-1][4],dp[i-1][3]+prices[i]);
        }

        return dp[prices.len()-1][4];
    }
}