This is a test repository that shows how I'm connecting to a mongo database.
I'm seeing a performance problem when connecting to a database.
The very first call made after Connect is slow.
The second call is always MUCH faster.

Is there a way to speed up the first call?
I'm seeing this both connecting to a local mongo instance running in Docker or connecting to CosmosDB.

**Connected to Cosmos DB**
```
$ go run main.go
Ping 1.549611747s
Ping 65.599248ms
Connected!
```

**Connected to Atlas**
```
$ go run main.go
Ping 624.839234ms
Ping 33.384318ms
Connected!
```

**Connected to local Mongo**
```
go run main.go
Ping 8.688944ms
Ping 820.645µs
Connected!
```
