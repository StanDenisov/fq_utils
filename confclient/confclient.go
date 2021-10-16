package confclient

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/StanDenisov/fq_utils/confstruct"
)

type ResponseForConf struct {
	AppMode string `json:"app_mode,omitempty"`
	AppName string `json:"app_name,omitempty"`
}

//ParseFlagsAndGetConfig - take 2 flags (1: -mod 2:-name)
//then send request to conf server and return realized ConfStruct
//if flag != "test" send response to conf server else parse local conf.json and send json
func ParseFlagsAndGetConfig() confstruct.ConfStruct {
	m, n := parse_flags()
	if m != "test" {
		confStruct, err := sendRequestToConfigServer(m, n)
		if err != nil {
			fmt.Println("Send request fail")
			os.Exit(1)
		}
		return confStruct
	}
	ex, err := os.Executable()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	exPath := filepath.Dir(ex)
	jsonFile, err := os.Open(exPath + "/conf/conf.json")
	if err != nil {
		fmt.Println("Cant read local conf.json")
		os.Exit(1)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var confStruct confstruct.ConfStruct
	json.Unmarshal(byteValue, &confStruct)
	return confStruct
}

func sendRequestToConfigServer(appMode string, appName string) (confstruct.ConfStruct, error) {
	conf := confstruct.ConfStruct{}
	data := ResponseForConf{AppMode: appMode, AppName: appName}
	jsonResp, err := json.Marshal(data)
	if err != nil {
		fmt.Println("marshaling json filed")
		fmt.Println(err)
		os.Exit(1)
	}
	resp, err := http.Post("https://localhost:13200", "application/json", bytes.NewBuffer(jsonResp))
	if err != nil {
		log.Fatalln(err)
		return conf, err
	}
	json.NewDecoder(resp.Body).Decode(&conf)
	return conf, err
}

func parse_flags() (string, string) {
	appMod := flag.String("mode", "test", "a app_mod for get config")
	appName := flag.String("name", "", "a app_name for get config")
	flag.Parse()
	return *appMod, *appName
}
