package runtime

import "log"

type Result struct {
    Code   int `json:"code"`
    Output string `json:"output"`
    Error  string `json:"error"`
}

func (r Result) success() bool {
    return r.Code == 0
}

func (r Result) PrintVerbose() {
    log.Println("stdout:")
    log.Println(r.Output)

    log.Println("stderr:")
    log.Println(r.Error)

    log.Printf("Exit status: %d\n", r.Code)
}