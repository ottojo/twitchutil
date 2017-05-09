package twitchutil

import (
	"encoding/json"
	"net/http"
)

func getJson(url string, target interface{}) error {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("client-ID", clientId)
	r, err := client.Do(req)

	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

type following struct {
	Follows []follow
	Links   links `json:"_Links"`
	Total   int   `json:"_Total"`
}

type links struct {
	Next string
}

type follow struct {
	Created_at string
	Channel    channel
}

type channel struct {
	Name, Game, Status string
}

type chattersResponse struct {
	Chatters chatterGroups
}

type chatterGroups struct {
	Moderators  []string
	Staff       []string
	Admins      []string
	Global_mods []string
	Viewers     []string
}
