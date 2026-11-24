
package main

import (
    "bufio"
    "errors"
    "os"
    "time"
)

func ExtractLog(inputFileName string, start, end time.Time) ([]string, error) {
    file, err := os.Open(inputFileName)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var result []string
    scanner := bufio.NewScanner(file)

    layout := "02.01.2006"

    startDate, _ := time.Parse(layout, start.Format(layout))
    endDate, _ := time.Parse(layout, end.Format(layout))

    for scanner.Scan() {
        line := scanner.Text()

        if len(line) < 10 {
            continue
        }

        date, err := time.Parse(layout, line[:10])
        if err != nil {
            continue
        }

        if !date.Before(startDate) && !date.After(endDate) {
            result = append(result, line)
        }
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    if len(result) == 0 {
        return nil, errors.New("no logs found")
    }

    return result, nil
}
