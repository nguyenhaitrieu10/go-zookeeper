package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"go-zookeeper/zk"
	"time"
)

func sample() {
	c, _, err := zk.Connect([]string{"127.0.0.1:2181"}, time.Second) //*10)
	//c, _, err := zk.Connect([]string{"103.60.17.237:62181"}, time.Second) //*10)
	if err != nil {
		panic(err)
	}

	domain := "test-for-dba.coredns"
	ip := "10.84.6.112"
	err = DeleteIP(c, domain, ip)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Deleted IP: %v\n", domain+"/"+ip)
	}

	ip = "10.84.6.111"
	createdPath, err := CreateIP(c, domain, ip)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Created Path: %v\n", createdPath[(len(PATH_ROOT)+1):])
	}
}

func syncDns(database *sql.DB, c *zk.Conn) ([]string, error) {

	domainList, err := GetDomainList(c)
	if err != nil {
		return nil, err
	}

	syncedDomains := []string{}
	for _, domain := range domainList {
		fmt.Println(domain)
		ips, err := GetIPList(c, domain)
		if err != nil {
			return syncedDomains, err
		}
		fmt.Println(ips)

		syncedDomains = append(syncedDomains, domain)

	}

	return []string{}, nil
}

func main() {
	//sample()

	//database, err := sql.Open("sqlite3", "/home/tony/working/devops/admin-cloud-go/admin.db")
	//statement, _ := database.Prepare("INSERT INTO dns (domain, ip) VALUES (?, ?)")
	//res, err := statement.Exec("local", "127.0.0.1")
	//fmt.Println("--------------")
	//fmt.Println(err)
	//fmt.Println(res)
	//fmt.Println("--------------")

	//l1 := []string{"127.0.0.1", "0.0.0.0", "10.84.54.112"}
	//l2 := []string{"103.60.17.231", "127.0.0.1"}

}
