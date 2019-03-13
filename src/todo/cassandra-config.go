package todo

import "github.com/gocql/gocql"

func InitCluster() *gocql.Session {
	cluster := gocql.NewCluster("172.17.0.1")
	cluster.Keyspace = "example"
	cluster.Consistency = gocql.Quorum
	session, _ := cluster.CreateSession()
	return session
}
