fn main() {
    println!("{}",Solution::min_cost_climbing_stairs(vec![10,15,20]));
}

struct Solution{
    
}

// dp[i]代表跳到i阶所花费的最小花费
// dp[i]由dp[i-1]和dp[i-2]决定，即dp[i]=min(dp[i-1]+cost[i-1],dp[i-2]+cost[i-2])
// 因为从0和1阶开始是不耗费体力，所以dp[0]=0,dp[1]=0
// 从2开始遍历
impl Solution {
    pub fn min_cost_climbing_stairs(cost: Vec<i32>) -> i32 {
        let mut dp=vec![0;cost.len()+1];
        for i in 2..cost.len()+1 {
            dp[i]=min(dp[i-1]+cost[i-1],dp[i-2]+cost[i-2]);
        }
        return dp[dp.len()-1];
    }

    
}

fn min(a:i32,b:i32)->i32 {
    if a<b{
        return a;
    }
    return b;
}