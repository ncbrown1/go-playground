package runtime

import (
    "crypto/sha1"
    "io"
    "fmt"
    "os"
    "io/ioutil"
    "log"
)

func FmtCode(code string) (*Result) {
    // generate sha1 hash of code contents
    h := sha1.New()
    io.WriteString(h, code)
    sha := fmt.Sprintf("%x", h.Sum(nil))
    src_filename := fmt.Sprintf("/tmp/%s.go", sha)

    // if there was an error, compile and run the code again
    log.Println("Formatting code")
    err := ioutil.WriteFile(src_filename, []byte(code), 0644)
    defer os.Remove(src_filename)
    if err != nil {
        panic(err)
    }
    result := run("gofmt", src_filename)

    return &result
}
