fn main() {
    println!("Hello, world!");
}
struct Solution{}
/**
 * 可转换为求两数组的最大子序列数量
 */
impl Solution {
    pub fn max_uncrossed_lines(nums1: Vec<i32>, nums2: Vec<i32>) -> i32 {
        let mut dp:Vec<Vec<i32>>=vec![vec![0;nums2.len()+1];nums1.len()+1];
        for i in 1..=nums1.len(){
            for j in 1..=nums2.len(){
                if nums1[i-1]==nums2[j-1]{
                    dp[i][j]=dp[i-1][j-1]+1;
                }else{
                    dp[i][j]=std::cmp::max(dp[i-1][j], dp[i][j-1]);
                }
            }
        }

        return dp[nums1.len()][nums2.len()];
    }
}