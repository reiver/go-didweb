package didweb

import (
	"strings"
	"net/url"
)

// PathSplit splits a DID-Web path into its parts.
//
// For example:
//
//	var path string = "users:joeblow"
//	
//	pathParts, ok := didweb.PathSplit(path)
//	
//	// pathParts == []string{"users", "joeblow"}
//
// Note that "path" in this sense is a DID path (not a file system path).
// So the separator is ":" and not "/".
func PathSplit(path string) ([]string, bool) {
	if "" == path {
		return nil, true
	}

	encodedParts := strings.Split(path, ":")

	parts := make([]string, len(encodedParts))
	for i:=0; i<len(encodedParts); i++ {
		part, err := url.QueryUnescape(encodedParts[i])
		if nil != err {
			return nil, false
		}

		parts[i] = part
	}

	return parts, true
}
