package twitchutil

import (
	"reflect"
	"strings"
)

func IsInChannel(user string, channel string) bool {
	response := new(chattersResponse)
	getJson("https://tmi.twitch.tv/group/user/"+channel+"/chatters", &response)
	groups := response.Chatters

	groupsValuePtr := reflect.ValueOf(&groups)
	groupsValue := groupsValuePtr.Elem()

	for i := 0; i < groupsValue.NumField(); i++ {
		group := groupsValue.Field(i)
		groupArray := group.Interface().([]string)
		for u := 0; u < len(groupArray); u++ {
			if strings.EqualFold(groupArray[u], user) {
				return true
			}
		}
	}

	return false
}
