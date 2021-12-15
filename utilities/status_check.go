package utilities

import (
	"net/http"
)

func CheckToken(auth string) int {
	url := "https://discord.com/api/v9/users/@me/affinities/guilds"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return -1
	}
	req.Close = true
	req.Header.Set("authorization", auth)
	httpClient := http.DefaultClient
	resp, err := httpClient.Do(CommonHeaders(req))
	if err != nil {
		return -1
	}

	return resp.StatusCode

}
