package output

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/natemarks/makemine/input"
)

// global directory to save the makemine data
const globalDataDir = "/etc/makemine"

func makeDataDir() error {
	err := os.MkdirAll(globalDataDir, 0755)
	return err
}

func WriteData(imm input.MakeMineInput) error {
	err := makeDataDir()
	if err != nil {
		return err
	}
	err = ToJsonFile(imm, path.Join(globalDataDir, "makemine.json"))
	return err
}

func ToJsonFile(data input.MakeMineInput, oFile string) error {

	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	_ = ioutil.WriteFile(oFile, file, 0644)
	return err
}
