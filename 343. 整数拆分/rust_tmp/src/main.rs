fn main() {
    println!("Hello, world!");
}

struct Solution {}


use std::cmp::max;
// dp[i]表示整数i拆分可得的最大乘积
// 遍历从1到i，将i拆分成j和i-j，此时dp[i]有两部分构成
// i-j不能拆分,dp[i]=j*(i-j)
// 可以拆分,dp[i]=j*dp[i-j]，这里dp[i-j]的作用是对(i-j)的拆分
// j已经在1->i的过程中被拆分，所以对i-j拆分
// 因为0和1不可拆分，因此dp[0/1]=0
impl Solution {
    pub fn integer_break(n: i32) -> i32 {
        let mut dp = vec![0_i32; (n + 1) as usize];
        dp[0] = 0;
        dp[1] = 0;
        dp[2] = 1;
        for i in 3..(n + 1) as usize {
            for j in 1..i {
                // 比较dp[i]是因为保存了j之前的最大值
                dp[i] = max(dp[i], max((dp[i - j] as usize) * j, j * (i - j)) as i32);
            }
        }
        return dp[n as usize];
    }
}
