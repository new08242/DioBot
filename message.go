package main

const (
	MessageTypeText = "text"
)

type RawMessage struct {
	Events []Event `json:"events"`
}

type Event struct {
	Type       string  `json:"type"`
	ReplyToken string  `json:"replyToken"`
	Timestamp  int     `json:"timestamp"`
	Source     Source  `json:"source"`
	Message    Message `json:"message"`
}

type Source struct {
	Type   string `json:"type"`
	UserID string `json:"userId"`
	GroupID string `json:"groupId"`
	RoomID string `json:"roomId"`
}

type Message struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Text string `json:"text"`
}
