package todo

import (
	"log"
	"os"

	"github.com/gocql/gocql"
)

const (
	CREATE_KEYSPACE = " CREATE KEYSPACE IF NOT EXISTS " + KEYSPACE_NAME + " WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };"
	CREATE_TABLE    = "create table if not exists example.todo (id uuid PRIMARY KEY, text text);"
	KEYSPACE_NAME   = "example"
)

func InitCluster() *gocql.Session {
	cassandra, exists := os.LookupEnv("CASSANDRA_URL")
	if !exists {
		log.Fatal("URL de cassandra no encontrada")
	} else {
		log.Println("INFO", "Conectando a Cassandra: "+cassandra)
	}
	cluster := gocql.NewCluster(cassandra)
	createKeyspace(cluster)
	cluster.Keyspace = KEYSPACE_NAME
	cluster.Consistency = gocql.Quorum
	session, _ := cluster.CreateSession()
	createTodoTable(session)
	return session
}

func createKeyspace(cluster *gocql.ClusterConfig) {
	session, _ := cluster.CreateSession()
	log.Println("INFO", "Creando keyspace "+KEYSPACE_NAME+" si no existe...")
	session.Query(CREATE_KEYSPACE).Exec()
	session.Close()
}

func createTodoTable(session *gocql.Session) {
	log.Println("INFO", "Creando tabla si no existe...")
	session.Query(CREATE_TABLE).Exec()
}
