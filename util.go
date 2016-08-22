
package detree

import "strings"

func RemoveQuotes(str string) string {
	if str[0] == '"' {
		return strings.Replace(str, string(str[0]), "", 2)
	} else if str[0] == '\'' {
		return strings.Replace(str, string(str[0]), "", 2)
	} else {
		return str
	}
}
