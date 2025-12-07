package main

import (
    "fmt"
    "slices"
)

func main() {
var input []string
fmt.Println("uninitialize:", input, input == nil, len(input) == 0)

input = make([]string, 3)
fmt.Println("emp:", input, "len:", len(input), "cap:", cap(input))

input[0] = "i"
input[1] = "r"
input[2] = "a"
fmt.Println("set:", input)
fmt.Println("get:", input[2])
fmt.Println("len:", len(input))
input = append(input, "b")
input = append(input, "a", "t")
fmt.Println("Append --:", input)

c := make([]string, len(input))
copy(c, input)
fmt.Println("cpy:", c)
l := input[2:5]
fmt.Println("sl1:", l)
l = input[:5]
fmt.Println("sl2:", l)

t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)
t2 := []string{"g", "h", "i"}
if slices.Equal(t, t2) {
        fmt.Println("t == t2") }
twoDArray := make([][]int, 3)
for i := range 3 {
        innerLen := i + 1
        twoDArray[i] = make([]int, innerLen)
        for j := range innerLen {
            twoDArray[i][j] = i + j
        }
    }
fmt.Println("2d: ", twoDArray)
}
