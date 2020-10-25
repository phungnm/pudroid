package config
import (
    "encoding/json"
    "os"
    "fmt"
)
type Configuration struct {
    Port              int
    Base_URL   		string
      DBDriver string
	    DBUser string
	    DBPass string
	    DBName string

}
func init() {
	dir, _ := os.Getwd()
	fmt.Println(dir)
	file, err := os.Open(dir+"/config/config.dev.json")
	if err != nil {
	  fmt.Println("error:", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	Config = Configuration{}
	err = decoder.Decode(&Config)
	if err != nil {
	  fmt.Println("error:", err)
	}

}

var Config Configuration


