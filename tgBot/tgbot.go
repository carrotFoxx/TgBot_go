package tgBot

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

var token string /*= "1096211287:AAEq9jysmBjjbZRIJxK6zxQ2-j2pizgRNbk"*/
var lastUpdateId int =0;
func GetLastUpdateId()int{
	return lastUpdateId
}

func SetLastUpdateId(val int){
	lastUpdateId = val;
}

func SendMsg(chatId int , text string)  {
	var request string = "https://api.telegram.org/bot" + token + "/sendMessage?chat_id=" + strconv.Itoa(chatId) + "&text=" + text

	client := http.Client{
		Timeout: 60 * time.Second,
	}
	resp, err := client.Get(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	/*var result string;
	for true {
		bodyChan := make([]byte, 1024)
		n, err := resp.Body.Read(bodyChan)
		result += string(bodyChan[:n])
		if n == 0 || err != nil {
			break
		}

	}
	fmt.Println(result)*/
	
}

type updateArray []struct {
	UpdateId int `json:"update_id"`
	Message  struct {
		MessageId int `json:"message_id"`
		From      struct {
			Id    int    `json:"id"`
			IsBot bool   `json:"is_bot"`
			Name  string `json:"first_name"`
		} `json:"from"`

		Chat struct {
			Id   int    `json:"id"`
			Name string `json:"first_name"`
			Type string `json:"type"`
		} `json:"chat"`

		Date int    `json:"date"`
		Text string `json:"text"`
	} `json:"message"`
}


func tgUpdateRequest(response string) updateArray {

	type TgUpdateRequest struct {
		Stat   bool `json:"ok"`
		Result []struct {
			UpdateId int `json:"update_id"`
			Message  struct {
				MessageId int `json:"message_id"`
				From      struct {
					Id    int    `json:"id"`
					IsBot bool   `json:"is_bot"`
					Name  string `json:"first_name"`
				} `json:"from"`

				Chat struct {
					Id   int    `json:"id"`
					Name string `json:"first_name"`
					Type string `json:"type"`
				} `json:"chat"`

				Date int    `json:"date"`
				Text string `json:"text"`
			} `json:"message"`
		} `json:"result"`
	}

	tgStrct := TgUpdateRequest{}
	bytes := []byte(response)
	err := json.Unmarshal(bytes, &tgStrct)
	if err != nil {
		log.Println(err)
		return tgStrct.Result
	}
	/*for _, value := range tgStrct.Result {
		lastUpdateId = value.UpdateId
	}*/
	return tgStrct.Result
}


func GetUpdates(userToken string) updateArray {
	token = userToken
	var request string = "https://api.telegram.org/bot" + token + "/getUpdates"
	if lastUpdateId != 0{
		request = request + "?offset=" + strconv.Itoa(lastUpdateId);
	}
	client := http.Client{
		Timeout: 4 * time.Second,
	}
	resp, err := client.Get(request)
	if err != nil {
		fmt.Println(err)
		return tgUpdateRequest("err")
	}
	defer resp.Body.Close()

	var result string;
	for true {

		bodyChan := make([]byte, 1024)
		n, err := resp.Body.Read(bodyChan)
		result += string(bodyChan[:n])
		if n == 0 || err != nil {
			break
		}
	}
	fmt.Println(result)

	return tgUpdateRequest(result)
}
