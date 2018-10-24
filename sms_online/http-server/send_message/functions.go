package send_message

import (
	"encoding/json"
	"fmt"
	"time"
)

func AddToQueue(msg string) (error,string) {
	QuData := Message{Text:msg,GetByHTTPDate:time.Now().Unix()}

	jsonData, err := json.Marshal(QuData)
	if err != nil{
		return err,err.Error()
	}

	fmt.Println(jsonData)
	return nil,""
}
