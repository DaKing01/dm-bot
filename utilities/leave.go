package utilities

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/fatih/color"
)

func Leave(serverid string, token string) int {
	url := "https://discord.com/api/v9/users/@me/guilds/" + serverid
	json_data := "{\"lurking\":false}"
	req, err := http.NewRequest(http.MethodDelete, url, bytes.NewBuffer([]byte(json_data)))
	if err != nil {
		color.Red("Error: %s", err)
		return 0
	}
	req.Close = true
	cookie, err := Cookies()
	if err != nil {
		color.Red("Error: %s", err)
		return 0
	}
	req.Header.Set("authorization", token)
	req.Header.Set("Cookie", cookie)
	httpClient := http.DefaultClient
	resp, errq := httpClient.Do(CommonHeaders(req))
	if errq != nil {
		fmt.Println(errq)
		return 0
	}
	return resp.StatusCode
}
