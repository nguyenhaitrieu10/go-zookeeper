package main

import (
	"fmt"
	"go-zookeeper/zk"
	"strings"
	"time"
)

func concatString(strList []string) string {
	return strings.Join(strList, "|")
}

func splitString(str string) []string {
	return strings.Split(str, "|")
}

func diff2Sets(l1 []string, l2 []string) (needDeleteList []string, needCreateList []string) {
	s := make(map[string]bool)

	for _, e := range l1 {
		s[e] = true
	}

	for _, e := range l2 {
		if !s[e] {
			needCreateList = append(needCreateList, e)
		}
		s[e] = false
	}

	for i, v := range s {
		if v {
			needDeleteList = append(needDeleteList, i)
		}
	}
	return needDeleteList, needCreateList
}

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
