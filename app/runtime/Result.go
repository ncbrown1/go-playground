package runtime

import (
    "log"
    "bytes"
    "encoding/gob"
)

type Result struct {
    Code   int `json:"code"`
    Output string `json:"output"`
}

func (r Result) success() bool {
    return r.Code == 0
}

func (r Result) PrintVerbose() {
    log.Println("output:")
    log.Println(r.Output)

    log.Printf("Exit status: %d\n", r.Code)
}

func (r Result) encode() ([]byte) {
    var buf bytes.Buffer
    enc := gob.NewEncoder(&buf)
    err := enc.Encode(r)
    if err != nil {
        log.Fatal("encode error:", err)
    }
    return buf.Bytes()
}

func decodeResults(b []byte) (Result, error) {
    var r Result
    dec := gob.NewDecoder(bytes.NewBuffer(b))
    err := dec.Decode(&r)
    return r, err
}

