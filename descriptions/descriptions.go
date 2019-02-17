package descriptions

import "strings"

// ShortDesc returns the first line of a long description.
func ShortDesc(longDesc string) string {
	return strings.Split(longDesc, "\n")[0]
}
