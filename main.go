package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
)

type Record struct {
}

func main() {
    pf, err := os.Open("input.txt")
    if err != nil {
        log.Fatalf("while opening file %q: %s", pf.Name(), err)
    }
    defer pf.Close()

    timeStr := ""
    distStr := ""

    scnr := bufio.NewScanner(pf)

    for scnr.Scan() {
        line := scnr.Text()
        switch {
        case strings.Contains(line, "Time"):
            times := strings.Split(strings.TrimSpace(strings.Split(line, ":")[1]), " ")
            for _, time := range times {
                if time == "" {
                    continue
                }
                timeStr += time
            }
        case strings.Contains(line, "Distance"):

            dists := strings.Split(strings.TrimSpace(strings.Split(line, ":")[1]), " ")
            for _, dist := range dists {
                if dist == "" {
                    continue
                }
                distStr += dist
            }
        }
    }

    time, _ := strconv.Atoi(timeStr)
    dist, _ := strconv.Atoi(distStr)

    fmt.Printf("Time %d, Dist %d\n", time, dist)
    fmt.Println(newRecordsCount(int64(time), int64(dist)))

}

func newRecordsCount(t, d int64) uint64 {
    nrc := uint64(0)
    for i := int64(0); i < t; i++ {
        if i*(t-i) > d {
            nrc++
        }
    }
    return nrc
}

func concurrentRecordsCount(t, d int64, ch chan<- int64) {
    nrc := int64(0)
    for i := int64(0); i < t; i++ {
        if i*(t-i) > d {
            nrc++
        }
    }
    ch <- nrc
}
