package diskspace

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type ITelegram interface {
	SendMessage(param url.Values)
}

type Telegram struct {
	baseURL string
	token   string
	groupId string
}

func NewTelegram(baseUrl, token, groupId string) *Telegram {
	return &Telegram{
		baseURL: baseUrl,
		token:   token,
		groupId: groupId,
	}
}

func (telegram *Telegram) SendMessage(param url.Values) {
	client := &http.Client{}

	param.Set("chat_id", telegram.groupId)

	request, err := http.NewRequest("POST", fmt.Sprintf("%s%s/%s", telegram.baseURL, telegram.token, "sendMessage"), bytes.NewBufferString(param.Encode()))
	if err != nil {
		log.Println(err)
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
	}

	defer response.Body.Close()

	b, _ := io.ReadAll(response.Body)

	if response.StatusCode == 200 || response.StatusCode == 201 {
		log.Println("Data success sent", string(b))
	} else {
		log.Println("Telegram", string(b))
	}

	b = nil
}
