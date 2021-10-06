package flags

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/StanDenisov/fq_utils/confclient"
	"github.com/StanDenisov/fq_utils/confstruct"
)

func GetAppConfig() (confstruct.ConfStruct, error) {
	confStruct := confclient.ParseFlagsAndGetConfig()
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)
	jsonFile, err := os.Open(exPath + "/conf/conf.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &confStruct)
	return confStruct, nil
}
