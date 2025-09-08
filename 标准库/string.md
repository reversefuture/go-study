# `string` 是**不可变类型**，不能修改其中的字节。
- 只要你想“修改字符串”，就必须转为 `[]byte` 或者 `[]rune`
- `string` 可以作为 `map` 的 key，而 `[]byte` 不能（切片不可比较）。
- 处理中文/emoji 一定要用 rune


```go
s := "hello"
// s[0] = 'H'  // ❌ 编译错误

b := []byte(s)
b[0] = 'H' // ✅ 合法
s = string(b) // "Hello"
```
# 反引号
在 Go 语言中，**反引号（`` ` ``）** 是一种特殊的字符串字面量语法，称为 **原始字符串字面量（raw string literal）**。它与双引号定义的字符串不同，反引号内的内容**不会被转义**，几乎原样保留。

---

### ✅ 一、基本语法对比

| 写法 | 类型 | 是否转义 |
|------|------|----------|
| `"hello\nworld"` | 解释型字符串（interpreted string） | 是，`\n` 会被解释为换行 |
| `` `hello\nworld` `` | 原始字符串（raw string） | 否，`\n` 就是两个字符 `\` 和 `n` |

---

### ✅ 二、反引号的核心特点

1. **不进行转义处理**
   ```go
   fmt.Println("hello\nworld")
   // 输出：
   // hello
   // world

   fmt.Println(`hello\nworld`)
   // 输出：
   // hello\nworld
   ```

2. **可以跨行书写**
   ```go
   multiLine := `第一行
第二行
第三行`
   fmt.Println(multiLine)
   // 输出三行文本，换行符保留
   ```

3. **包含双引号也不需要转义**
   ```go
   json := `{"name": "Alice", "age": 30}`
   fmt.Println(json) // 直接输出 JSON，无需转义引号
   ```

4. **适合写正则表达式、SQL、模板、JSON 等**
   ```go
   regex := regexp.MustCompile(`\d+`) // 不用写成 "\\d+"
   sql := `
   SELECT * FROM users
   WHERE age > ?
   ORDER BY name
   `
   template := `
   <html>
   <body>
     <h1>{{.Title}}</h1>
   </body>
   </html>
   `
   ```

---

### ✅ 三、典型使用场景

#### 1. 正则表达式（避免双重转义）
```go
// ❌ 使用双引号：需要对 \ 转义
r1 := regexp.MustCompile("\\d+\\s+\\w+")

// ✅ 使用反引号：清晰简洁
r2 := regexp.MustCompile(`\d+\s+\w+`)
```

#### 2. SQL 查询语句
```go
query := `
SELECT id, name, email
FROM users
WHERE active = ?
  AND created_at > ?
ORDER BY name ASC
`
```

#### 3. JSON 或配置文本
```go
config := `{
  "host": "localhost",
  "port": 8080,
  "debug": true
}`
```

#### 4. HTML 或模板
```go
page := `
<!DOCTYPE html>
<html>
<head><title>测试页面</title></head>
<body>
  <p>欢迎访问！</p>
</body>
</html>
`
```

#### 5. 包含大量引号或反斜杠的内容
```go
path := `C:\Users\Go\Projects\test` // Windows 路径不再需要写成 "C:\\Users\\Go\\..."
```

---

### ✅ 四、注意事项

1. **不能嵌套反引号**
   ```go
   // ❌ 错误：不能在反引号字符串中包含反引号
   s := `This is a backtick: `` // 语法错误！

   // ✅ 解决方法：用双引号或其他方式
   s := "This is a backtick: `"
   ```

2. **开头和结尾的换行也会被包含**
> s := `

可以这样避免：
>   s := `hello

3. **反引号不能用于 rune（字符）**
   ```go
   // ❌ 错误
   var r rune = `a`

   // ✅ 正确
   var r rune = 'a'
   ```

---

### ✅ 五、总结

| 特性 | 反引号 `` `...` `` | 双引号 `"..."` |
|------|------------------|---------------|
| 转义 | ❌ 不转义 | ✅ 会转义（如 `\n`, `\"`） |
| 换行 | ✅ 支持多行 | ❌ 单行（除非用 `\n`） |
| 性能 | 相同 | 相同 |
| 适用场景 | 正则、SQL、JSON、模板等 | 一般字符串 |

---

### ✅ 推荐使用反引号的场景

- 正则表达式
- 多行文本
- 包含很多 `\` 或 `"` 的字符串
- 需要保持原样的脚本或配置
