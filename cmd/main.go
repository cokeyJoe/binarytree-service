package main

import (
	"binarytree/pkg/api"
	"binarytree/pkg/logging"
	"binarytree/pkg/logging/bstlogger"
	"binarytree/pkg/logging/httplog"
	"binarytree/pkg/tree/utils"
	"flag"
	"log"
	"os"
)

func main() {

	path := flag.String("path", "ints.json", "path to ints json file")
	flag.Parse()

	f, err := os.Open(*path)
	if err != nil {
		log.Fatal(err)
	}

	bst, err := utils.FromReader(f, utils.WithMinCount(30))

	// NOTE: немного не смотрится здесь, но дефереом закрывать - будет висеть пока не остановится сервер
	// при этом файл прочитан, все остальное - не его забота
	f.Close()

	if err != nil {
		log.Fatal(err)
	}

	jsonLogger := logging.New(os.Stdout)

	loggedBst := bstlogger.NewBSTLogger(bst, jsonLogger)

	httpLogger := httplog.New(jsonLogger)

	httpAPI := api.New(loggedBst, ":8000")
	httpAPI.Use(httpLogger.LogHTTP)

	if err := httpAPI.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
