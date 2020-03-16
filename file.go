package utilities

import (
	"io"
	"os"
	"os/signal"
	"syscall"
)
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

func RegisterCloseFileSafety(files ...io.Writer) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		for _, f := range files {
			_ = f
		}
	}()
}
