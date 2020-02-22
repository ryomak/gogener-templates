package control

import (
	"fmt"
	"os"
	"strings"

	"[[.ModName]]/peer"
	log "github.com/sirupsen/logrus"
)

type Control struct {
	UpdatedText         chan string
	UpdatedTextFromUser chan string
	UpdateUserList      chan []peer.User
}

func (ctrl *Control) StartControlLoop() {
	log.Println("Running Control loop")
	for {
		select {
		case update := <-ctrl.UpdatedText:
			ctrl.updateText(update)
		case userListChanged := <-ctrl.UpdateUserList:
			ctrl.updateList(userListChanged)
		case updateFromUser := <-ctrl.UpdatedTextFromUser:
			ctrl.handleUserInput(updateFromUser)
		default:
		}
	}
}

func (ctrl *Control) handleUserInput(input string) {
	whatever := strings.Split(input, "*")
	//close signal
	if input == "close" {
		msg := peer.Message{"DISCONNECT", peer.User{peer.GetMyName(), "", ""}, "", make([]peer.User, 0)}
		msg.Send()
		fmt.Println("bye")
		os.Exit(1)
		return
	}
	//private (message)*(user)
	if len(whatever) > 1 {
		msg := peer.Message{"PRIVATE", peer.User{peer.GetMyName(), "", ""}, whatever[0], make([]peer.User, 0)}
		msg.SendToUser(whatever[1], ctrl.UpdatedTextFromUser)
		ctrl.UpdatedText <- "(private) from " + peer.GetMyName() + ": " + msg.MSG
		return
	}
	//public
	msg := peer.Message{"PUBLIC", peer.User{peer.GetMyName(), "", ""}, whatever[0], make([]peer.User, 0)}
	msg.Send()
	ctrl.UpdatedText <- peer.GetMyName() + ": " + msg.MSG
}

func (ctrl *Control) updateText(str string) {
	fmt.Printf("\x1b[35m%s\x1b[0m", str)
	fmt.Println()
	fmt.Print("-> ")
}

func (ctrl *Control) updateList(list []peer.User) {
	log.Info("list:", list)
}
