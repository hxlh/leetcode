fn main() {
    Solution::update_matrix(vec![
        vec![1, 0, 1, 1, 0, 0, 1, 0, 0, 1],
        vec![0, 1, 1, 0, 1, 0, 1, 0, 1, 1],
        vec![0, 0, 1, 0, 1, 0, 0, 1, 0, 0],
        vec![1, 0, 1, 0, 1, 1, 1, 1, 1, 1],
        vec![0, 1, 0, 1, 1, 0, 0, 0, 0, 1],
        vec![0, 0, 1, 0, 1, 1, 1, 0, 1, 0],
        vec![0, 1, 0, 1, 0, 1, 0, 0, 1, 1],
        vec![1, 0, 0, 0, 1, 1, 1, 1, 0, 1],
        vec![1, 1, 1, 1, 1, 1, 1, 0, 1, 0],
        vec![1, 1, 1, 1, 0, 1, 0, 0, 1, 1],
    ]);
}
struct Solution {}
// 因为需要搜索四个方向，动态规划从左上搜到右下，再从右下搜到左上，实现四个方向的遍历
impl Solution {
    pub fn update_matrix(mat: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let m = mat.len();
        let n = mat[0].len();
        let mut dp = vec![vec![100000; n]; m];

        for i in 0..m {
            for j in 0..n {
                if mat[i][j] != 0 {
                    if j > 0 {
                        dp[i][j] = min(dp[i][j], dp[i][j - 1] + 1);
                    }
                    if i > 0 {
                        dp[i][j] = min(dp[i][j], dp[i - 1][j] + 1);
                    }
                } else {
                    dp[i][j] = 0;
                }
            }
        }
        for v in dp.iter() {
            println!("{:?}", v);
        }

        for i in 1..m + 1 {
            for j in 1..n + 1 {
                if mat[m - i][n - j] != 0 {
                    let i = m - i;
                    let j = n - j;
                    if j < n - 1 {
                        dp[i][j] = min(dp[i][j], dp[i][j + 1] + 1);
                    }
                    if i < m - 1 {
                        dp[i][j] = min(dp[i][j], dp[i + 1][j] + 1);
                    }
                }
            }
        }
        println!("");
        for v in dp.iter() {
            println!("{:?}", v);
        }
        return dp;
    }
}

fn min(a: i32, b: i32) -> i32 {
    if a < b {
        return a;
    }
    return b;
}

// [
// [1,0,1,1,0,0,1,0,0,1],
// [0,1,1,0,1,0,1,0,1,1],
// [0,0,1,0,1,0,0,1,0,0],
// [1,0,1,0,1,1,1,1,1,1],
// [0,1,0,1,1,0,0,0,0,1],
// [0,0,1,0,1,1,1,0,1,0],
// [0,1,0,1,0,1,0,0,1,1],
// [1,0,0,0,1,2,1,1,0,1],
// [2,1,1,1,1,2,1,0,1,0],
// [3,2,2,1,0,1,0,0,1,1]
// ]
