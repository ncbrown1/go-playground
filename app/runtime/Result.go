package runtime

import "log"

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