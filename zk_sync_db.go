package main

import (
	"database/sql"
	"fmt"
	"go-zookeeper/zk"
)

var (
	NULL = "NULL"
)

type SyncStatistic struct {
	createdDomains []string
	updatedDomains []string
	deletedDomains []string
}

func (statistic SyncStatistic) String() string {
	return fmt.Sprintf("createdDomains: %v\nupdatedDomains: %v\ndeletedDomains: %v\n",
		statistic.createdDomains, statistic.updatedDomains, statistic.deletedDomains)
}

func syncZk2Db(database *sql.DB, c *zk.Conn) (statistic *SyncStatistic, err error) {

	domainListFromZk, err := GetDomainList(c) // from zookeeper
	if err != nil {
		return statistic, err
	}

	dnsFromDb, err := getCurrentDNS(database) // from db
	if err != nil {
		return statistic, err
	}

	statistic = &SyncStatistic{}
	for _, domain := range domainListFromZk {
		fmt.Println(domain)
		ipListFromZk, err := GetIPList(c, domain)
		if err != nil {
			return statistic, err
		}
		if len(ipListFromZk) == 0 {

		}
		ipStringFromZk := concatString(ipListFromZk)

		ipStringFromDb, exist := dnsFromDb[domain]
		if !exist {
			err = createDomain(database, domain, concatString(ipListFromZk))
			if err != nil {
				return statistic, err
			}
			statistic.createdDomains = append(statistic.createdDomains, domain)
		} else if ipStringFromZk != ipStringFromDb {
			err = updateDomainIP(database, domain, concatString(ipListFromZk))
			if err != nil {
				return statistic, err
			}
			statistic.updatedDomains = append(statistic.updatedDomains, domain)
		}
		dnsFromDb[domain] = NULL
	}

	for domain, _ := range dnsFromDb {
		if dnsFromDb[domain] != NULL {
			err = deleteDomain(database, domain)
			if err != nil {
				return statistic, err
			}
			statistic.deletedDomains = append(statistic.deletedDomains, domain)
		}
	}

	return statistic, nil
}
