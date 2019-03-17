# go-todo-rest-api-cassandra
Example using CQL and golang REST API. 

# Cassandra
Cassandra can be deployed using Docker. Cassandra is a powerfull NoSQL database used to store big amount of data. 

## Configuration
To use Cassandra is necessary config it. In this example we can see how to config it using Docker.

### Docker
To run Cassandra in Docker container: 
```bash
docker run -p 9042:9042 -d --name cassandra cassandra
```
### Config keyspace and table

```bash
docker run -it --link cassandra --rm cassandra sh -c 'exec cqlsh "$CASSANDRA_PORT_9042_TCP_ADDR"'

CREATE KEYSPACE "example" WITH REPLICATION = {     'class' : 'SimpleStrategy',     'replication_factor' : 1    };

create table example.todo (id uuid PRIMARY KEY, text text); 
``` 
