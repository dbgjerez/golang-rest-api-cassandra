package todo

import "github.com/gocql/gocql"

func InitCluster() *gocql.Session {
	cluster := gocql.NewCluster("cassandra")
	cluster.Keyspace = "example"
	cluster.Consistency = gocql.Quorum
	session, _ := cluster.CreateSession()
	return session
}
