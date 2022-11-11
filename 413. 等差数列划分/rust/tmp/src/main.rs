fn main() {
    println!("{}",Solution::number_of_arithmetic_slices(vec![1,2,3,8,9,10]));
}
 
struct Solution{

}

// 因为等差数列随时停止，因此最后需要所有dp[i]相加
impl Solution {
    pub fn number_of_arithmetic_slices(nums: Vec<i32>) -> i32 {
        if nums.len()<3{
            return 0;
        }
        let mut dp=vec![0;nums.len()];
        for i in 2..nums.len(){
            if nums[i]-nums[i-1]==nums[i-1]-nums[i-2]{
                dp[i]=dp[i-1]+1;
            }
        }
        
        let mut res=0;
        for i in 2..dp.len(){
            res+=dp[i];
        }
        return res;
    }
}