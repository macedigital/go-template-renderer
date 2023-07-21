package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"text/template"
)

func indexedList(itemTemplate string, length int64) (string, error) {
	if length < 0 {
		return "", fmt.Errorf("expecting non-negative integer, got #{length}")
	}

	result := make([]string, length)
	for i := int64(0); i < length; i++ {
		result[i] = strings.ReplaceAll(itemTemplate, "$", fmt.Sprint(i))
	}

	return strings.Join(result, ","), nil
}

func loadContext(filename string) (*interface{}, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var result interface{}
	err = json.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func render(content string, values *interface{}, writer io.Writer) {
	funcMap := template.FuncMap{
		"split":        strings.Split,
		"manifestHash": strings.Clone,
		"indexedList":  indexedList,
	}

	tmpl, err := template.New("tpl").Option("missingkey=error").Funcs(funcMap).Parse(content)
	if err != nil {
		log.Fatalf("cannot process template: #{err}")
	}

	err = tmpl.Execute(writer, values)
	if err != nil {
		log.Fatalf("cannot render template: #{err}")
	}
}

func main() {
	tmpl, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("cannot load template file: #{err}")
	}
	values, err := loadContext(os.Args[2])
	if err != nil {
		log.Fatalf("cannot load values file: #{err}")
	}

	render(string(tmpl), values, os.Stdout)
}
