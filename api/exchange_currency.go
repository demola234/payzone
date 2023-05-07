package api

import (
	"io/ioutil"
	"log"
	"net/http"
)

func MakeRequest() {
	resp, err := http.Get("https://exchange-rates.abstractapi.com/v1/live/?api_key=09ad830c8bf643ea93b2349e5b9cdd98&base=USD&target=EUR")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}
