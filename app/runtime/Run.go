package runtime

import (
    "crypto/sha1"
    "io"
    "fmt"
)

func RunCode(code string) (*Result) {
    h := sha1.New()
    io.WriteString(h, code)
    sha := fmt.Sprintf("%x", h.Sum(nil))

    return &Result{
        0,
        sha,
        "",
    }
}