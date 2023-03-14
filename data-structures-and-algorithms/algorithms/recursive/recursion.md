# Recursion

确定边界条件,否则无限循环
通过函数进行循环

模版：
```go
func recursion(level, param1, param2, ...) {
	// 递归终止条件
	if level > MAX_LEVEL {
		process_result
		return
	}
	
	// 处理
	process(level, data...)
	
	// 递归
	self.recursion(level+1, p1,...)
}
```
