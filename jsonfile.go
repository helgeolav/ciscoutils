package ciscoutils

import (
	"encoding/json"
	"errors"
	"os"
)

// SaveJsonToFile save a struct to file as JSON with readable formatting. Returned is the error, if any.
func SaveJsonToFile(fn string, o interface{}) error {
	var err error
	if o != nil {
		b, err := json.MarshalIndent(o, "", " ")
		if err == nil {
			f, err := os.Create(fn)
			if err == nil {
				_, _ = f.Write(b)
				err = f.Close()
			}
		}
	} else {
		err = errors.New("No interface passed in")
	}
	return err
}
