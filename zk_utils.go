package main

import "go-zookeeper/zk"

const PATH_ROOT = "/dns"

func GetPath(c *zk.Conn, path string) ([]string, error) {
	znode, _, _, err := c.ChildrenW(path)
	return znode, err
}

func CreatePath(c *zk.Conn, path string) (string, error) {
	p, err := c.Create(path, nil, 0, zk.WorldACL(zk.PermAll))
	return p, err
}

func DeletePath(c *zk.Conn, path string) error {
	err := c.Delete(path, -1)
	return err
}

func GetDomainList(c *zk.Conn) ([]string, error) {
	return GetPath(c, PATH_ROOT)
}

func GetIPList(c *zk.Conn, domain string) ([]string, error) {
	path := PATH_ROOT + "/" + domain
	return GetPath(c, path)
}

func CreateDomain(c *zk.Conn, domain string) (string, error) {
	createPath := PATH_ROOT + "/" + domain
	return CreatePath(c, createPath)
}

func CreateIP(c *zk.Conn, domain string, ip string) (string, error) {
	createPath := PATH_ROOT + "/" + domain + "/" + ip
	return CreatePath(c, createPath)
}

func DeleteDomain(c *zk.Conn, domain string, ip string) error {
	deletePath := PATH_ROOT + "/" + domain + "/" + ip
	return DeletePath(c, deletePath)
}

func DeleteIP(c *zk.Conn, domain string, ip string) error {
	deletePath := PATH_ROOT + "/" + domain + "/" + ip
	return DeletePath(c, deletePath)
}
