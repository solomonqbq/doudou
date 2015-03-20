package db

import (
	"fmt"
	"github.com/dmotylev/goproperties"
	"os"
	"path/filepath"
	"testing"
)

func TestDB(t *testing.T) {
	current_path, _ := os.Getwd()
	properties_file := filepath.Join(current_path, "../config.properties")
	Global_properties, _ = properties.Load(properties_file)
	InitAllDS(Global_properties)
	err := GenModel()
	if err != nil {
		fmt.Println(err)
	}
}
