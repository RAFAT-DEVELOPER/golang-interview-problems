package main

import (
	"fmt"
	"strings"
)

func concatPlusOperator(a, b string) string {
	return a + b
}

func concatFmtSprintf(a, b string) string {
	return fmt.Sprintf("%s%s", a, b)
}

func concatStringBuilder(a, b string) string {
	var builder strings.Builder
	builder.WriteString(a)
	builder.WriteString(b)
	return builder.String()
}
