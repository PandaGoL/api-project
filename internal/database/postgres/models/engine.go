package models

import "fmt"

type ModelEngine struct {
	Columns        []string
	PlaceholderMap map[string]string
	Placeholders   []string
}

func NewModelEngine(columns []string) *ModelEngine {
	m := new(ModelEngine)
	m.Columns = columns
	m.PlaceholderMap = initPlaceholderMap(columns)
	m.Placeholders = initPlaceholder(columns)
	return m
}

func initPlaceholderMap(list []string) map[string]string {
	result := make(map[string]string)
	for i, v := range list {
		result[v] = fmt.Sprintf("$%d", i+1)
	}
	return result
}

func initPlaceholder(list []string) []string {
	result := make([]string, len(list))
	for i := range list {
		result[i] = fmt.Sprintf("$%d", i+1)
	}
	return result
}
