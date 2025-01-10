package main

func main() {

	json.Marshal()
}

type 

type Event struct {
	MessageType string    `json:"message_type"`
	Timestamp   time.Time `json:"timestamp"`
	Data        any       `json:"data"`
}