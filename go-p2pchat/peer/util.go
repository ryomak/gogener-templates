package peer

import (
	"errors"
	"strings"

	log "github.com/sirupsen/logrus"
)

func GetFromUserMap(mapData map[string]User) []User {
	values := make([]User, len(mapData))
	idx := 0
	for _, value := range mapData {
		values[idx] = value
		idx++
	}
	return values
}

func EchoUser(text string, u User, level string) {
	switch level {
	case "Fatal":
		log.Fatal(text + " by " + u.Name + "(" + u.IP + ":" + u.Port + ")")
	case "Info":
		log.Info(text + " by " + u.Name + "(" + u.IP + ":" + u.Port + ")")
	case "Debug":
		log.Debug(text + " by " + u.Name + "(" + u.IP + ":" + u.Port + ")")
	default:
		log.Println(text + " by " + u.Name + "(" + u.IP + ":" + u.Port + ")")
	}
}

func GetUserFromStr(str string) User {
	//ex test@192.168.10.1:3000
	at := strings.Index(str, "@")
	col := strings.Index(str, ":")
	if at == -1 || col == -1 {
		panic(errors.New("-user is invalid"))
	}
	return User{
		Name: str[:at],
		IP:   str[at+1 : col],
		Port: str[col+1:],
	}
}
