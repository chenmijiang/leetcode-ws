/**
给你三个长度为 n 的数组，分别描述 n 个优惠券的属性：code、businessLine 和 isActive。其中，第 i 个优惠券具有以下属性：

code[i]：一个 字符串，表示优惠券的标识符。
businessLine[i]：一个 字符串，表示优惠券所属的业务类别。
isActive[i]：一个 布尔值，表示优惠券是否当前有效。
当以下所有条件都满足时，优惠券被认为是 有效的 ：

code[i] 不能为空，并且仅由字母数字字符（a-z、A-Z、0-9）和下划线（_）组成。
businessLine[i] 必须是以下四个类别之一："electronics"、"grocery"、"pharmacy"、"restaurant"。
isActive[i] 为 true 。
返回所有 有效优惠券的标识符 组成的数组，按照以下规则排序：

先按照其 businessLine 的顺序排序："electronics"、"grocery"、"pharmacy"、"restaurant"。
在每个类别内，再按照 标识符的字典序（升序）排序。


示例 1：

输入： code = ["SAVE20","","PHARMA5","SAVE@20"], businessLine = ["restaurant","grocery","pharmacy","restaurant"], isActive = [true,true,true,true]

输出： ["PHARMA5","SAVE20"]

解释：

第一个优惠券有效。
第二个优惠券的标识符为空（无效）。
第三个优惠券有效。
第四个优惠券的标识符包含特殊字符 @（无效）。
示例 2：

输入： code = ["GROCERY15","ELECTRONICS_50","DISCOUNT10"], businessLine = ["grocery","electronics","invalid"], isActive = [false,true,true]

输出： ["ELECTRONICS_50"]

解释：

第一个优惠券无效，因为它未激活。
第二个优惠券有效。
第三个优惠券无效，因为其业务类别无效。


提示：

n == code.length == businessLine.length == isActive.length
1 <= n <= 100
0 <= code[i].length, businessLine[i].length <= 100
code[i] 和 businessLine[i] 由可打印的 ASCII 字符组成。
isActive[i] 的值为 true 或 false。©leetcode
**/

package main

import (
	"regexp"
	"sort"
)

func validateCoupons(code []string, businessLine []string, isActive []bool) []string {
	// 存储有效优惠券及其对应的业务类别
	type coupon struct {
		code         string
		businessLine string
	}

	validCoupons := []coupon{}
	for i := 0; i < len(code); i++ {
		if isActive[i] && isValidCode(code[i]) && isValidBusinessLine(businessLine[i]) {
			validCoupons = append(validCoupons, coupon{code[i], businessLine[i]})
		}
	}

	// 定义业务类别的优先级顺序
	businessOrder := map[string]int{
		"electronics": 0,
		"grocery":     1,
		"pharmacy":    2,
		"restaurant":  3,
	}

	sort.Slice(validCoupons, func(i, j int) bool {
		// 首先按业务类别排序
		if businessOrder[validCoupons[i].businessLine] != businessOrder[validCoupons[j].businessLine] {
			return businessOrder[validCoupons[i].businessLine] < businessOrder[validCoupons[j].businessLine]
		}
		// 在同一业务类别内按标识符字典序排序
		return validCoupons[i].code < validCoupons[j].code
	})

	// 提取排序后的优惠券代码
	result := make([]string, len(validCoupons))
	for i, coupon := range validCoupons {
		result[i] = coupon.code
	}

	return result
}

func isValidCode(code string) bool {
	return len(code) > 0 && regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString(code)
}

func isValidBusinessLine(businessLine string) bool {
	return businessLine == "electronics" || businessLine == "grocery" || businessLine == "pharmacy" || businessLine == "restaurant"
}

// func main() {
// 	code := []string{"1OFw", "0MvB"}
// 	businessLine := []string{"electronics", "pharmacy"}
// 	isActive := []bool{true, true}
// 	fmt.Println(validateCoupons(code, businessLine, isActive))
// }
