package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	c := session.DB("local").C("$cmd.sys.inprog")

	fmt.Println("Waiting for an index to be applied...\n ")
	for true {
		result := inProgress{}
		c.Find(bson.M{}).One(&result)
		if len(result.Inprog) > 0 {
			printProgress(extractProgress(result.Inprog[0]))
		} else {
			fmt.Println("\033[1A\r\033[K")
		}
		time.Sleep(time.Duration(200) * time.Millisecond)
	}
}

func extractProgress(operation operation) (collection string, name string, curr int64, total int64, duration string) {
	duration = fmt.Sprint(time.Duration(operation.Duration*1000) * time.Nanosecond)
	return operation.Insert.Collection, operation.Insert.Name, operation.Progress.Done, operation.Progress.Total, duration
}

func printProgress(collection string, name string, current int64, total int64, duration string) {
	if total == 0 || current == 0 {
		return
	}
	percent := int((float64(current) / float64(total)) * 100)
	progress := fmt.Sprintf("%d/%d", current, total)
	collectionName := fmt.Sprintf("%s/%s", collection, name)
	doneParts := int((float64(current) / float64(total)) * 30)
	fmt.Printf("\033[1A\r\033[K%s%s %s %s %d%% (time %s)\n",
		drawBar(color.BgGreen, doneParts),
		drawBar(color.BgRed, int(30-doneParts)),
		collectionName, progress, percent, duration)
}

func drawBar(colour color.Attribute, width int) (result string) {
	colourize := color.New(colour).SprintFunc()
	return colourize(strings.Repeat(" ", width))
}

type operation struct {
	Progress progress
	Insert   insert
	Duration int64 `bson:"microsecs_running"`
}

type inProgress struct {
	Inprog []operation
}

type progress struct {
	Done  int64
	Total int64
}

type insert struct {
	Name       string
	Background bool
	Collection string `bson:"ns"`
}
