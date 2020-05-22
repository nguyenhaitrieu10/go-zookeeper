package main

import (
	"database/sql"
	"errors"
)

var (
	INVALID_DOMAIN  = errors.New("Invalid domain format")
	INVALID_IP_LIST = errors.New("Invalid ip list format")
)

func createDomain(database *sql.DB, domain string, ip string) error {
	check := validateDomainFormat(domain)
	if !check {
		return INVALID_DOMAIN
	}

	statement, err := database.Prepare("INSERT INTO dns (domain, ip) VALUES (?, ?)")
	if err != nil {
		return err
	}
	_, err = statement.Exec(domain, ip)
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

func getCurrentDNS(database *sql.DB) (map[string]string, error) {
	rows, err := database.Query("SELECT domain, ip FROM dns")
	if err != nil {
		return nil, err
	}

	dns := make(map[string]string)
	var domain string
	var ip string
	for rows.Next() {
		err = rows.Scan(&domain, &ip)
		if err != nil {
			return nil, err
		}
		dns[domain] = ip
	}

	return dns, nil
}
