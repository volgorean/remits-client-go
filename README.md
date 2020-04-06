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
