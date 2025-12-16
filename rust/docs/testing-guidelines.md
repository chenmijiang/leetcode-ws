# LeetCode Rust 题解测试规范

本文档定义了 LeetCode Rust 题解的测试用例编写标准，确保所有题解都有完整、规范的测试覆盖。

## 1. 测试模块结构规范

### 1.1 基本结构

每个题解文件末尾应添加 `#[cfg(test)]` 测试模块：

```rust
#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_example_1() {
        // 测试代码
    }
}
```

### 1.2 必要的导入语句

**题解文件顶部**（在 `@lc code=start` 之前）：

```rust
use crate::Solution;
```

**测试模块内部**：

```rust
use super::*;  // 导入当前模块的所有公开项
```

### 1.3 测试函数命名约定

| 类型     | 命名格式              | 示例                               |
| -------- | --------------------- | ---------------------------------- |
| 题目示例 | `test_example_N`      | `test_example_1`, `test_example_2` |
| 边界情况 | `test_edge_case_描述` | `test_edge_case_empty_input`       |
| 特殊情况 | `test_special_描述`   | `test_special_all_same_values`     |

## 2. 测试用例覆盖要求

### 2.1 必须覆盖的用例

1. **题目示例**：所有题目描述中给出的示例必须有对应测试
2. **边界情况**：
   - 空输入（空数组、空字符串）
   - 单元素输入
   - 最大/最小值边界
   - 全相同元素
3. **异常情况**（如适用）：
   - 无解情况
   - 多解情况

### 2.2 覆盖检查清单

```
[ ] 所有题目示例已覆盖
[ ] 空输入/单元素边界已测试
[ ] 最大/最小值边界已测试
[ ] 特殊模式已考虑（全相同、递增、递减等）
```

## 3. 代码风格指南

### 3.1 测试函数注释规范

每个测试函数应包含中文注释说明：

```rust
#[test]
fn test_example_1() {
    // 示例1：[3,2,1,4] -> 7
    // 平滑下降阶段：[3], [2], [1], [4], [3,2], [2,1], [3,2,1]
    let prices = vec![3, 2, 1, 4];
    assert_eq!(Solution::get_descent_periods(prices), 7);
}
```

### 3.2 断言宏使用规范

| 宏           | 使用场景   | 示例                                 |
| ------------ | ---------- | ------------------------------------ |
| `assert_eq!` | 比较具体值 | `assert_eq!(result, expected)`       |
| `assert!`    | 布尔条件   | `assert!(Solution::is_valid(input))` |
| `assert_ne!` | 确保不相等 | `assert_ne!(result, wrong_value)`    |

**最佳实践**：

```rust
// ✅ 好：expected 在右边，便于阅读错误信息
assert_eq!(Solution::solve(input), expected);

// ✅ 好：对于布尔返回值，直接使用 assert_eq!
assert_eq!(Solution::is_valid(board), true);
assert_eq!(Solution::is_valid(board), false);
```

## 4. 实际代码示例

### 4.1 完整测试模块模板

```rust
#[cfg(test)]
mod tests {
    use super::*;

    // ========== 题目示例 ==========

    #[test]
    fn test_example_1() {
        // 示例1说明
        let input = vec![1, 2, 3];
        assert_eq!(Solution::solve(input), expected_output);
    }

    #[test]
    fn test_example_2() {
        // 示例2说明
    }

    // ========== 边界情况 ==========

    #[test]
    fn test_edge_case_empty() {
        // 边界情况：空输入
    }

    #[test]
    fn test_edge_case_single() {
        // 边界情况：单元素
    }

    // ========== 特殊情况 ==========

    #[test]
    fn test_special_all_same() {
        // 特殊情况：所有元素相同
    }
}
```

### 4.2 辅助函数示例

对于复杂输入类型，建议定义辅助函数：

```rust
#[cfg(test)]
mod tests {
    use super::*;

    /// 辅助函数：将字符串数组转换为 Vec<Vec<char>>（用于数独类题目）
    fn to_board(input: &[&[&str; 9]; 9]) -> Vec<Vec<char>> {
        input
            .iter()
            .map(|row| row.iter().map(|s| s.chars().next().unwrap()).collect())
            .collect()
    }

    /// 辅助函数：创建链表（用于链表类题目）
    fn to_list(values: &[i32]) -> Option<Box<ListNode>> {
        let mut head = None;
        for &val in values.iter().rev() {
            let mut node = ListNode::new(val);
            node.next = head;
            head = Some(Box::new(node));
        }
        head
    }

    #[test]
    fn test_with_helper() {
        let board = to_board(&[/* ... */]);
        assert_eq!(Solution::is_valid_sudoku(board), true);
    }
}
```

### 4.3 不同题型的测试示例

#### 数组类题目

```rust
#[test]
fn test_array_example() {
    // 两数之和
    let nums = vec![2, 7, 11, 15];
    let target = 9;
    assert_eq!(Solution::two_sum(nums, target), vec![0, 1]);
}
```

#### 字符串类题目

```rust
#[test]
fn test_string_example() {
    // 最长回文子串
    let s = String::from("babad");
    let result = Solution::longest_palindrome(s);
    // 可能有多个答案
    assert!(result == "bab" || result == "aba");
}
```

#### 树类题目

```rust
/// 辅助函数：从数组构建二叉树
fn to_tree(values: &[Option<i32>]) -> Option<Rc<RefCell<TreeNode>>> {
    // 实现略...
}

#[test]
fn test_tree_example() {
    let root = to_tree(&[Some(1), Some(2), Some(3)]);
    assert_eq!(Solution::max_depth(root), 2);
}
```

## 5. 验证流程

### 5.1 运行测试命令

```bash
# 运行所有测试
cargo test

# 运行特定模块的测试
cargo test valid_sudoku

# 运行特定测试函数
cargo test test_example_1

# 显示测试输出（包括 println!）
cargo test -- --nocapture

# 运行被忽略的测试
cargo test -- --ignored
```

### 5.2 测试输出解读

```
running 15 tests
test valid_sudoku::tests::test_example_1 ... ok
test valid_sudoku::tests::test_example_2 ... ok
...
test result: ok. 15 passed; 0 failed; 0 ignored
```

### 5.3 提交前检查清单

```
[ ] cargo test 全部通过
[ ] 无编译警告
[ ] 所有题目示例已覆盖
[ ] 边界情况已测试
[ ] 代码风格符合规范
```

### 5.4 新增题解工作流

1. 完成题解代码
2. 在文件顶部添加 `use crate::Solution;`
3. 在 `lib.rs` 中注册模块：
   ```rust
   #[path = "solutions/题号.题目名.rs"]
   pub mod 模块名;
   ```
4. 添加测试模块和用例
5. 运行 `cargo test` 验证
6. 提交代码

---

**维护说明**：本规范应随项目发展持续更新，如有新的最佳实践请及时补充。
