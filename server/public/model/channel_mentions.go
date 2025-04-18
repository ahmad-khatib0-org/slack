package model

import (
	"regexp"
	"strings"
)

var channelMentionRegexp = regexp.MustCompile(`\B~[a-zA-Z0-9\-_]+`)

// What It Matches:
//   Valid matches:
//     ~username in "hello~username"
//     ~channel-1 in "msg~channel-1"
//     ~beta_test in "a~beta_test"
//   What it rejects:
//     ~admin at start of string ("~admin") ← Word boundary violation
//     ~user! ← Contains invalid character (!)
//     ~ alone ← Needs at least 1 valid char after tilde

func ChannelMentions(message string) []string {
	var names []string

	if strings.Contains(message, "~") {
		alreadyMentioned := make(map[string]bool)
		for _, match := range channelMentionRegexp.FindAllString(message, -1) {
			name := match[1:]
			if !alreadyMentioned[name] {
				names = append(names, name)
				alreadyMentioned[name] = true
			}
		}
	}

	return names
}
