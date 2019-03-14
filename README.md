# go-todo-rest-api-cassandra
Example using CQL and Go REST API

CREATE KEYSPACE "example" 
   ... WITH REPLICATION = {     'class' : 'SimpleStrategy',     'replication_factor' : 1    };

create table example.todo (id uuid PRIMARY KEY, text text);
