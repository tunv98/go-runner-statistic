package main

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func makeSummaries(names [][]interface{}, amounts [][]interface{}) []summary {
	var summaries []summary
	for i := 0; i < len(names); i++ {
		var name string
		var amount uint64

		if len(names[i]) > 0 {
			name = fmt.Sprintf("%v", names[i][0])
		}

		if len(amounts[i]) > 0 {
			amountStr := fmt.Sprintf("%v", amounts[i][0])
			amount = formatAmount(amountStr)
		}
		summaries = append(summaries, summary{
			name:   name,
			amount: amount,
		})
	}
	sort.Slice(summaries, func(i, j int) bool {
		return summaries[i].amount > summaries[j].amount
	})
	return summaries
}

func formatAmount(amountStr string) uint64 {
	re := regexp.MustCompile(`[ .đ]`)
	amountStr = re.ReplaceAllString(amountStr, "")
	parsedAmount, err := strconv.ParseUint(amountStr, 10, 64)
	if err != nil {
		log.Printf("Failed to parse amount: %v", err)
		parsedAmount = 0
	}
	return parsedAmount
}

func formatCurrency(amount uint64) string {
	amountStr := strconv.FormatUint(amount, 10)
	var result strings.Builder
	n := len(amountStr)
	for i := n - 1; i >= 0; i-- {
		result.WriteByte(amountStr[i])
		if (n-i-1)%3 == 2 && i != 0 {
			result.WriteByte(',')
		}
	}
	output := result.String()
	var finalResult strings.Builder
	for i := len(output) - 1; i >= 0; i-- {
		finalResult.WriteByte(output[i])
	}
	return finalResult.String() + " đ"
}
