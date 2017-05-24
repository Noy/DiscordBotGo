package main

import (
	"strings"
	"log"
)

func addPermissionFor(msg BotMessage) {
	if strings.HasPrefix(msg.Message, "addperm") {
		if HasPermissionUser(msg.Author.Name) {
			name := strings.TrimLeft(msg.Message, "addperm ")
			AddPerm(name)
			log.Println(Perm.Users)
			msg.SendMessage("Added " + name + " to the permissions list.")
		} else {
			msg.SendMessage("You don't have permission")
		}
	}
}

func removePermissionFor(msg BotMessage) {
	if strings.HasPrefix(msg.Message, "rmperm") {
		if HasPermissionUser(msg.Author.Name) {
			name := strings.TrimLeft(msg.Message, "rmperm ")
			RmPerm(name)
			log.Println(Perm.Users)
			msg.SendMessage("Removed " + name + " to the permissions list.")
		} else {
			msg.SendMessage("You don't have permission")
		}
	}
}
