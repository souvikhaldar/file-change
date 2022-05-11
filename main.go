package main

import (
	"flag"
	"log"

	"github.com/fsnotify/fsnotify"
)

func main() {
	filePath := flag.String("file", "", "path to the file to watch")
	flag.Parse()
	log.Println("Watching file: ", *filePath)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Println("Error creating watcher: ", err)
		return
	}
	defer watcher.Close()
	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("Event: ", event)
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("Error in watching: ", err)
			}
		}
	}()

	err = watcher.Add(*filePath)
	if err != nil {
		log.Fatal("Error in adding file to watch: ", err)
	}
	<-done
}
