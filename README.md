# natses
Event Sourcing with NATS Streaming

### Installation

```
go get github.com/plimble/natses
```

### Usage

```go
// Pub
sc, _ := stan.Connect(clusterID, clientID)

data, _ := natses.NewEvent("Registered", registeredEvent)
sc.Publish("foo", data)

data, _ = natses.NewEvent("EmailUpdated", emailUpdatedEvent)
sc.Publish("foo", data)
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

