package directmessage

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/DaKing01/dm-bot/utilities"
)

type UserInf struct {
	User   User     `json:"user"`
	Mutual []Guilds `json:"mutual_guilds"`
}
type User struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
}
type Guilds struct {
	ID string `json:"id"`
}

func UserInfo(token string, userid string) (UserInf, error) {
	url := "https://discord.com/api/v9/users/" + userid + "/profile?with_mutual_guilds=true"
	cookie, err := utilities.Cookies()
	if err != nil {
		return UserInf{}, err
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return UserInf{}, err
	}
	fingerprint, err := utilities.Fingerprint()
	if err != nil {
		return UserInf{}, err
	}
	req.Close = true
	req.Header.Set("Authorization", token)
	req.Header.Set("Cookie", cookie)
	req.Header.Set("x-fingerprint", fingerprint)
	req.Header.Set("host", "discord.com")

	client := http.DefaultClient

	resp, err := client.Do(utilities.CommonHeaders(req))
	if err != nil {
		return UserInf{}, err
	}

	body, err := utilities.ReadBody(*resp)
	if err != nil {
		return UserInf{}, err
	}

	if body == nil {

		return UserInf{}, fmt.Errorf("body is nil")
	}

	var info UserInf
	errx := json.Unmarshal(body, &info)
	if errx != nil {
		return UserInf{}, errx
	}
	return info, nil
}
