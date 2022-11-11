fn main() {
    println!("{}",Solution::fib(3));
}

struct Solution{

}
// dp[i]第i个数的斐波拉契值
// dp[i]=dp[i-1]+dp[i-2]
// 初始化dp[0]=0,dp[1]=1
// 遍历顺序：从小到大
impl Solution {
    pub fn fib(n: i32) -> i32 {
        if n==0{
            return 0;
        }
        let mut prev1=1;
        let mut prev2=0;
        for _ in 2..n+1{
            let tmp=prev1+prev2;
            prev2=prev1;
            prev1=tmp;
        }
        return prev1;
    }
}