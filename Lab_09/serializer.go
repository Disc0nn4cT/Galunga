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

// ==========================================
func ToYAML(v any) (string, error) {
	val := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}

	if val.Kind() != reflect.Struct {
		return "", fmt.Errorf("ToYAML підтримує лише структури")
	}

	var sb strings.Builder

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fieldVal := val.Field(i)

		tag := field.Tag.Get("json")
		if tag == "" {
			tag = field.Name
		}

		sb.WriteString(fmt.Sprintf("%s: ", tag))

		switch fieldVal.Kind() {
		case reflect.String:
			sb.WriteString(fmt.Sprintf("\"%s\"\n", fieldVal.String()))
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			sb.WriteString(fmt.Sprintf("%d\n", fieldVal.Int()))
		case reflect.Bool:
			sb.WriteString(fmt.Sprintf("%t\n", fieldVal.Bool()))
		case reflect.Slice:
			sb.WriteString("\n")
			for j := 0; j < fieldVal.Len(); j++ {
				elem := fieldVal.Index(j)
				if elem.Kind() == reflect.String {
					sb.WriteString(fmt.Sprintf("  - \"%s\"\n", elem.String()))
				} else {
					sb.WriteString(fmt.Sprintf("  - %v\n", elem.Interface()))
				}
			}
		default:
			sb.WriteString("null\n")
		}
	}

	return strings.TrimRight(sb.String(), "\n"), nil
}

// ==========================================
func ToJSON(v any) (string, error) {
	val := reflect.ValueOf(v)
	typ := reflect.TypeOf(v)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}

	if val.Kind() != reflect.Struct {
		return "", fmt.Errorf("ToJSON підтримує лише структури")
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
