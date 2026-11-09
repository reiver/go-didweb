package didweb

import (
	"net/url"

	"strings"
)

// Split splits a DID-Web path into a host and path.
//
// For example:
//
//	var did string = "did:web:example.com:users:joeblow"
//	
//	host, path, ok := didweb.Split(did)
//	
//	// host == "example.com"
//	// path == "users:joeblow"
//
// Note that "path" in this sense is a DID path (not a file system path).
// So the separator is ":" and not "/".
func Split(did string) (host string, path string, ok bool) {
	if !IsDIDWeb(did) {
		return "", "", false
	}

	const prefix string = "did:web:"

	str := did[len(prefix):]

	index := strings.Index(str, ":")

	escapedHost := str
	if 0 <= index {
		escapedHost = str[:index]
		path = str[1+index:]
		if "" == path {
			return "", "", false
		}
	}


	var err error
	//@TODO: should we only percent-decode the colon ("%3A"), or percent-decode everything?
	host, err = url.QueryUnescape(escapedHost)
	if nil != err {
		return "", "", false
	}


	return host, path, true
}
