package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const targetDirectory string = "/var/file-spammer"

func init() {
	rand.Seed(time.Now().UnixNano())
	os.MkdirAll(targetDirectory, os.ModePerm)
}

func main() {
	for i := 0; i < 10; i++ {
		go createFiles(targetDirectory)
	}

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-c
}

func createFiles(dir string) {
	for {
		fileName := rand.Uint64()
		filePath := fmt.Sprintf("%s/%d", targetDirectory, fileName)
		file, err := os.Create(filePath)
		if err != nil {
			log.Fatalln("Failed to create file:", err)
		}
		fileContent := make([]byte, 1024)
		rand.Read(fileContent)
		file.Write(fileContent)
		file.Close()
		fmt.Printf("Created file %d\n", fileName)
		time.Sleep(time.Second * time.Duration(rand.Intn(5)))
	}
}
