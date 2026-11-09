package didweb

import (
	"net/url"
	gopath "path"
)


// DocURI returns the URI used to resolve a DID-Web (ex: "did:web:example.com", "did:web:example.com:users:joeblow") into a DID document.
//
// DocURI returns an empty string if what it is given is not an DID-Web.
//
// For example:
//
//	var did string = "did:web:example.com"
//
//	uri := didweb.DocURI(did)
//
//	// uri == "https://example.com/.well-known/did.json"
//
// Also, for example:
//
//	var did string = "did:web:example.com:users:joeblow"
//
//	uri := didweb.DocURI(did)
//
//	// uri == "https://example.com/users/joeblow/did.json"
func DocURI(did string) string {
	if !IsDIDWeb(did) {
		return ""
	}

	host, path, ok := Split(did)
	if !ok {
		return ""
	}

	var uri url.URL
	uri.Scheme = "https"
	uri.Host = host //@TODO: should we make this canonical?
	switch path {
	case "":
		uri.Path = "/.well-known/did.json"
	default:
		pathParts, ok := PathSplit(path)
		if !ok {
			return ""
		}
		pathParts = append(pathParts, "did.json")

		uri.Path = gopath.Join(pathParts...)
	}

	return uri.String()
}
