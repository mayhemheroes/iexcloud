package fuzzGetBytes

import "github.com/goinvest/iexcloud/v2"
import "context"

func mayhemit(bytes []byte) int {
    content := string(bytes)
    ctx := context.Background()
    test := iex.NewClient(content)
    test.GetBytes(ctx, content)
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}