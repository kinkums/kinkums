package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func sendMsgToMQ() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	channel, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer channel.Close()

	q, err := channel.QueueDeclare(
		"hello", //name
		false,   //durable
		false,   //delete when unused
		false,   //exclusive
		false,   //no wait
		nil,     //arguments
	)

	failOnError(err, "Failed to declare a queue")

	body := "Image stored to RabbitMQ"
	err = channel.Publish(
		"",     //exchange
		q.Name, //routing key
		false,  //mandatory
		false,  //immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	fmt.Println(" Sent ", body)
	//log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")

}

func uploadImage(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, nil)
		fmt.Println(" Method is Get ")
	} else if req.Method == "POST" {
		file, handler, err := req.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Println(" Method is Post ")

		fmt.Fprintf(w, "%v", handler.Header)
		img, err := os.OpenFile("./imageOutput/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		io.Copy(img, file)
		sendMsgToMQ()
	} else {
		fmt.Println("Unknown HTTP " + req.Method + " Method")
	}

}

func main() {
	http.HandleFunc("/upload", uploadImage)
	http.ListenAndServe(":9090", nil)
	//sendMsgToMQ()
}
