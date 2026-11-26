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

func Compare(name1, name2 string) (string, error) {
	mark1, err := fetchStudentMark(name1)
	if err != nil {
		return "", err
	}

	mark2, err := fetchStudentMark(name2)
	if err != nil {
		return "", err
	}

	switch {
	case mark1 > mark2:
		return ">", nil
	case mark2 > mark1:
		return "<", nil
	default:
		return "=", nil
	}

}
