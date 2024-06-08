package main

import (
	"bufio"
	"danhenderson95/websockets/openai"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func getUserOpenAIKey() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your OpenAI API key: ")
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func setOpenAIKeyEnvVar(key string) {
	os.Setenv("OPENAI_API_KEY", key)
}

func chatHandler(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)

	for {
		msgType, msg, err := conn.ReadMessage()

		if err != nil {
			return
		}

		openAIResponse := openai.MakeOpenAICall(string(msg))

		if err = conn.WriteMessage(msgType, []byte(openAIResponse)); err != nil {
			return
		}
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "websockets.html")
}

func main() {
	key := getUserOpenAIKey()
	setOpenAIKeyEnvVar(key)

	http.HandleFunc("/chat", chatHandler)
	http.HandleFunc("/", homeHandler)

	fmt.Printf("Starting frontend at: http://localhost:%d\n", 8080)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
