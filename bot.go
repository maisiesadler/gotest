package main

import (
	"fmt"
	"os"

	"github.com/nlopes/slack"
)

func main() {

	token := os.Getenv("SLACK_TOKEN")
	fmt.Printf("ST: %v\n", token)
	api := slack.New(token)
	rtm := api.NewRTM()
	go rtm.ManageConnection()

Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			fmt.Print("Event Received: ")
			switch ev := msg.Data.(type) {
			case *slack.ConnectedEvent:
				fmt.Println("Connection counter:", ev.ConnectionCount)

			case *slack.MessageEvent:
				fmt.Printf("Message: %v\n", ev)
				//	info := rtm.GetInfo()
				//prefix := fmt.Sprintf("<@%s> ", info.User.ID)

				// user, _ := rtm.GetUserInfo(ev.User)

				// user := info.GetUserByID(ev.User)
				output := get(ev.Text)
				fmt.Printf("%v ", output)

				//if ev.User != info.User.ID && strings.HasPrefix(ev.Text, prefix) {
				rtm.SendMessage(rtm.NewOutgoingMessage(output, ev.Channel))
				//}

			case *slack.RTMError:
				fmt.Printf("Error: %s\n", ev.Error())

			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid credentials")
				break Loop

			default:
				//Take no action
			}
		}
	}
}
