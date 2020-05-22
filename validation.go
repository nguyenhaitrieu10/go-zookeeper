package main

import "regexp"

func validateDomainFormat(domain string) bool {
	if domain == "" || len(domain) > 255 {
		return false
	}

	matched, err := regexp.MatchString(`^[a-z0-9\-]*\.[a-z0-9\.\-]*$`, domain)
	if err != nil || !matched {
		return false
	}

	return true
}

func validateIPListFormat(ipList string) bool {
	if ipList == "" || len(ipList) > 160 {
		return false
	}

	matched, err := regexp.MatchString(`^([0-9]+\.[0-9]+\.[0-9]+\.[0-9]+\|?)+$`, ipList)
	if err != nil || !matched {
		return false
	}

	return true
}
