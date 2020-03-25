package main

import (
	"KinoBox/database"
	"log"
	//	"net/http"
	"os"
)

func main() {

	// Инициализируем лог-файл
	logFile, _ := os.OpenFile("./log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	defer logFile.Close()
	logger := log.New(logFile, "", log.LstdFlags|log.Lshortfile)

	// Подключимся к базе данных
	db, err := database.InitDB("KinoBox.db", logger)
	if err != nil {
		logger.Fatalf("Подключиться к базе данных не удалось\r\n(%s)", err)
	}
	defer db.Close()

	database.InitModels(db, logger)

	//	http.HandleFunc("/countries", countries)
	//	http.ListenAndServe(":80", nil)
}

//func countries(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-type", "text/plain")
//	w.Write([]byte("Hello World!!!"))
//}
