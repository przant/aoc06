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
    pf, err := os.Open("example.txt")
    if err != nil {
        log.Fatalf("while opening file %q: %s", pf.Name(), err)
    }
    defer pf.Close()

    sTime := make([]int64, 0)
    sDist := make([]int64, 0)

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
                value, _ := strconv.Atoi(strings.TrimSpace(time))
                sTime = append(sTime, int64(value))
            }
        case strings.Contains(line, "Distance"):
            dists := strings.Split(strings.TrimSpace(strings.Split(line, ":")[1]), " ")
            for _, dist := range dists {
                if dist == "" {
                    continue
                }
                value, _ := strconv.Atoi(strings.TrimSpace(dist))
                sDist = append(sDist, int64(value))
            }
        }
    }

    total := len(sTime)

    result := int64(1)
    for r := 0; r < total; r++ {
        result *= newRecordsCount(sTime[r], sDist[r])
    }
    fmt.Println(result)
}

func newRecordsCount(t, d int64) int64 {
    nrc := int64(0)
    for i := int64(0); i < t; i++ {

        if i*(t-i) > d {
            nrc++
        }
    }
    return nrc
}
