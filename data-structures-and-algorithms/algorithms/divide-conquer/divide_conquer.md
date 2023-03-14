# divide_conquer

分治：分而治之，大问题拆分为子问题
分治 递归 区别：最后要将子结果组合
```go
// 分治

模版
func devide_conquer(problem, param1, param2 ...) {
	// 终止条件
	if problem == nil {
		return
	}

	// 准备数据
	data := prepare_data(problem)
	// 分解子问题
	subproblems := split_problem(problem, data)
	subresult1 := divide_conquer(subproblems[0], p1, ...)
	subresult2 := divide_conquer(subproblems[1], p2...)
	// 组合最终结果
	result := process_result(subresult1, subresult2...)
}
```
