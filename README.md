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
sc.Publish("foo", natses.NewEvent("Registered", registeredEvent))
sc.Publish("foo", natses.NewEvent("EmailUpdated", emailUpdatedEvent))
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

