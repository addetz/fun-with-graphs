package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

func main() {
    filename := "skip.txt"
    file, err := os.Create(filename)
    if err != nil {
        return
    }
    defer file.Close()

    file.WriteString("hola skip")

    data, err := ioutil.ReadFile(filename)
    //data, err := ioutil.ReadAll(file) - BROKEN BITCH
    if err != nil {
        return
    }
    fmt.Println("Data: ", string(data))
    fmt.Println("len: ",len(data))
}
