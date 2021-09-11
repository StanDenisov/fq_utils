package confclient

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/StanDenisov/fq_utils/confstruct"
)

type ResponseForConf struct {
	AppMode string
	AppName string
}

func GetConfig(appMode string, appName string) (confstruct.ConfStruct, error) {
	conf := confstruct.ConfStruct{}
	data := ResponseForConf{AppMode: appMode, AppName: appName}
	jsonResp, err := json.Marshal(data)
	if err != nil {
		panic("OMG")
	}
	resp, err := http.Post("https://localhost:13200", "application/json", bytes.NewBuffer(jsonResp))
	if err != nil {
		log.Fatalln(err)
		return conf, err
	}
	json.NewDecoder(resp.Body).Decode(conf)
	return conf, err
}
