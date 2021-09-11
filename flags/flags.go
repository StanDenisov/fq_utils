package flags

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/StanDenisov/fq_utils/confclient"
	"github.com/StanDenisov/fq_utils/confstruct"
)

func GetAppConfig(appMod string, appName string) (confstruct.ConfStruct, error) {
	if appMod != "" {
		confStruct, err := confclient.GetConfig(appMod, appName)
		if err != nil {
			panic("omg conf structure not recognize")
		}
		return confStruct, err
	}
	jsonFile, err := os.Open("../conf/conf.json")
	if err != nil {
		fmt.Println(appMod)
		fmt.Println(appName)
		panic("file not recongnized " + err.Error())
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var confStruct confstruct.ConfStruct
	json.Unmarshal(byteValue, &confStruct)
	return confStruct, nil
}
