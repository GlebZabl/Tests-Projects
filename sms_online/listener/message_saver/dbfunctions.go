package message_saver

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func InsertMessage(msg Message) error  {
	db, err := sql.Open("mysql", DBConnectionString)
	if err != nil {
		return err
	}

	defer db.Close()

	sql := "INSERT INTO `messages`(`text`,`http_get_date`,`listener_get_date`,`insert_date`) VALUES(?,?,?,?);"

	insertDate := time.Now().Unix()

	result, err := db.Exec(sql, msg.Text,msg.GetByHTTPDate,msg.GetByListenerDate, insertDate)

	if err != nil{
		return err
	}

	affected,err := result.RowsAffected()
	if affected != 1 || err != nil{
		return *new(error)
	}

	return nil
}
