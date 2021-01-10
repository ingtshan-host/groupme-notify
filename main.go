package main

import (
	"flag"
	"fmt"
	"github.com/nhomble/groupme.go/groupme"
	"math/rand"
	"time"
)

var (
	groupId = flag.String("groupId", "", "id of the group to post the message into")
	message = flag.String("m", "", "the message")
)

func main() {
	flag.Parse()
	if len(*groupId) <= 0 {
		flag.Usage()
		panic("groupId is empty")
	}
	rand.Seed(time.Now().UnixNano())
	id := fmt.Sprintf("%d%d", rand.Int63(), rand.Int63())
	fmt.Printf("messageId=%s\n", id)
	provider := groupme.EnvironmentTokenProvider{}
	client, _ := groupme.NewClient(provider)
	_, err := client.Messages.Send(*groupId, &groupme.SendMessageCommand{
		SourceGuid: id,
		Text:       *message,
	})
	must(err)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
