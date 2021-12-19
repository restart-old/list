package whitelist

import (
	"strings"
)

func (w *WhiteList) addCache(username string) {
	w.cMutex.Lock()
	defer w.cMutex.Unlock()
	w.List = append(w.List, username)
}
func (w *WhiteList) add(username string) error {
	w.addCache(username)
	return w.settings.Gophig.SetConf(w)
}

func (w *WhiteList) removeCache(username string) {
	w.cMutex.Lock()
	defer w.cMutex.Unlock()
	for i, u := range w.List {
		if strings.ToLower(u) == strings.ToLower(username) {
			w.List = append(w.List[:i], w.List[i+1:]...)
		}
	}
}
func (w *WhiteList) remove(username string) error {
	w.removeCache(username)
	return w.settings.Gophig.SetConf(w)
}

func (w *WhiteList) whitelistedCache(username string) bool {
	w.cMutex.Lock()
	defer w.cMutex.Unlock()
	for _, u := range w.List {
		if strings.ToLower(u) == strings.ToLower(username) {
			return true
		}
	}
	return false
}

func (w *WhiteList) whitelisted(username string) bool {
	err := w.settings.Gophig.GetConf(w)
	return err == nil && w.whitelistedCache(username)
}

func (w *WhiteList) closeCache() {
	w = nil
}
func (w *WhiteList) close() error {
	err := w.settings.Gophig.SetConf(w)
	w.closeCache()
	return err
}
