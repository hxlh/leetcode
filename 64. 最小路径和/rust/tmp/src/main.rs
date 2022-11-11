fn main() {
    println!(
        "{}",
        Solution::min_path_sum(vec![vec![1, 3, 1], vec![1, 5, 1], vec![4, 2, 1],])
    );
    println!(
        "{}",
        Solution::min_path_sum(vec![vec![1, 2, 3], vec![4, 5, 6],])
    );
}

struct Solution {}

// 显然dp[i][j]为走到(i,j)坐标的最小步数
// 每次只能向下或者向右移动一步。
// 因此dp[i][j]=min(dp[i-1][j],dp[i][j-1])+grid[i][j]
impl Solution {
    pub fn min_path_sum(grid: Vec<Vec<i32>>) -> i32 {
        let m=grid.len();
        let n=grid[0].len();
        let mut dp=vec![vec![0;n];m];

        for i in  0..m {
            for j in 0..n {
                if i==0 && j==0{
                    dp[i][j]=grid[i][j];
                }else if i==0{
                    dp[i][j]=dp[i][j-1]+grid[i][j];
                }else if j==0{
                    dp[i][j]=dp[i-1][j]+grid[i][j];
                }else{
                    dp[i][j]=min(dp[i-1][j],dp[i][j-1])+grid[i][j]
                }
            }
        }

        return dp[m-1][n-1];
    }
}


fn min(a: i32, b: i32) -> i32 {
    if a > b {
        return b;
    }
    return a;
}
