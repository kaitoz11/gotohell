package utils

import (
	"strconv"
	"strings"
)

type Notification struct {
	targetURL string
	topic     string
	title     string
	message   string
	tags      []string
	priority  int
	icon      string
}

func MakeNotification(target string) *Notification {
	ntfy := &Notification{
		targetURL: target,
	}
	return ntfy
}

func (ntfy *Notification) SetOptions(topic string, icon string, title string, message string, tags []string, priority int) {
	ntfy.topic = topic
	ntfy.icon = icon
	if ntfy.icon == "" {
		ntfy.icon = "https://avatars.githubusercontent.com/u/43519768?v=4"
	}
	ntfy.title = title
	ntfy.message = message
	ntfy.tags = tags
	if ntfy.tags == nil {
		ntfy.tags = []string{"skull_and_crossbones"}
	}
	ntfy.priority = priority
	if ntfy.priority < 1 || ntfy.priority > 5 {
		ntfy.priority = 3
	}
}

func (ntfy *Notification) SendNotification() (bool, error) {
	if ntfy.targetURL == "" || ntfy.topic == "" || ntfy.message == "" {
		return false, nil
	}

	headers := map[string][]string{
		"X-Icon":     {ntfy.icon},
		"X-Title":    {ntfy.title},
		"X-Tags":     {strings.Join(ntfy.tags, ",")},
		"X-Priority": {strconv.Itoa(ntfy.priority)},
	}

	url := ntfy.targetURL
	if ntfy.targetURL[len(ntfy.targetURL)-1:] != "/" {
		url += "/"
	}
	url += ntfy.topic

	res, err := SendPOST(url, ntfy.message, nil, headers)
	if err != nil {
		return false, err
	}

	defer res.Body.Close()

	return true, nil
}
