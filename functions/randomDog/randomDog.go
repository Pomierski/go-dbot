package randomDog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func Get() string {
	response, err := http.Get("https://dog.ceo/api/breeds/image/random")

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

	return responseObject.Message
}
