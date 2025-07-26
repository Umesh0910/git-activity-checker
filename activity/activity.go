package activity

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type UserActResponse struct {
	UserName    string `json:"userName"`
	Activity    string `json:"activity"`
	Repository  string `json:"repository"`
	PerformedAt string `json:"performedAt"`
}

func (u *UserActResponse) GetUserActivity(user string) (*UserActResponse, error){
	log.Println("Getting user activity")
	resp, err := http.Get("https://api.github.com/users/" + user + "/events")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil{
		return nil, err
	}
	fmt.Printf("Response: %s", body)
	return u, nil
}
  