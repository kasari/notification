package main

import "github.com/kasari/notification"

func main() {
	notification.Message("hello").WithTitle("I'm Title").Push()
}
