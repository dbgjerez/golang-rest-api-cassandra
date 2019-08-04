# go-todo-rest-api-cassandra
Example using CQL and golang REST API. 

# Cassandra
Cassandra can be deployed using Docker. Cassandra is a powerfull NoSQL database used to store big amount of data. 

# Deploy
## Docker
To run Cassandra in Docker container: 
```bash
docker run -p 9042:9042 -d --name cassandra cassandra
```

When Cassandra is running, the following step is run the application, linking it with the database:

To build the container: 
```bash
docker build -t todo-api .
```

When the application has been built as Docker image, to run it:
```bash
docker run -p 8000:8000 -e CASSANDRA_URL=cassandra:9042 --link=cassandra cassandra-test
```

## Kubernetes


### Docker

### Config keyspace and table

```bash
docker run -it --link cassandra --rm cassandra sh -c 'exec cqlsh "$CASSANDRA_PORT_9042_TCP_ADDR"'

CREATE KEYSPACE IF NOT EXISTS "example" WITH REPLICATION = {     'class' : 'SimpleStrategy',     'replication_factor' : 1    };

create table if not exists example.todo (id uuid PRIMARY KEY, text text); 
``` 
