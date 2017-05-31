# natses
Event Sourcing with NATS Streaming

### Installation

```
go get github.com/plimble/natses
```

### Useage

```go
// Pub
sc, _ := stan.Connect(clusterID, clientID)
sc.Publish("foo", natses.NewEvent("user", data))
```

```go
// Sub

sc.Subscribe("foo", natses.SubEvent(UserEvent))

func UserEvent(msg natses.EventMsg) {
  switch msg.EventType {
  case "Registered":
    ...
  case "EmailUpdated":
    ...
  }
}

```

