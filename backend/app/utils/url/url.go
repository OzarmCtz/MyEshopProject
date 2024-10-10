package url

import (
	"regexp"
	"strconv"
	"strings"
)

func GetResourceName(url string, pattern string) string {
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(url)
	path := ""

	if len(matches) >= 1 {
		parts := strings.Split(matches[1], "/")

		for i, part := range parts {
			if _, err := strconv.Atoi(part); err != nil {
				if i == 0 {
					path = part
				} else {
					path = path + "/" + part
				}
			}
		}
	}
	return path
}
