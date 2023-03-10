package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	jsonStr, err := getJson("https://www.jma.go.jp/bosai/forecast/data/overview_forecast/140000.json")
	if err != nil {
		panic(err)
	}

	wetherReport, err := parseJson(jsonStr)
	if err != nil {
		panic(err)
	}
	fmt.Println(wetherReport)
}

// JSONファイルをインターネットから取得して、構造体にマッピングする
// string, nil で返す
func getJson(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "" ,err
	}
	defer res.Body.Close()

	byteArray, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(byteArray), nil
}

// JSONから取得したデータを格納する構造体
type WetherReport struct {
	PublishingOffice string `json:"publishingOffice"`
	ReportDatetime string `json:"reportDatetime"`
	TargetArea string `json:"targetArea"`
	Text string `json:"text"`
}

func parseJson(jsonStr string) (WetherReport, error) {
	var wetherReport WetherReport
	err := json.Unmarshal([]byte(jsonStr), &wetherReport)
	return wetherReport, err
}