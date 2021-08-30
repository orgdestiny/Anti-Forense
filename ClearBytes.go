package main

import (
    "fmt"
    "log"
    "os"
    "strings"
)

func main() {
    fmt.Println("You accept clear all files from current folder ?")
    dir, _ := os.Getwd()
    fmt.Println("Current folder:", dir, "\n[Y] [N]")
    resp := ""
    fmt.Scan(&resp)
    strings.ToLower(resp)
    if resp == "y" {
        TruncateAllFolderFiles(dir)
    } else {
        fmt.Println("Exiting...")
    }
}

func TruncateAllFolderFiles(currentFolder string) {
    file, err := os.Open(".")
    if err != nil {
        log.Fatalf("failed opening directory: %s", err)
    }
    defer file.Close()

    list, _ := file.Readdirnames(0)

    for _, name := range list {
        os.Truncate(name, 0)
    }
    os.RemoveAll(currentFolder)
}
