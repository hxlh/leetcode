fn main() {
    println!("Hello, world!");
}
struct Solution{}

/**
 * dp[i]为以nums[i]结尾的最大子数组和
 * 两个来源
 * dp[i-1]+nums[i] 加上当前
 * nums[i] 从新开始
 * dp[i]=max(dp[i-1]+nums[i],nums[i])
 */
impl Solution {
    pub fn max_sub_array(nums: Vec<i32>) -> i32 {
        let mut dp:Vec<i32>=vec![0;nums.len()];
        let mut res=nums[0];
        dp[0]=nums[0];
        for i in 1..nums.len(){
            dp[i]=std::cmp::max(dp[i-1]+nums[i],nums[i]);
            if dp[i]>res{
                res=dp[i];
            }
        }
        return res;
    }
}