package message_saver

func SaveMessage(msg Message)  {
		if  InsertMessage(msg) == nil{
			return
		}else{
			beckToQueue(msg)
		}
}

func beckToQueue(msg Message){
	return
}