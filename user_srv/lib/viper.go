package lib

import (
	"bytes"
	"io"
	"log"
	"os"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Viper struct {
  instance *viper.Viper
}

func NewViper () *Viper {
	return &Viper{}
}

func (libViper *Viper) Init() {
  vipper := viper.New()
	libViper.instance = vipper
	vipper.SetConfigType("toml")
} 

func (libViper *Viper) Unmarshal(path string, instance interface{}, name string) error {
  file, err := os.Open(path)

	if err != nil {
		log.Printf("[ERROR] lib.viper.Unmarshal.Open: %s", err.Error())
		return err
	}	

	data, err := io.ReadAll(file)

	if err != nil {
		log.Printf("[ERROR] lib.viper.Unmarshal.ReadAll: %s", err.Error())
		return err
	}

	libViper.instance.ReadConfig(bytes.NewBuffer(data))

	if err := libViper.instance.Unmarshal(instance); err != nil {
		log.Printf("[ERROR] lib.viper.Unmarshal.ReadAll: %s", err.Error())
		zap.S().Errorf("lib.viper.Unmarshal.Unmarshal: %s %+v", err, data)
		return err
	}

	log.Printf("[INFO] viper unmarshal %s", name)
	return nil
}