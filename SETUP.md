## Setup

To run the poc

Start the `Neo4J` DB
```shell
docker run --publish=7474:7474 --publish=7687:7687 -e NEO4J_AUTH=none neo4j
```

If data need to be persistent the issue:
```shell
docker run --publish=7474:7474 --publish=7687:7687 --volume=$PWD/neo4j_data:/data -e NEO4J_AUTH=none neo4j
```


And issue:
```shell
 go run cmd/main.go
```

View the data at [http://localhost:7474/browser/preview/](http://localhost:7474/browser/preview/)