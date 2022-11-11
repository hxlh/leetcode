fn main() {
    print!("{}",Solution::climb_stairs(12));
}


struct Solution{
    
}
// dp[i]=dp[i-1]+dp[i-2]
// dp[i]表示走到第i阶台阶所需方法树
// impl Solution {
//     pub fn climb_stairs(n: i32) -> i32 {
//         let mut dp = vec![0;(n+1) as usize];
//         dp[0]=1;
//         dp[1]=1;

//         let n=n as usize;
//         for i in 2..n + 1 {
//             dp[i] = dp[i - 1] + dp[i - 2];
//         }

//         return dp[n];
//     }
// }

// 压缩版，因dp[i]只与dp[i-1]和dp[i-2]有关
impl Solution {
    pub fn climb_stairs(n: i32) -> i32 {
        let mut prev1=1;
        let mut prev2=1;

        for _ in 2..n + 1 {
            let cur = prev1+prev2;
            prev2=prev1;
            prev1=cur;
        }

        return prev1;
    }
}