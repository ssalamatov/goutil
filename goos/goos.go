package goos

import (
	"fmt"
	"io"
	"os"
)

var _ = os.CreateTemp


func ReadFileByChunk(tempfile string) []byte {
    f, err := os.Open(tempfile)
    if err != nil {
        fmt.Println("Error opening file: ", tempfile)
        return nil
    }
    defer f.Close()

    buf := make([]byte, 10)
    data := make([]byte, 100)
    for {
        n, err := f.Read(buf)
        if err == io.EOF {
            break
        }
        data = append(data, buf[:n]...)
    }
    return data
}

func PrintFile(temfile string) {
    data := ReadFileByChunk(temfile)
    if data == nil {
        return
    }
    for _, s := range data {
        fmt.Print(string(s))
    }
}


func GetEnv() string {
    err := os.Setenv("x", "y")
    if err != nil {
        panic("could not set env")
    }
    value, _ := os.LookupEnv("x")
    return value
}
