package main

import (
	"flag"
	"log"
	"multiplicator-app/multiplier"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {

	// Работа с флагом
	rtpStr := flag.String("rtp", "0,9", "RTP in (0, 1.0]")
	flag.Parse()
	rtpFloat, err := strconv.ParseFloat(*rtpStr, 64)
	if err != nil {
		log.Fatal("err")
	}
	RTP := rtpFloat
	if RTP <= 0 || RTP > 1.0 {
		log.Fatal("rtp must be between (0, 1.0]")
	}
	log.Println(RTP)
	// Нахождение параметра alpha для формулы
	alpha := multiplier.GetAlphaRTP(RTP)

	//Определяем service
	service := multiplier.NewMultiplierService(alpha)

	//Определяем роутер (я использовал gorrila mux)
	router := mux.NewRouter()

	//Определяем handler
	multiplier.NewMultiplierHandler(router, &multiplier.MultiplierHandler{
		MultiplierService: service,
	})

	//Настраиваем сервер
	server := http.Server{
		Handler: router,
		Addr:    ":64333",
	}

	//Включение сервера
	log.Println("Server started in port :64333")
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("Server failed to start")
	}
}
