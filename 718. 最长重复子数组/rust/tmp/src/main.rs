fn main() {
    println!("{}",Solution::find_length(vec![
        2,3
    ], vec![
        1,4
    ]));
}

struct Solution{}
/**
 * dp[i][j]表示以nums1[i]和nums2[j]结尾的子数组最大长度
 * if nums[i]==nums[j] then
 * dp[i][j]=dp[i-1][j-1]+1
 */
impl Solution {
    pub fn find_length(nums1: Vec<i32>, nums2: Vec<i32>) -> i32 {
        let mut dp:Vec<Vec<i32>>=vec![vec![0;nums2.len()];nums1.len()];
        let mut res=0i32;
        for i in 0..nums2.len(){
            if nums1[0]==nums2[i]{
                dp[0][i]=1;
                if dp[0][i]>res{
                    res=dp[0][i];
                }
            }
        }

        for i in 0..nums1.len(){
            if nums1[i]==nums2[0]{
                dp[i][0]=1;
                if dp[i][0]>res{
                    res=dp[i][0];
                }
            }
        }

        for i in 1..nums1.len(){
            for j in 1..nums2.len(){
                if nums1[i]==nums2[j]{
                    dp[i][j]=dp[i-1][j-1]+1;
                    if dp[i][j]>res{
                        res=dp[i][j];
                    }
                }
            }
        }

        return res;
    }
}
