package models

import (
	"fmt"
	"strings"
)

func getColumnList(columnList []string) string {
	return strings.Join(columnList, ", ")
}

func getColumnListWhithTableAbbreviation(abbr string, columnList []string) string {
	resultList := make([]string, len(columnList))
	for i := range columnList {
		resultList[i] = fmt.Sprintf("%s.%s", abbr, columnList[i])
	}
	return strings.Join(resultList, ", ")
}
func GetPlaceholders(input []interface{}) string {
	var placeholders []string
	for i := 0; i < len(input); i++ {
		placeholders = append(placeholders, fmt.Sprintf("$%d", i+1))
	}
	return strings.Join(placeholders, ", ")
}
