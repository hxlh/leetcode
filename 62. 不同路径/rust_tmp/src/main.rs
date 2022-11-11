fn main() {
    println!(
        "{}",
        Solution::unique_paths_with_obstacles(vec![vec![0, 0, 0], vec![0, 1, 0],vec! [0, 0, 0]])
    );
}

struct Solution {}
/*
dp[i][j]表示从(0,0)到(i,j)有多少种方法
由于机器人每次只能从左或从上方走到(i,j)
因此dp[i][j]=dp[i-1][j]+dp[i][j-1]
初始化顺序: dp[0][j]=1,dp[i][0]=1
遍历循序：从左到右
*/
impl Solution {
    pub fn unique_paths(m: i32, n: i32) -> i32 {
        let m = m as usize;
        let n = n as usize;

        let mut dp = vec![vec![0; n]; m];
        for i in 0..m {
            dp[i][0] = 1;
        }
        for j in 0..n {
            dp[0][j] = 1;
        }

        for i in 1..m {
            for j in 1..n {
                dp[i][j] = dp[i - 1][j] + dp[i][j - 1];
            }
        }

        return dp[m - 1][n - 1];
    }
}

// 63. 不同路径 II
impl Solution {
    pub fn unique_paths_with_obstacles(obstacle_grid: Vec<Vec<i32>>) -> i32 {
        if obstacle_grid[0][0] == 1 {
            return 0;
        }

        let m = obstacle_grid.len();
        let n = obstacle_grid[0].len();

        let mut dp = vec![vec![0; n]; m];
        for i in 0..m {
            if obstacle_grid[i][0] == 1 {
                break;
            }
            dp[i][0] = 1;
        }
        for j in 0..n {
            if obstacle_grid[0][j] == 1 {
                break;
            }
            dp[0][j] = 1;
        }

        for i in 1..m {
            for j in 1..n {
                if obstacle_grid[i][j] == 1 {
                    continue;
                }
                dp[i][j] = dp[i - 1][j] + dp[i][j - 1];
            }
        }
        for i in 0..dp.len() {
            println!("{:?}",dp[i]);
        }
        return dp[m - 1][n - 1];
    }
}
