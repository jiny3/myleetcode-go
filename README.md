# myleetcode-go

## package

- container(container/list 双向链表, container/ring 循环链表, **container/heap 堆**)

## tips

1. for 遍历时减少判断语句的计算量

> example: 将数组长度先存储在域内变量中

```go
tmp := []int{5, 2, 9, 13}
// for i := 0; i < len(tmp); i++ {	// 每次循环都会计算 len(tmp)
for i, n := 0, len(tmp); i < n; i++ {
  // ...
}
```

2. slice 不连续索引赋值

> 先扩容至该索引位置，保证访问不越界

```go
tmp := []int{}
visit_i := []int{5, 2, 9, 13}
for _, i := range visit_i {
    // 若当前索引 i 会越界则补足至该索引位置(或更长的位置)
    if len(tmp) <= i {
        tmp = append(tmp, make([]int, i-len(tmp)+1)...)
    }
    // 对 tmp[i] 进行操作
}
```

3. str -> int

```go
i, _ := strconv.Atoi(s)
```
