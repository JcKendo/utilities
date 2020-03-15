package utilities

import "os"
import "github.com/BurntSushi/toml"

func LoadConfig(fileName string, config interface{}) error {
	HandlePrintf("Load config: " + fileName)
	_, err := os.Stat(fileName)
	if err != nil {
		return err
	}
	if _, err := toml.DecodeFile(fileName, config); err != nil {
		return err
	}
	return err
}

