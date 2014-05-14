package main

import (
    "bufio"
    "math/rand"
    "os"
    "time"
)

func filereader(Filename string) []string {
    file, _ := os.Open(Filename)
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines
}

func RandomQuote() string {
    quotes := filereader("quotes.txt")
    rand.Seed(time.Now().Unix())
    return quotes[rand.Intn(len(quotes))]
}

func main() {
    print(RandomQuote())
}
