package main

import (
	"flag"
	"fmt"
	"github.com/nhomble/groupme.go/groupme"
	"math/rand"
	"time"
)

var (
	groupID = flag.String("groupId", "", "id of the group to post the message into")
	botID   = flag.String("botId", "", "id of the bot post the message")
	message = flag.String("m", "", "the message")
)

func main() {
	flag.Parse()
	if len(*groupID) <= 0 && len(*botID) <= 0 {
		flag.Usage()
		panic("groupId and botId cannot both be empty")
	}
	rand.Seed(time.Now().UnixNano())
	id := fmt.Sprintf("%d%d", rand.Int63(), rand.Int63())
	fmt.Printf("messageId=%s\n", id)
	provider := groupme.EnvironmentTokenProvider{}
	var err error

	client, err := groupme.NewClient(provider)
	must(err)

	if len(*botID) > 0 {
		err = client.Bots.Send(groupme.BotMessageCommand{
			BotID:   *botID,
			Message: *message,
		})
	} else {
		_, err = client.Messages.Send(*groupID, &groupme.SendMessageCommand{
			SourceGuid: id,
			Text:       *message,
		})
	}
	must(err)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
