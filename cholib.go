package main

import (
	"fmt"
	"time"
	"golang.org/x/net/websocket"
)
// IMPORTS, BITCH!

func onInit() string {
	return "[INIT] Golang 10 min lib for Shitty mongo/Chatango (@Coil) initiated. " + time.Now().String()
}

func main() {
	init := 1
	fmt.Println(onInit()) // Friendly little init message.
	origin := "http://st.chatango.com" // Origin required, thank Chatango for that.
	url := "ws://s23.chatango.com:1800" // Websocket server for the chat.
	ws, err := websocket.Dial(url, "", origin) // Create a websocket connection
	if err != nil { // IS THIS THING ON?
		fmt.Println("[FATAL] ERR NIGNOG")
	}
	ws.Write([]byte("bauth:hazerdklan::Coil:password\x00")) // tested at the klan
	time.Sleep(10000000000) // Needs a delay. HAVE YOUR FUCKING DELAY, CHATANGO
	ws.Write([]byte("bm:t12j:256:<nED1C24/><f x11555555=\"0\">get rekt mofucka\r\n\x00")) // Send a message
	for init == 1 { // Golang's while loop equivalent
		time.Sleep(100) // MORE SLEEPING FUUUUUUUUUUUUUUUUUUUUU
		buf := make([]byte,1024) // Receive from the websocket
		var n int // cursor
		if n, err = ws.Read(buf); err != nil { // Make sure this shit isn't null, nignog
			fmt.Println("[FATAL] ERR NIGNOG") // but in the event that it is null...
		}
		fmt.Printf("%s\n", buf[:n]) // Print the data!
	}
}
