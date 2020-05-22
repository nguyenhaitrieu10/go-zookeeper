package main

import (
	"database/sql"
	"errors"
	"strings"
)

var (
	INVALID_DOMAIN  = errors.New("Invalid domain format")
	INVALID_IP_LIST = errors.New("Invalid ip list format")
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

func createDomain(database *sql.DB, domain string, ip []string) error {
	check := validateDomainFormat(domain)
	if !check {
		return INVALID_DOMAIN
	}

	statement, err := database.Prepare("INSERT INTO dns (domain, ip) VALUES (?, ?)")
	if err != nil {
		return err
	}
	_, err = statement.Exec(domain, concatString(ip))
	return err
}

func deleteDomain(database *sql.DB, domain string) error {
	check := validateDomainFormat(domain)
	if !check {
		return INVALID_DOMAIN
	}

	statement, err := database.Prepare("DELETE FROM dns WHERE domain=?")
	if err != nil {
		return err
	}
	_, err = statement.Exec(domain)
	return err
}

func updateDomainIP(database *sql.DB, domain string, ipList string) error {
	check := validateDomainFormat(domain)
	if !check {
		return INVALID_DOMAIN
	}

	check = validateIPListFormat(ipList)
	if !check {
		return INVALID_IP_LIST
	}

	statement, err := database.Prepare("UPDATE dns SET ip=? WHERE domain=?")
	if err != nil {
		return err
	}
	_, err = statement.Exec(ipList, domain)
	return err
}

func getDomainIP(database *sql.DB, domain string) (string, error) {
	check := validateDomainFormat(domain)
	if !check {
		return "", INVALID_DOMAIN
	}

	rows, err := database.Query("SELECT ip FROM dns WHERE domain=" + domain)
	if err != nil {
		return "", err
	}

	var ipList string
	if rows.Next() {
		err = rows.Scan(&ipList)
	}
	if err != nil {
		return "", err
	}

	return ipList, nil
}

//func getCurrentDNS(database *sql.DB) error {
//
//}
