## Discription
This module provides functions to reverse string and to get length of the string
## Usage 
# ReverseString(str string) string
```go
str := "Hello" 
revStr := ReverseString(str)
//revStr == "olleH"
```
For UTF-8
```go
str := "汉字"
revStr := ReverseString(str)
//revStr == "字汉"
```

# StringLength(str string) int 
```go
str := "Hello" 
length := StringLength(str)
//length == 5
```
For UTF-8
```go
str := "汉字"
length := StringLength(str)
//lenght == 2
```
