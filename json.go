package twitchutil

import (
	"net/http"
	"encoding/json"
)

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

type following struct {
	Follows []follow
	Links   links  `json:"_Links"`
	Total   int `json:"_Total"`
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