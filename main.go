package main

import (
    "fmt"
    "github.com/atotto/clipboard"
)

func main() {
    // Try reading text from the clipboard
    text, err := clipboard.ReadAll()
    if err != nil {
        fmt.Println("Error reading from clipboard:", err)
        return
    }

    if text == "" {
        fmt.Println("Clipboard is empty.")
    } else {
        fmt.Println("Clipboard content:", text)
    }
}
