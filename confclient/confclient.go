package confclient

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/StanDenisov/fq_utils/confstruct"
)

type ResponseForConf struct {
	AppMode string
	AppName string
}

//ParseFlagsAndGetConfig - take 2 flags (1: -mod 2:-name)
//then send request to conf server and return realized ConfStruct
func ParseFlagsAndGetConfig() confstruct.ConfStruct {
	m, n := parse_flags()
	conf, err := sendRequestToConfigServer(m, n)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return conf
}

func sendRequestToConfigServer(appMode string, appName string) (confstruct.ConfStruct, error) {
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

func parse_flags() (string, string) {
	appMod := flag.String("mod", "test", "a app_mod for get config")
	appName := flag.String("name", "auth", "a app_name for get config")
	flag.Parse()
	return *appMod, *appName
}
