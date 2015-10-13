// Package conf loads toml formatted config files
// either using current working directory or
// revel basepath if available
package config

import (
	"errors"
	"github.com/BurntSushi/toml"
	"github.com/revel/revel"
	"log"
	"os"
	"strings"
)

const (
	defaultFile = `settings.toml`
	PS          = string(os.PathSeparator)
)

var (
	ErrConfigFileNotFound = errors.New(`Configuration file was not found.`)
	configFileName        = defaultFile
)

func Load(i interface{}, file ...string) error {
	var err error

	if len(file) > 0 && file[0] != `` {
		configFileName = file[0]
	}

	if err = fromCwd(i); err != nil {
		os.Chdir(revel.BasePath)
		log.Println(`load new `, revel.BasePath)
		err = fromCwd(i)
	}

	return err
}

func fromCwd(i interface{}) error {
	var (
		err error
		pwd string
	)

	if pwd, err = os.Getwd(); err != nil {
		return err
	}

	chunks := strings.Split(pwd, PS)

	var fileName string
	for i := len(chunks); i > 0; i-- {
		testFileName := strings.Join(chunks[0:i], PS) + PS + configFileName
		if _, err := os.Stat(testFileName); err == nil {
			fileName = testFileName
			break
		}
	}

	if fileName == `` {
		return ErrConfigFileNotFound
	}

	if _, err = toml.DecodeFile(fileName, i); err == nil {
		log.Printf("Loaded configuration file: %s\n", configFileName)
	}

	return err
}
