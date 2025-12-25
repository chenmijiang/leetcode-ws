// LeetCode 题解模块
// 定义 Solution 结构体供各题解使用
pub struct Solution;

// 引入各题解模块
#[path = "solutions/36.valid-sudoku.rs"]
pub mod valid_sudoku;

#[path = "solutions/2110.number-of-smooth-descent-periods-of-a-stock.rs"]
pub mod number_of_smooth_descent_periods;

#[path = "solutions/54.spiral-matrix.rs"]
pub mod spiral_matrix;

#[path = "solutions/3074.apple-redistribution-into-boxes.rs"]
pub mod apple_redistribution_into_boxes;

#[path = "solutions/3075.maximize-happiness-of-selected-children.rs"]
pub mod maximize_happiness_of_selected_children;
