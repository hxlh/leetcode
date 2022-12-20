use std::ops::Index;

fn main() {
    let res=input("leetcode",vec!["leet","code"]);
    println!("{}",Solution::word_break(res.0,res.1));
}

fn input(s:&str,v:Vec<&str>)->(String,Vec<String>){
    let mut rv:Vec<String>=Vec::with_capacity(v.len());
    for i in v.iter(){
        rv.push(i.to_string());
    }
    return (s.to_string(),rv);
}

struct Solution{

}
/**思路:
 * 本题是多重背包，背包大小为单词s的长度，物品是单词表的单词
 * 同时因为word的顺序有要求，所以是排列问题
 * dp[i]->true/false表示单词长度为i时，是否可以拼接成s
 * 状态转移方程dp[i]=前面0..j长度的字符串是否满足，以及j..i是否能找到符合的word
 * 初始化dp[0]=true;
 */
impl Solution {
    pub fn word_break(s: String, word_dict: Vec<String>) -> bool {
        let mut dp=vec![false;s.len()+1];
        dp[0]=true;
        for i in 1..s.len()+1{
            for j in 0..i {
                if dp[j]{
                    // 可以换成map实现O(1)查找
                    for w in word_dict.iter(){
                        if w==&s[j..i]{
                            dp[i]=true;
                            break;
                        }
                    }
                }   
            }
        }

        return dp[s.len()];
    }
}