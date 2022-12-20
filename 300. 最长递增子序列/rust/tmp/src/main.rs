fn main() {
    println!("{}",Solution::length_of_lis(vec![
        1,3,6,7,9,4,10,5,6
    ]));
}
struct Solution{}

/**
 * dp[i]以nums[i]结尾的最长子序列的长度
 * if nums[i]>nums[j]
 * dp[i]=max(dp[j]+1,dp[i])
 */
use std::cmp::max;
impl Solution {
    pub fn length_of_lis(nums: Vec<i32>) -> i32 {
        let mut dp:Vec<i32>=vec![1;nums.len()];
        let mut res=1;
        for i in 0..nums.len(){
            for j in 0..i{
                if nums[j]<nums[i]{
                    dp[i]=max(dp[j]+1,dp[i]);
                    if dp[i]>res{
                        res=dp[i];
                    }
                }
            }
        }
        return res;
    }
}
