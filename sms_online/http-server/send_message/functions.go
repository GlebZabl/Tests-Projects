package send_message

import (
	"encoding/json"
	"fmt"
)

func AddToQueue(msg string) (error,string) {
	QuData, err := json.Marshal(msg)
	if err != nil{
		return err,err.Error()
	}
	fmt.Println(QuData)
	return nil,""
}
