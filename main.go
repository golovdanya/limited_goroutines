package main
import (
    "bufio"
    "fmt"
    "os"
    "net/http"
    "io/ioutil"
    "strings"
    "time"
)


func worker(id int, urls chan string, results chan <- int, doneWorkers chan <- bool) {
    var data string
    for url := range urls {
        if (strings.HasPrefix(url, "http") ||  strings.HasPrefix(url, "https")) {
            resp, err := http.Get(url)
            if err != nil {
                fmt.Println(err)
            }
            defer  resp.Body.Close()
            body, err := ioutil.ReadAll(resp.Body)
            if err != nil {
                fmt.Println(err)
            }
            data = string(body)
        } else {
            body, err := ioutil.ReadFile(url)
            if err != nil {
                fmt.Println(err)
            }
            data = string(body)
        }
        number := strings.Count(data, "Go")
        fmt.Println("FINISH: worker", id, " job", url, "Number of Go is", number)
        results <- number
    }
    doneWorkers <- true
    return
}

func main() {
    maxWorkers := 3
    workers := 0

    urls := make(chan string)
    results := make(chan int)
    doneWorkers := make(chan bool)
    doneAnalyze := make(chan bool)

    go func() {
        totalCount := 0
        for res := range results {
            totalCount += res
        }
        fmt.Println("Total: ", totalCount)
        doneAnalyze <- true
    }()

    scanner := bufio.NewScanner(os.Stdin)
    start := time.Now()
    t := time.Now()
    for scanner.Scan() {
        if workers < maxWorkers {
            go worker(workers, urls, results, doneWorkers)
            workers += 1
        }
        url := scanner.Text()
        urls <- url
    }
    close(urls)
    for i := 0; i < workers; i++ {
        <-doneWorkers
    }
    elapsed := t.Sub(start)
    fmt.Println("Time = ", elapsed)
    close(results)
    <-doneAnalyze
}