use std::collections::HashMap;

struct Solution {}

impl Solution {
    pub fn frequency_sort(s: String) -> String {
        let b = s.as_bytes();

        // 计算频次
        let mut f: HashMap<&u8, i32> = HashMap::new();
        for c in b {
            if let Some(r) = f.get(c) {
                f.insert(c, r.clone() + 1);
            } else {
                f.insert(c, 1);
            }
        }
        let mut buckets: Vec<(&u8, i32)> = Vec::with_capacity(f.len());
        for t in f {
            buckets.push(t.clone());
        }
        buckets.sort_by(|a, b| b.1.cmp(&a.1));
        let mut res = String::with_capacity(s.len());
        for t in buckets {
            for i in 0..t.1 {
                res.push((*t.0) as char);
            }
        }

        return res;
    }
}

fn main() {
    let s = Solution::frequency_sort(String::from("Aabb"));
    print!("{}\n", s);
}
