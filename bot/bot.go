package bot

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/accnameowl/twitchbot/bot/automessage"
)

type Bot struct {
	server         string
	port           string
	nick           string
	channel        string
	conn           net.Conn
	mods           map[string]bool
	userMaxLastMsg int
	lastfm         string
	quoteBuffer    []automessage.Message
}

func New() *Bot {
	bot := &Bot{
		server:         os.Getenv("BOT_SERVER"),
		port:           os.Getenv("BOT_PORT"),
		nick:           os.Getenv("BOT_NICK"),
		channel:        os.Getenv("BOT_CHANNEL"),
		conn:           nil,
		mods:           make(map[string]bool),
		userMaxLastMsg: 2,
		lastfm:         "NexTliFE_",
	}
	bot.quoteBuffer = defaultMessages
	return bot
}

func (bot *Bot) Connect() error {
	var err error
	fmt.Printf("Attempting to connect to server...\n")

	//attempt connection to twitch IRC server
	bot.conn, err = net.Dial("tcp", bot.server+":"+bot.port)
	if err != nil {
		fmt.Printf("Unable to connect to Twitch IRC server! Reconnecting in 10 seconds...\n")
		//return error then reiterate Bot.Connect() after 10 seconds
		return err
	}

	fmt.Printf("Connected to IRC server %s\n", bot.server)
	return err
}

func (bot *Bot) Message(message string) {
	if message == "" {
		return
	}
	//send message to chat
	fmt.Fprintf(bot.conn, bot.channel+" :"+message+"\r\n")
}

//add a quote to the bot, then append it to the slice of quoteBuffer
func (bot *Bot) NewQuote(_quote string, _autodelay time.Duration, _toggleDelay bool) {
	newQuote := automessage.Message{
		Msg:         _quote,
		AutoDelay:   _autodelay,
		ToggleDelay: _toggleDelay,
	}
	bot.quoteBuffer = append(bot.quoteBuffer, newQuote)
}

func (bot *Bot) RuntimeQuotes() {
	quoteCh := make(chan string)
	for {
		for _, messages := range bot.quoteBuffer {
			go messages.AutoMessage(quoteCh)
		}
		msg := <-quoteCh
		bot.Message(msg)
		time.Sleep(1)
	}
}

var defaultMessages = []automessage.Message{
	automessage.Message{
		Msg:         "Hit the follow button!",
		Timestamp:   time.Now(),
		AutoDelay:   360,
		ToggleDelay: true,
	},
	automessage.Message{
		Msg:         "Use your twitch prime to your favourite streamer!",
		Timestamp:   time.Now(),
		AutoDelay:   360,
		ToggleDelay: true,
	},
}
