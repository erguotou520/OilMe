package main

import (
	// "crypto/tls"
	"encoding/json"
	"net/url"
	// "gopkg.in/gomail.v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Oil 油价
type Oil struct {
	Prov string `json:"prov"`
	P90  string `json:"p90"`
	P0   string `json:"p0"`
	P95  string `json:"p95"`
	P97  string `json:"p97"`
	P89  string `json:"p89"`
	P92  string `json:"p92"`
	CT   string `json:"ct"`
	P93  string `json:"p93"`
}

// ResBody res结构体
type ResBody struct {
	RetCode int8  `json:"ret_code"`
	List    []Oil `json:"list"`
}

// ResponseBody res返回的结构体
type ResponseBody struct {
	ResCode  int8    `json:"showapi_res_code"`
	ResError string  `json:"showapi_res_error"`
	ResBody  ResBody `json:"showapi_res_body"`
}

var (
	apiURL    string
	serverURL string
)

func request(url string) (*ResponseBody, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Request error: " + err.Error())
		return &ResponseBody{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Request parse error: " + err.Error())
		return &ResponseBody{}, err
	}
	var resBody ResponseBody
	err = json.Unmarshal(body, &resBody)
	if err != nil {
		log.Println("Request convert json error: " + err.Error())
		return &ResponseBody{}, err
	}
	return &resBody, nil
}

// func sendMail(content string) error {
// 	m := gomail.NewMessage()
// 	m.SetHeader("From", Config.Mail.Username)
// 	m.SetHeader("To", Config.Mail.To)
// 	m.SetHeader("Subject", Config.Mail.Subject)
// 	m.SetBody("text/html", content)
//
// 	d := gomail.NewPlainDialer(Config.Mail.Host, Config.Mail.Port, Config.Mail.Username, Config.Mail.Password)
// 	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
// 	// Send the email
// 	if err := d.DialAndSend(m); err != nil {
// 		debugLog.Println("Send main error: " + err.Error())
// 		return err
// 	}
// 	return nil
// }

func sendServer(url string, data map[string][]string) error {
	resp, err := http.PostForm(url, data)
	if err != nil {
		log.Println("Send wechat error: " + err.Error())
		return err
	}
	defer resp.Body.Close()
	return nil
}

func init() {
	apiURL = os.Getenv("API_URL")
	serverURL = os.Getenv("SERVER_URL")
	log.Print(apiURL)
	log.Print(serverURL)
}

func main() {
	// 判断环境变量
	if apiURL == "" || serverURL == "" {
		log.Fatal("Need api url and server url")
		os.Exit(-1)
	}
	resBody, err := request(apiURL)
	if err != nil {
		log.Print(err)
		return
	}
	// sendMail("今日油价:" + resBody.ResBody.List[0].P92 + "。")
	sendServer(serverURL, url.Values{
		"text": {"今日油价"},
		"desp": {resBody.ResBody.List[0].P92 + "元。"},
	})
}
