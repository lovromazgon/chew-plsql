package chew_plsql

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"sync"

	"github.com/lovromazgon/chew/funcmap"
)

var (
	plsqlf     *PlsqlFuncMap
	plsqlfInit sync.Once
)

func PlsqlNS() *PlsqlFuncMap {
	plsqlfInit.Do(func() { plsqlf = &PlsqlFuncMap{} })
	return plsqlf
}

func init() {
	funcmap.AddFunc(
		&funcmap.Func{
			Func: PlsqlNS,
			Doc: funcmap.FuncDoc{
				Name: "plsql",
				Text: "PL/SQL related functions",

				NestedFuncs: []funcmap.FuncDoc{
					funcmap.FuncDoc{
						Name: "ParameterType",
						Text: "Takes a parameter type (\"in\", \"out\", \"in out\") and outputs the uppercase version of the" +
							" parameter while padding it to 6 characters with spaces, so it's always aligned the same way.",
						Example: "{{ plsql.ParameterType \"in\" }}",
					},
					funcmap.FuncDoc{
						Name: "Comment",
						Text: "Takes a multi-line string and prepends dashes so it is considered as a comment in PL/SQL." +
							" It also indents the comment with spaces.",
						Example: "{{ plsql.Comment \"My comment can be\n multiple lines\n long\" 2 }}",
					},
					funcmap.FuncDoc{
						Name: "Separator",
						Text: "Outputs a separator for use between procedures and functions:\n" +
							"--------------------------------------------------------------------",
						Example: "{{ plsql.Separator }}",
					},
				},
			},
		})
}

type PlsqlFuncMap struct{}

func (*PlsqlFuncMap) ParameterType(paramType string) string {
	switch strings.ToUpper(paramType) {
	case "IN":
		return "IN    "
	case "OUT":
		return "   OUT"
	case "IN OUT":
		return "IN OUT"
	default:
		panic(fmt.Sprintf("Unknown parameter type: %s", paramType))
	}
}

func (*PlsqlFuncMap) Comment(comment string, indentSize int) string {
	scanner := bufio.NewScanner(strings.NewReader(comment))
	var buffer bytes.Buffer

	i := 0
	for scanner.Scan() {
		buffer.WriteString(strings.Repeat(" ", indentSize))
		buffer.WriteString("-- ")
		buffer.WriteString(scanner.Text())
		buffer.WriteString("\n")
		i++
	}

	return strings.TrimRight(buffer.String(), "\n")
}

func (*PlsqlFuncMap) Separator() string {
	return "--------------------------------------------------------------------"
}
