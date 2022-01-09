package randomCat

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	File string `json:"file"`
}

func Get() string {
	response, err := http.Get("http://aws.random.cat/meow")

	if err != nil {
		fmt.Println(err)
		return ""
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	return responseObject.File
}
