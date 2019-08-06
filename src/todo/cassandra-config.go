package todo

import (
	"log"
	"os"

	"github.com/gocql/gocql"
)

const (
	CREATE_KEYSPACE    = " CREATE KEYSPACE IF NOT EXISTS " + KEYSPACE_NAME + " WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };"
	CREATE_TABLE       = "create table if not exists example.todo (id uuid PRIMARY KEY, text text);"
	KEYSPACE_NAME      = "example"
	CASSANDRA_URL      = "CASSANDRA_URL"
	CASSANDRA_USERNAME = "CASSANDRA_USERNAME"
	CASSANDRA_PASSWORD = "CASSANDRA_PASSWORD"
)

func InitCluster() *gocql.Session {
	cassandra := envVar(CASSANDRA_URL)
	auth := gocql.PasswordAuthenticator{Username: envVar(CASSANDRA_USERNAME), Password: envVar(CASSANDRA_PASSWORD)}
	cluster := createCluster(cassandra, KEYSPACE_NAME, auth)
	session, err := cluster.CreateSession()
	//defer session.Close()
	if err != nil {
		log.Fatal("FATAL", err)
	}
	createTodoTable(session)
	return session
}

func createCluster(host string, keyspace string, authentication gocql.PasswordAuthenticator) *gocql.ClusterConfig {
	cluster := gocql.NewCluster(host)
	cluster.Authenticator = authentication
	createKeyspace(keyspace, cluster)
	cluster.Keyspace = keyspace
	cluster.Consistency = gocql.One
	return cluster
}

func envVar(key string) string {
	cassandra, exists := os.LookupEnv(key)
	if !exists {
		log.Fatal("FATAL", "Valor de "+key+" no encontrado")
	} else {
		log.Println("INFO", "Recuperado valor de "+key)
	}
	return cassandra
}

func createKeyspace(keyspace string, cluster *gocql.ClusterConfig) {
	session, err := cluster.CreateSession()
	defer session.Close()
	if err != nil {
		log.Fatal("FATAL", err)
	}
	if err := session.Query(CREATE_KEYSPACE).Exec(); err != nil {
		log.Fatal("FATAL", err)
	}
	log.Println("INFO", "Configurado keyspace: "+keyspace)
}

func createTodoTable(session *gocql.Session) {
	log.Println("INFO", "Creando tabla si no existe...")
	session.Query(CREATE_TABLE).Exec()
}
