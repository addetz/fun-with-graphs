package main

import (
    "fmt"
    "strings"
    "sync"
)

/*
Your task is to design a following string processing pipeline:
put on queue -> strip whitespace -> change to UPPER -> reverse -> print
 */
func main() {
    done := make(chan struct{})
    defer close(done)
    in := gen("real", "   housewives   ", " of", "beverly", "hills   ")
    stripped1 := strip(done, in)
    stripped2 := strip(done, in)
    stripped := merge(done, stripped1, stripped2)
    uppered := upper(done, stripped)
    reversed := reverse(done, uppered)
    for r := range reversed {
        fmt.Println(r)
    }
}

func merge(done <-chan struct{}, cs ...<-chan string) <-chan string {
    var wg sync.WaitGroup
    out := make(chan string)

    // Start an output goroutine for each input channel in cs.  output
    // copies values from c to out until c or done is closed, then calls
    // wg.Done.
    output := func(c <-chan string) {
        defer wg.Done()
        for n := range c {
            select {
            case out <- n:
            case <-done:
                return
            }
        }
    }
    wg.Add(len(cs))
    for _, c := range cs {
        go output(c)
    }

    // Start a goroutine to close out once all the output goroutines are
    // done.  This must start after the wg.Add call.
    go func() {
        wg.Wait()
        close(out)
    }()
    return out
}

func gen(strings ...string) <-chan string {
    out := make(chan string)
    go func() {
        for _, s := range strings {
            out <- s
        }
        close(out)
    }()
    return out
}

func strip(done <-chan struct{}, in <-chan string) <-chan string {
    out := make(chan string)
    go func() {
        defer close(out)
        for s := range in {
            select {
            case out <- strings.Trim(s, " "):
            case <-done:
                return
            }
        }
    }()
    return out
}

func upper(done <-chan struct{}, in <-chan string) <-chan string {
    out := make(chan string)
    go func() {
        defer close(out)
        for s := range in {
            select {
            case out <- strings.ToUpper(s):
            case <-done:
                return
            }
        }
    }()
    return out
}

func reverse(done <-chan struct{}, in <-chan string) <-chan string {
    out := make(chan string)
    go func() {
        defer close(out)
        for s := range in {
            select {
            case out <- rev(s):
            case <-done:
                return
            }
        }
    }()
    return out
}

func rev(s string) string{
    var reverse string
    for i := len(s)-1; i >= 0; i-- {
        reverse += string(s[i])
    }
    return reverse
}

