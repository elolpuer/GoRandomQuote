package config

import (
	"encoding/json"
	"fmt"
	"os"
)

//Setting data of config
type Setting struct {
	Host   string
	Port   string
	PgHost string
	PgPort string
	PgUser string
	PgPass string
	PgDB   string
}

var cfg Setting

//Config return config data
func Config() Setting {

	//Открыть файл
	file, err := os.Open("config/settings.cfg")

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Случилось какое-то дерьмо 1 ")
	}

	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Случилось какое-то дерьмо 2 ")
	}

	readByte := make([]byte, stat.Size())

	_, err = file.Read(readByte)

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Случилось какое-то дерьмо 3")
	}

	err = json.Unmarshal(readByte, &cfg)

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Случилось какое-то дерьмо 4")
	}

	return cfg
}
