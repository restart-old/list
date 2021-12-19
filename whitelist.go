package whitelist

import (
	"os"
	"sync"
)

type WhiteList struct {
	settings *Settings
	cMutex   sync.RWMutex

	List    []string
	Enabled bool
}

func New(settings *Settings) (*WhiteList, error) {
	whitelist := &WhiteList{settings: settings}

	if err := settings.Gophig.GetConf(whitelist); os.IsNotExist(err) {
		if err = settings.Gophig.SetConf(&WhiteList{List: make([]string, 0), Enabled: true}); err != nil {
			return whitelist, err
		}
		if err = settings.Gophig.GetConf(whitelist); err != nil {
			return whitelist, err
		}
	}
	return whitelist, nil
}

// Add adds a new username to the whitelist
func (w *WhiteList) Add(username string) error {
	if !w.Whitelisted(username) {
		if w.settings.CacheOnly {
			w.addCache(username)
			return nil
		}
		return w.add(username)
	}
	return nil
}

// Remove removes the username provided from the whitelist
func (w *WhiteList) Remove(username string) error {
	if w.Whitelisted(username) {
		if w.settings.CacheOnly {
			w.removeCache(username)
			return nil
		}
		return w.remove(username)
	}
	return nil
}

// Whitelisted returns a bool of if the player is whitelisted or not
func (w *WhiteList) Whitelisted(username string) bool {
	if username == "" {
		return false
	}
	if w.settings.CacheOnly {
		return w.whitelistedCache(username)
	}
	return w.whitelisted(username)
}

func (w *WhiteList) Close() error {
	return w.close()
}
