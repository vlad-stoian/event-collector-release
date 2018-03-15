package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {

	pksTileVersion := os.Getenv("PKS_TILE_VERSION")
	aggregatorEndpoint := os.Getenv("AGGREGATOR_ENDPOINT")

	fmt.Println("PKS Tile Version: ", pksTileVersion)
	fmt.Println("Aggregator Endpoint: ", aggregatorEndpoint)

	sendEvent(aggregatorEndpoint, pksTileVersion)

	for {
		time.Sleep(time.Minute)
	}

}

func sendEvent(url, pksTileVersion string) {
	var jsonStr = []byte(fmt.Sprintf(`{"product_version":"%s"}`, pksTileVersion))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	fmt.Println("Request Headers:", req.Header)
	fmt.Println("Request Body:", req.Body)

	client := &http.Client{
		Timeout: time.Minute,
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Response Body:", string(body))
}
