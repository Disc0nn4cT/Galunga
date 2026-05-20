package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Server struct {
	Host       string   `json:"host"`
	Port       int      `json:"port"`
	Debug      bool     `json:"debug"`
	AllowedIPs []string `json:"allowed_ips"`
}

func ToJSON(v any) (string, error) {
	val := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}

	if val.Kind() != reflect.Struct {
		return "", fmt.Errorf("ToJSON підтримує лише структури, отримано: %s", val.Kind())
	}

	var sb strings.Builder
	sb.WriteString("{\n")

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fieldVal := val.Field(i)

		tag := field.Tag.Get("json")
		if tag == "" {
			tag = field.Name
		}

		sb.WriteString(fmt.Sprintf("\t\"%s\": ", tag))

		switch fieldVal.Kind() {
		case reflect.String:
			sb.WriteString(fmt.Sprintf("\"%s\"", fieldVal.String()))
		case reflect.Int:
			sb.WriteString(fmt.Sprintf("%d", fieldVal.Int()))
		case reflect.Bool:
			sb.WriteString(fmt.Sprintf("%t", fieldVal.Bool()))
		case reflect.Slice:

			sb.WriteString("[\n")
			for j := 0; j < fieldVal.Len(); j++ {
				elem := fieldVal.Index(j)
				sb.WriteString(fmt.Sprintf("\t\t\"%s\"", elem.String()))
				if j < fieldVal.Len()-1 {
					sb.WriteString(",")
				}
				sb.WriteString("\n")
			}
			sb.WriteString("\t]")
		default:
			sb.WriteString("null")
		}

		if i < val.NumField()-1 {
			sb.WriteString(",")
		}
		sb.WriteString("\n")
	}

	sb.WriteString("}")
	return sb.String(), nil
}

func main() {

	srv := Server{
		Host:       "localhost",
		Port:       8080,
		Debug:      true,
		AllowedIPs: []string{"192.168.1.1", "10.0.0.1"},
	}

	result, err := ToJSON(srv)
	if err != nil {
		fmt.Println("Помилка:", err)
		return
	}

	fmt.Println("=== Результат виконання вашої ToJSON ===")
	fmt.Println(result)
}
