package dependencies

import (
	"fmt"
	"os"
)

func CreateStructure(basePath string, structure map[string]interface{}) error {
	for name, content := range structure {
		currentPath := fmt.Sprintf("%s/%s", basePath, name)
		switch v := content.(type) {
		case map[string]interface{}:
			// Jika item adalah folder, buat folder dan isi rekursif
			if err := os.MkdirAll(currentPath, 0755); err != nil {
				return err
			}
			if err := CreateStructure(currentPath, v); err != nil {
				return err
			}
		case []string:
			// Jika item adalah list of files
			if err := os.MkdirAll(currentPath, 0755); err != nil {
				return err
			}
			for _, file := range v {
				filePath := fmt.Sprintf("%s/%s", currentPath, file)
				if _, err := os.Create(filePath); err != nil {
					return err
				}
			}
		default:
			// Jika item adalah file
			if _, err := os.Create(currentPath); err != nil {
				return err
			}
		}
	}
	return nil
}
