package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func fetchStudentMark(name string) (int, error) {
	resp, err := http.Get("http://localhost:8082/mark?name=" + name)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("Оценка студента %s не может быть получена! (HTTP: %d)", name, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(strings.TrimSpace(string(body)))
}

func Average(names []string) (int, error) {
	sum := 0

	for _, name := range names {
		mark, err := fetchStudentMark(name)
		if err != nil {
			return 0, err
		}

		sum += mark
	}

	average := sum / len(names)

	return average, nil
}
