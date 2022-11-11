fn main() {
    println!(
        "{}",
        Solution::maximal_square(vec![vec!['0'],])
    );
}

struct Solution {}

// dp[i][j]表示以右下角为标准的正方形的最大边长，dp[i][j]由其左上，上。左的顶点构成的正方形决定
// dp[i][j]=min(dp[i][j-1],dp[i-1][j],dp[i-1][j-1])
impl Solution {
    pub fn maximal_square(matrix: Vec<Vec<char>>) -> i32 {
        let m = matrix.len();
        let n = matrix[0].len();
        let mut max = 0;
        let mut dp = vec![vec![0; n]; m];
        for i in 0..m {
            if matrix[i][0] == '1' {
                dp[i][0] = 1;
                if dp[i][0] > max {
                    max = dp[i][0];
                }
            }
        }
        for j in 0..n {
            if matrix[0][j] == '1' {
                dp[0][j] = 1;
                if dp[0][j] > max {
                    max = dp[0][j];
                }
            }
            
        }

        // display(&dp);

       

        if m != 1 {
            for i in 1..m {
                for j in 1..n {
                    if matrix[i][j] == '1' {
                        dp[i][j] = min(dp[i][j - 1], dp[i - 1][j]);
                        dp[i][j] = min(dp[i - 1][j - 1], dp[i][j]);
                        dp[i][j] += 1;
                        if dp[i][j] > max {
                            max = dp[i][j];
                        }
                    }
                }
            }
        }
        // for i in 0..m {
        //     for j in 0..n {
        //         if dp[i][j] > max {
        //             max = dp[i][j];
        //         }
        //     }
        // }

        return max * max;
    }
}

fn min(a: i32, b: i32) -> i32 {
    if a < b {
        return a;
    }
    return b;
}

fn display(dp: &Vec<Vec<i32>>) {
    println!("display----------------------------------------------");
    for i in 0..dp.len() {
        println!("{:?}", dp[i]);
    }
}
