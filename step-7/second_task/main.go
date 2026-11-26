package main

import (
	"strings"
	"time"
)

func QuizRunner(questions, answers []string, answerCh chan string) int {
	correct := 0

	for i := range questions {
		select {
		case answer := <-answerCh:
			if strings.EqualFold(strings.TrimSpace(answer), strings.TrimSpace(answers[i])) {
				correct++
			}
		case <-time.After(time.Second):
			continue
		}
	}

	return correct
}
