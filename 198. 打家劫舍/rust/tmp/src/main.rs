
fn main() {
    println!("{}",Solution::rob(vec![1,2,3,1]));
}

struct Solution{

}

/**dp[i]表示到该i个房子所获得的的最大数额
 * 第i个房子：
 * 状态转移方程dp[i]=max(偷dp[i-2]+该房子的金额,不偷dp[i-1])
 */
// impl Solution {
//     pub fn rob(nums: Vec<i32>) -> i32 {
//         if nums.len()==1{
//             return nums[0];
//         }
//         let n=nums.len();
//         let mut dp=vec![0;n];
//         dp[0]=nums[0];
//         dp[1]=max(dp[0],nums[1]);

//         for i in 2..nums.len(){
//             dp[i]=max(dp[i-1],nums[i]+dp[i-2]);
//         }

//         return dp[dp.len()-1];
//     }
// }

// 同理该第i个房子金额只与i-1和i-2的房子有关，因此可以压缩
impl Solution {
    pub fn rob(nums: Vec<i32>) -> i32 {
        if nums.len()==1{
            return nums[0];
        }
        let n=nums.len();
        let mut prev2=nums[0];
        let mut prev1=max(prev2,nums[1]);

        for i in 2..nums.len(){
            let tmp=max(prev1,nums[i]+prev2);
            prev2=prev1;
            prev1=tmp;
        }

        return prev1;
    }
}

fn max(a:i32,b:i32) -> i32{
    if a>b{
        a
    }else{
        b
    }
}