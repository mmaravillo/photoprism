package sanitize

import (
	"strings"

	"github.com/photoprism/photoprism/pkg/txt"
)

// Name sanitizes and capitalizes names.
func Name(name string) string {
	if name == "" {
		return ""
	}

	// Remove double quotes and other special characters.
	name = strings.Map(func(r rune) rune {
		switch r {
		case '"', '`', '~', '\\', '/', '*', '%', '&', '|', '+', '=', '$', '@', '!', '?', ':', ';', '<', '>', '{', '}':
			return -1
		}
		return r
	}, name)

	name = strings.TrimSpace(name)

	if name == "" {
		return ""
	}

	// Shorten and capitalize.
	return txt.Clip(txt.Title(name), txt.ClipDefault)
}
