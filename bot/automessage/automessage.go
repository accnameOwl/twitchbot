package automessage

import (
	"time"
)

type Message struct {
	Msg         string
	Timestamp   time.Time
	AutoDelay   time.Duration //seconds
	ToggleDelay bool
}

var defaultMessages = []Message{
	Message{
		Msg:         "Hit the follow button!",
		Timestamp:   time.Now(),
		AutoDelay:   360,
		ToggleDelay: true,
	},
	Message{
		Msg:         "Use your twitch prime to your favourite streamer!",
		Timestamp:   time.Now(),
		AutoDelay:   360,
		ToggleDelay: true,
	},
}

func (msg *Message) AutoMessage(c chan string) {
	currTime := time.Now()

	if currTime.After(msg.Timestamp) {
		msg.Timestamp = currTime.Add(time.Second * msg.AutoDelay)
		c <- msg.Msg
	}
}
