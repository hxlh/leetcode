fn main() {
    println!("{}",Solution::find_length_of_lcis(vec![
        2,2,2,2,2
    ]));
}
struct Solution{}

impl Solution {
    pub fn find_length_of_lcis(nums: Vec<i32>) -> i32 {
        let mut dp:Vec<i32>=vec![1;nums.len()];
        let mut res=1;
        for i in 1..nums.len(){
            if nums[i]>nums[i-1]{
                dp[i]=dp[i-1]+1;
                if dp[i]>res{
                    res=dp[i];
                }
            }
        }
        return res;
    }
}