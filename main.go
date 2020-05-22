package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"go-zookeeper/zk"
	"time"
)

func main() {
	database, err := sql.Open("sqlite3", "/home/tony/working/devops/admin-cloud-go/admin.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	c, _, err := zk.Connect([]string{"127.0.0.1:2181"}, time.Second) //*10)
	//c, _, err := zk.Connect([]string{"103.60.17.237:62181"}, time.Second) //*10)
	if err != nil {
		fmt.Println(err)
		return
	}

	statistic, err := syncZk2Db(database, c)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(statistic)
}
