package whitelist

import (
	"encoding/json"
	"os"
	"strings"
)

// WhiteList contains a list of the whitelisted users and the path of the whitelist file
type WhiteList struct {
	filepath string
	List     []string
	Enabled  bool
}

// New reads the file provided and returns a new *WhiteList
func New(filepath string) (*WhiteList, error) {
	var v *WhiteList

	b, err := read(filepath)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		return v, err
	}
	v.filepath = filepath
	return v, err
}

// Add adds a new username to the whitelist
func (w *WhiteList) Add(username string) error {
	if w.Whitelisted(username) {
		return nil
	}
	w.List = append(w.List, username)
	b, err := json.MarshalIndent(w, "", "\t")
	if err != nil {
		return err
	}
	return os.WriteFile(w.filepath, b, 727)
}

// Remove removes the username provided from the whitelist
func (w *WhiteList) Remove(username string) error {
	if !w.Whitelisted(username) {
		return nil
	}
	w.List = remove(username, w.List)
	b, err := json.MarshalIndent(w, "", "\t")
	if err != nil {
		return err
	}
	return os.WriteFile(w.filepath, b, 727)
}

// Whitelisted returns a bool of if the player is whitelisted or not
func (w *WhiteList) Whitelisted(username string) bool {
	return username != "" && whitelisted(username, w.List)
}

// remove...
func remove(username string, list []string) []string {
	for i, u := range list {
		if strings.ToLower(u) == strings.ToLower(username) {
			return append(list[:i], list[i+1:]...)
		}
	}
	return list
}

// whitelisted...
func whitelisted(username string, list []string) bool {
	for _, u := range list {
		if strings.ToLower(u) == strings.ToLower(username) {
			return true
		}
	}
	return false
}

// read...
func read(filepath string) ([]byte, error) {
	if err := checkFile(filepath); err != nil {
		return nil, err
	}
	return os.ReadFile(filepath)
}

// checkfile...
func checkFile(filepath string) error {
	if _, err := os.Open(filepath); os.IsNotExist(err) {
		return marshalAndWrite(filepath)
	} else {
		return err
	}
}

// marshalAndWrite
func marshalAndWrite(filepath string) error {
	b, _ := json.MarshalIndent(WhiteList{List: []string{}}, "", "\t")
	err := os.WriteFile(filepath, b, 727)
	if err != nil {
		return err
	}
	return nil
}
