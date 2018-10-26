package message_saver

type Message struct {
	Text              string `json:"text"`
	GetByHTTPDate     int64  `json:"get_date"`
	GetByListenerDate int64  `json:"get_by_listener_date"`
	InsertDate        int64  `json:"insert_date"`
}

