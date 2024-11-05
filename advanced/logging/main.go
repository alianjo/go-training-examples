package main

import (
	"log"
	"log/slog"
	"os"
)

func main() {
	log.Println("HIIII")
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	log.Println("with micro")

	myCustomLog := log.New(os.Stderr, "My: ", log.LstdFlags)
	myCustomLog.Println("HIIIIIII")

	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")

	myslog.Info("hello again", "key", "val", "age", 25)
}
