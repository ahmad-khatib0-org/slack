package version

import (
	"regexp"
	"strings"
)

var versionCommentRE = regexp.MustCompile(`^Minimum server version: (\d+\.\d+(?:\.\d+[\w-]*)?)$`)

func ExtractMinimumVersionFromComment(s string) string {
	lines := strings.Split(strings.TrimSpace(s), "\n")
	if len(lines) > 0 {
		lastLine := lines[len(lines)-1]
		if m := versionCommentRE.FindStringSubmatch(lastLine); len(m) >= 1 {
			return m[1]
		}
	}
	return ""
}
