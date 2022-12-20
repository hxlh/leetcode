
fn main() {
    println!("{}",Solution::rob(vec![
        2,3
    ]));
}

struct Solution{}

impl Solution {
    // 因为成环，考虑0~n-1和1~n两种情况，取大的值
    pub fn rob(nums: Vec<i32>) -> i32 {
        if nums.len()==1{
            return nums[0];
        }
        if nums.len()==2{
            return std::cmp::max(nums[0],nums[1]);
        }
        let res1=Solution::rob_normal(&nums,0,nums.len()-2);
        let res2=Solution::rob_normal(&nums,1,nums.len()-1);
        return std::cmp::max(res1, res2);
    }
    fn rob_normal(nums: &Vec<i32>,start:usize,end:usize)->i32{
        let mut dp:Vec<i32>=vec![0;nums.len()];
        dp[start]=nums[start];
        dp[start+1]=std::cmp::max(dp[start],nums[start+1]);

        for i in (start+2)..=end{
            dp[i]=std::cmp::max(dp[i-1],dp[i-2]+nums[i]);
        }
        return dp[end];
    }
}
