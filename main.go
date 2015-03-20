package doudou

import (
	"fmt"
	"github.com/dmotylev/goproperties"
	"jcloud/doudou/db"
	"os"
	"path/filepath"
)

func main() {
	current_path, _ := os.Getwd()
	properties_file := filepath.Join(current_path, "../config.properties")
	db.Global_properties, _ = properties.Load(properties_file)
	db.InitAllDS(db.Global_properties)
	err := db.GenModel()
	if err != nil {
		fmt.Println(err)
	}
}
