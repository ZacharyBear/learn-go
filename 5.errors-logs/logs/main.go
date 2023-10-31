package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/rs/zerolog"
	log1 "github.com/rs/zerolog/log"
)

func main() {
	log.SetPrefix("main(): ")
	log.Print("Hey, Je suis un log!")

	// this will make program exited
	// log.Fatal("Hey, Je suis un log!")

	// log.Panic("Hey, Je suis un log!")
	fmt.Println("Can you see me?")

	// # write to file
	testFileOutput()

	// # Zerolog(Framework)
	testZerolog()
}

func testFileOutput() {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.Print("Hey, Je suis un log!")
}

func testZerolog() {
	// try: go run main.go --debug
	debug := flag.Bool("debug", false, "Wether show debug logs")
	flag.Parse()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log1.Print("Hola, Je suis un log!")

	// Context Data
	log1.Info().
		Int("EmployeeID", 1001).
		Msg("Getting employee infomation")

	log1.Debug().
		Str("Name", "John").
		Send()
}
