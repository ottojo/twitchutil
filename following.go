package twitchutil

func FollowingChannels(user string) []string {

	response := new(following)
	getJson("https://api.twitch.tv/kraken/users/" + user + "/follows/channels", &response)
	follows := response.Follows
	var result []string
	result = append(result, user)
	for i := 0; i < len(follows); i++ {
		result = append(result, follows[i].Channel.Name)
	}
	for (len(result) < response.Total) {
		url := response.Links.Next
		response = new(following)
		getJson(url, &response)
		follows := response.Follows
		for i := 0; i < len(follows); i++ {
			result = append(result, follows[i].Channel.Name)
		}
	}

	return result
}

