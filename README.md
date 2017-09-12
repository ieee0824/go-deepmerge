# go-deepmerge
Merge the enumerable attributes of two interface deeply.

# example

```go
a := map[string]interface{}{
    "foo": "bar",
    "hoge": "huga",
    "array": []interface{}{1, 2, 3},
}

b := map[string]interface{}{
    "john": "doe",
    "hoge": "huga",
    "array": []interface{}{1, 2, "fizz", 4, "buzz"},
}

c, _ := Merge(a, b)

c => 
map[string]interface{}{
    "foo": "bar",
    "john": "doe",
    "hoge": "huga",
    "array": []interface{}{1, 2, 3, 1, 2, "fizz", 4, "buzz"},
}
```