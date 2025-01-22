package id

import "fmt"

func FmtId(first string, second string) string {
	return fmt.Sprintf("%s->%s", first, second)
}
