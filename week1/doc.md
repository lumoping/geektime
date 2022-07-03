# 作业

    我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

# 答案

**应该Warp。**

```go
package sql

var ErrNoRows = errors.New("sql: no rows in result set")
```

`sql.ErrNoRows`是一个`sentinel error`， 不添加上下文的话，难以定位错误发生位置。

为了方便定位位置，常见的方法有两种：
- 打印日志
  - 很明显，本题中的错误是需要返回给上层的，那这样我们就处理了这个错误两次，若上层也采取了同样的处理方法，那么就会有很多重叠的错误日志，对于排除问题的帮助很小。
- 添加上下文
  - `fmt.Errorf()`
      - 这种方式可以添加上下文、堆栈信息。上层在判断类型时，可以使用 `errors.Is(err, target error) bool` 、 `errors.As(err error, target any) bool`
        方法去判断、转换，非常方便

综上所述，应该wrap sql.ErrNoRows