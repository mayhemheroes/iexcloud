package fuzz

import "github.com/goinvest/iexcloud/v2"

func mayhemit(bytes []byte) int {
    content := string(bytes)
    _ = iex.NewClient(content)
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}