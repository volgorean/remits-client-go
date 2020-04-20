# remits-client-go
This is a Go client for the [Remits](https://github.com/badtuple/remits) Database... WIP


## Usage

Connecting to Remits
```go
client, err := remits.Connect("0.0.0.0:4242")
if err != nil {
    fmt.Println(err)
    os.Exit(1)
}
defer client.Close()
```

Adding a Log
```go
resp, err := client.AddLog("testlog")
if err != nil {
    panic(err)
}
fmt.Println(resp)
```

Listing Logs
```go
logs, err := client.ListLogs()
if err != nil {
    panic(err)
}
fmt.Println(logs)
```

Getting a Log
```go
log, err := client.ShowLog("testlog")
if err != nil {
    panic(err)
}
fmt.Println(log)
```

Adding a Message to a Log
```go
resp, err = client.AddMessage("testlog", []byte("reeee"))
if err != nil {
    panic(err)
}
fmt.Println(resp)
```

Deleting a Log
```go
resp, err = client.DeleteLog("testlog")
if err != nil {
    panic(err)
}
fmt.Println(resp)
```

Adding a `map` iterator function named `xiter` to a log
```go
resp, err = client.AddIterator("testlog", "xiter", "map", "return msg")
if err != nil {
    panic(err)
}
fmt.Println(resp)
```

Listing all iterators for a log
```go
iterators, err := client.ListIterators("testlog")
if err != nil {
    panic(err)
}
fmt.Println(iterators)
```

Getting a message from an iterator
```go
message, err := client.NextIterator("xiter", 0, 1)
if err != nil {
    panic(err)
}
fmt.Println(string(message))
```

Deleting an iterator
```go
resp, err = client.DeleteIterator("testlog", "xiter")
if err != nil {
    panic(err)
}
fmt.Println(resp)
```