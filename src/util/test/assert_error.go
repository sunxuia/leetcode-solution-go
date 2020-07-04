package test

import (
	"fmt"
	"strings"
)

// AssertError 断言错误
type AssertError struct {
	handler *Handler
	message string
}

func (err *AssertError) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprint("Error while validate ", err.handler.caseName,
		"\nExpect  : ", err.handler.FullExpect,
		"\nActual  : ", err.handler.FullActual,
		"\nMessage : ", err.message))

	if len(err.handler.FieldName) > 0 {
		sb.WriteString(fmt.Sprintln("\nDetail  :",
			"\n    Field name : ", err.handler.FieldName,
			"\n    Expect     : ", err.handler.Expect,
			"\n    Actual     : ", err.handler.Actual,
		))
	}
	sb.WriteString("\n")
	return sb.String()
}
