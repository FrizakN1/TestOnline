package setting

import (
	"encoding/json"
	"knocker/utils"
)

type Setting struct {
	Address string
	Port    string
	DbHost  string
	DbPort  string
	DbUser  string
	DbPass  string
	DbName  string
}

var settings Setting

func LoadSetting(filename string) *Setting {
	bytes, e := utils.LoadFile(filename)
	if e != nil {
		utils.Logger.Println(e)
		return nil
	}
	e = json.Unmarshal(bytes, &settings)
	if e != nil {
		utils.Logger.Println(e)
		return nil
	}
	return &settings
}
