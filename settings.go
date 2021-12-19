package whitelist

import "github.com/RestartFU/gophig"

type Settings struct {
	CacheOnly bool
	Gophig    *gophig.Gophig
}
