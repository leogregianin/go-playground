package main

import (
    "fmt"
    "golang.org/x/tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    var ret map[string]int
    ret = make(map[string]int)
    for _, v := range strings.Fields(s) {
        ret[v]++
    }
    fmt.Println(ret)
    return ret
}

func main() {
	wc.Test(WordCount)
}
