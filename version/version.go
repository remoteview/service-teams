package version

import (
	"io/ioutil"
	"strings"
)

// GetVersion -
func GetVersion() string {
	v, err := ioutil.ReadFile("VERSION.txt")
	if err != nil {
		return "unset"
	}
	return strings.TrimSuffix(string(v), "\n")
}
