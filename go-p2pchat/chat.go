package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"[[.ModName]]/control"
	"[[.ModName]]/peer"
	"[[.ModName]]/util"
	log "github.com/sirupsen/logrus"
)

var (
	port   = flag.String("p", "1111", "port")
	user   = flag.String("user", "name@192.168.10.1:1111", "connect user")
	myName = flag.String("name", "ryomak", "my name ")
	debug  = flag.Bool("debug", false, "debug mode")
	public = flag.Bool("public", false, "network")
)
var ctrl = control.Control{
	UpdatedText:         make(chan string, 10),
	UpdateUserList:      make(chan []peer.User, 10),
	UpdatedTextFromUser: make(chan string, 10),
}

func init() {
	flag.Parse()
	if *debug {
		log.Info("debug mode")
	} else {
		log.SetLevel(log.WarnLevel)
	}
	if *public {
		log.Warn("public network")
	} else {
		log.Info("private network")
	}
	peer.SetMyName(*myName)
}

func main() {
	if peer.GetUserFromStr(*user).IP != util.GetMyIP() || !*public {
		go peer.IntroduceMyself(peer.GetUserFromStr(*user))
	}
	go ctrl.StartControlLoop()
	go peer.RunServer(*port, ctrl.UpdatedText, ctrl.UpdateUserList)
	UserInput()
}

func UserInput() {
	fmt.Printf("\x1b[36m%s\x1b[0m", "Hello "+peer.GetMyName()+" from "+util.GetMyIP()+".\nPrivate messages ...'(message)*(user) .\n To leave...'close'")
	fmt.Println()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		ctrl.UpdatedTextFromUser <- text
	}
}
