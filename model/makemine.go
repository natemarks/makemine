package model

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

// DataDir is the directory where all the files get written
const DataDir = "/etc/makemine"

// MyData contains all the fields we care about
type MyData struct {
	FullName  string `yaml:"fullName" json:"fullName"`
	LocalUser string `yaml:"localUser" json:"localUser"`
	Email     string `yaml:"email" json:"email"`
}

// MyDataFromURL given a string, it assumes the string is a url to a MyData json file
// it tries ot get eh file, unmarshall the json and return the variable
func MyDataFromURL(url string) (MyData, error) {

	var data MyData
	hClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return data, err
	}

	res, err := hClient.Do(req)
	if err != nil {
		return data, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}

	return data, err
}

// MyDataFromFilePath given a string, it assumes the string is a path to a local MyData json file.
// it tries to unmarshall the json and return the variable
func MyDataFromFilePath(filePath string) (MyData, error) {

	var data MyData
	file, _ := ioutil.ReadFile(filePath)

	err := json.Unmarshal(file, &data)

	return data, err
}

// MyDataFromInput prompts the user to manually enter the values for MyData and returns the variable
func MyDataFromInput() MyData {

	data := MyData{}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Full Name (ex. Firstly Lastly):")
	data.FullName, _ = reader.ReadString('\n')
	data.FullName = strings.TrimSuffix(data.FullName, "\n")

	fmt.Println("local computer user account(ex. flastly): ")
	data.LocalUser, _ = reader.ReadString('\n')
	data.LocalUser = strings.TrimSuffix(data.LocalUser, "\n")

	fmt.Println("Email address (ex. flastly@somedomain.com): ")
	data.Email, _ = reader.ReadString('\n')
	data.Email = strings.TrimSuffix(data.Email, "\n")

	return data
}

// ToJSOM write the variable data to a JSON file in the DataDir
func (data MyData) ToJSOM() error {
	filePath := path.Join(DataDir, "makemine.json")
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filePath, file, 0644)
	return err
}

// ToYaml write the variable data to a YAML file in the DataDir
func (data MyData) ToYaml() error {
	filePath := path.Join(DataDir, "makemine.yaml")
	file, err := yaml.Marshal(data)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filePath, file, 0644)
	return err
}

// ToSourceScript write the variable data to a script that can be sourced to export the
// MyData fields as variables
func (data MyData) ToSourceScript() error {
	filePath := path.Join(DataDir, "makemine.sh")
	file := fmt.Sprintf("export FULLNAME=\"%s\"\n", data.FullName)
	file = file + fmt.Sprintf("export LOCALUSER=\"%s\"\n", data.LocalUser)
	file = file + fmt.Sprintf("export EMAIL=\"%s\"\n", data.Email)
	err := ioutil.WriteFile(filePath, []byte(file), 0644)
	return err
}
