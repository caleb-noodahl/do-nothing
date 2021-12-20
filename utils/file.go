package utils

import (
	"encoding/json"
	"os"
)

func WriteJson(path string, objs interface{}) error {
	data, err := json.Marshal(objs)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

func Write(path string, data []byte) error {
	return os.WriteFile(path, data, 0644)
}
