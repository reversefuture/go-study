# 命名
## getter
不要加Get，用uppercase即可
```go
owner := obj.Owner()
if owner != user {
    obj.SetOwner(user)
}
```
## interface
按照约定，只包含一个方法的接口应当以该方法的名称加上 - er 后缀来命名，如 Reader、Writer、 Formatter、CloseNotifier 等
请将字符串转换方法命名为 String 而非 ToString。

# 分号
规则是这样的：若在新行前的最后一个标记为标识符（包括 int 和 float64 这类的单词）、数值或字符串常量之类的基本字面或以下标记之一
```go
break continue fallthrough return ++ -- ) }
```
则词法分析将始终在该标记后面插入分号

分号也可在闭括号之前直接省略

警告：无论如何，你都不应将一个控制结构（if、for、switch 或 select）的左大括号放在下一行

# Redeclaration and reassignment
f, err := os.Open(name)
if err != nil {
    return err
}
d, err := f.Stat() // This duplication is legal: err is declared by the first statement, but only re-assigned in the second
if err != nil {
    f.Close()
    return err
}
codeUsing(f, d)

# For
It unifies for and while and there is no do-while. There are three forms, only one of which has semicolons.

// Like a C for
for init; condition; post { }

// Like a C while
for condition { }

// Like a C for(;;)
for { }

If you're looping over an array, slice, string, or map, or reading from a channel, a range clause can manage the loop.

for key, value := range oldMap {
    newMap[key] = value
}

If you only need the first item in the range (the key or index), drop the second:

for key := range m {
    if key.expired() {
        delete(m, key)
    }
}

If you only need the second item in the range (the value), use the blank identifier, an underscore, to discard the first:

sum := 0
for _, value := range array {
    sum += value
}

String: breaking out individual Unicode code points by parsing the UTF-8 with Erroneous encodings consume one byte and produce the replacement **rune** U+FFFD.
for pos, char := range "日本\x80語" { // \x80 is an illegal UTF-8 encoding
    fmt.Printf("character %#U starts at byte position %d\n", char, pos)
}

Finally, Go has no comma operator and ++ and -- are statements not expressions. Thus if you want to run multiple variables in a for you should use parallel assignment (although that precludes ++ and --).

// Reverse a
for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
    a[i], a[j] = a[j], a[i]
}

# Switch
Use with logical operators:
func unhex(c byte) byte {
    switch {
    case '0' <= c && c <= '9':
        return c - '0'
    case 'a' <= c && c <= 'f':
        return c - 'a' + 10
    case 'A' <= c && c <= 'F':
        return c - 'A' + 10
    }
    return 0
}

Use with cases presented in comma-separated lists.

func shouldEscape(c byte) bool {
    switch c {
    case ' ', '?', '&', '=', '#', '+', '%':
        return true
    }
    return false
}