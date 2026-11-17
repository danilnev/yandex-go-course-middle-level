package main

import (
	"fmt"
	"strings"
	"time"
)

type Ticket struct {
	Ticket string
	User   string
	Status string
	Date   time.Time
}

func StatusIsRight(status *string) bool {
	return *status == "Готово" || *status == "В работе" || *status == "Не будет сделано"
}

func GetTasks(text string, user *string, status *string) []Ticket {
	lines := strings.Split(text, "\n")

	tickets := &[]Ticket{}

	if status != nil && !StatusIsRight(status) {
		return *tickets
	}

	for _, line := range lines {
		splitted := strings.Split(line, "_")
		if len(splitted) != 4 {
			continue
		}
		line_ticket, line_user, line_status := splitted[0], splitted[1], splitted[2]
		line_date, err := time.Parse("2006-01-02", splitted[3])
		if err == nil {
			if (user == nil || *user == line_user) && (status == nil || *status == line_status) {
				*tickets = append(*tickets, Ticket{
					Ticket: line_ticket,
					User:   line_user,
					Status: line_status,
					Date:   line_date,
				})
			}
		}
	}

	return *tickets
}

func main() {
	chatHistory := `
TICKET-12345_Паша Попов_Готово_2024-01-01
TICKET-12346_Иван Иванов_В работе_2024-01-02
TICKET-12347_Анна Смирнова_Не будет сделано_2024-01-03
TICKET-12348_Паша Попов_В работе 1_2024-01-04
`
	status := "В работе 1"
	result := GetTasks(chatHistory, nil, &status)
	fmt.Println(result)
}
