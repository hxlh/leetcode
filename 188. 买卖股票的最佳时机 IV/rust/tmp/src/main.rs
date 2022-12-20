fn main() {
    println!("Hello, world!");
}

struct Solution{}

// 上板的2变成k了
/**
 * 状态：
 * 0：不操作
 * 1：第一次买入
 * 2：第一次卖出
 * 3：....
 * 总结：奇数项买入，偶数项卖出，主要是把之前的手工代码变成自动
 */
impl Solution {
    pub fn max_profit(k: i32, prices: Vec<i32>) -> i32 {
        if k==0{
            return 0;
        }
        
        let mut dp:Vec<Vec<i32>>=vec![vec![0;((2*k)+1) as usize];prices.len()];
        // init 
        dp[0][0]=0;
        for j in 1..dp[0].len() {
            if j%2==0{  
                dp[0][j]=0;
                continue;
            }
            // 相当于当前买入卖出，消耗了前面的次数
            dp[0][j]=-prices[0];
        }

        for i in 1..prices.len(){
           for j in 1..dp[0].len(){
                if j %2==0{
                    dp[i][j]=std::cmp::max(dp[i-1][j],dp[i-1][j-1]+prices[i])
                }else{
                    dp[i][j]=std::cmp::max(dp[i-1][j],dp[i-1][j-1]-prices[i])
                }
           } 
        }

        return dp[prices.len()-1][(2*k) as usize];
    }
}