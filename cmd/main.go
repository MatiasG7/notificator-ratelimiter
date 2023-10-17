package main

import "github.com/notificator-ratelimiter/cmd/notificator"

func main() {
	n := notificator.NewNotificator()
	n.Send("news", "user1", "news user1")
	n.Send("status", "user1", "status user1")
	n.Send("news", "user1", "news user1")
	n.Send("status", "user1", "status user1")

	n.Send("news", "user2", "news user2")
	n.Send("status", "user2", "status user2")
	n.Send("news", "user2", "news user2")
	n.Send("status", "user2", "status user2")

	n.Send("news", "user3", "news user3")
	n.Send("status", "user3", "status user3")
	n.Send("news", "user3", "news user3")
	n.Send("status", "user3", "status user3")

	n.Send("news", "user1", "news user1")
	n.Send("status", "user1", "status user1")
	n.Send("news", "user1", "news user1")
	n.Send("status", "user1", "status user1")

	n.Send("news", "user2", "news user2")
	n.Send("status", "user2", "status user2")
	n.Send("news", "user2", "news user2")
	n.Send("status", "user2", "status user2")

	n.Send("news", "user3", "news user3")
	n.Send("status", "user3", "status user3")
	n.Send("news", "user3", "news user3")
	n.Send("status", "user3", "status user3")
}
