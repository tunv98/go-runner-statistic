package main

import (
	"fmt"
	"strings"
)

type summary struct {
	name   string
	amount uint64
}

func (s summary) toString() string {
	return fmt.Sprintf("%s đóng góp %s \n", s.name, formatCurrency(s.amount))
}

type report struct {
	fromDate  string
	toDate    string
	summaries []summary
}

func (r report) toString() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Quỹ khuyến dục thông báo: từ %s đến %s\n", r.fromDate, r.toDate))
	sb.WriteString("Xin ghi nhận sự đóng góp của các a/c sau:\n")
	for _, s := range r.summaries {
		if s.amount == 0 {
			continue
		}
		sb.WriteString(fmt.Sprintf("%s", s.toString()))
	}
	sb.WriteString("Vì sự cống hiến trên, những thành viên sau sẽ ")
	return sb.String()
}
