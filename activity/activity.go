package activity

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type UserActResponse struct {
	Event GitHubEvent
}

func (u *UserActResponse) GetUserActivity(user string) ([]GitHubEvent, error) {
	log.Println("Getting user activity")
	resp, err := http.Get("https://api.github.com/users/" + user + "/events")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// unmarshal actor info
	var events []GitHubEvent
	err = json.Unmarshal([]byte(body), &events)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User is: ", events[0].Actor.Login)
	return events, nil
}

func (u *UserActResponse) PrintUserActivity(event []GitHubEvent) {
	// print the user activity for each event
	for _, event := range event {
		switch event.Type {
		case "PushEvent":
			fmt.Printf("PushEvent: %s Pushed to %s", event.Actor.Login, event.Repo.Name)
			// case "CreateEvent":
			// 	fmt.Println("CreateEvent: ", event.Payload.Ref)
			// case "PullRequestEvent":
			// 	fmt.Println("PullRequestEvent: ", event.Payload.Ref)
		}
	}
}
