package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/golang/glog"
	"github.com/lekkehart/slack-user-manager/lib"
)

func main() {

	// TODO ENABLE - token, active := lib.ParseFlags()
	token, _ := lib.ParseFlags()
	glog.Infoln("---------- main() START ----------")

	file, err := os.Open("list_of_userids.in")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// TODO ENABLE - lib.ActivateUserInSlack(token, active, scanner.Text())
		lib.RemoveTitlesAndPhoneInSlack(token, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	glog.Infoln("---------- main() END ----------")
}
