# fame-events

To update dependencies:
```
bazelisk run //:gazelle-update-repos && bazelisk run //:gazelle
```

Test all packages:
```
bazelisk run //pkg/event:event_test
```

Dapr deploy:
```
dapr run --app-id goapp --app-port 50001 --app-protocol grpc go run main.go
```

TODOs:
- Add tests for missing structs (such as play off optimizer)
- Implement an event scheduler
- A GRPc service:
    - Endpoint for consuming list of players/teams and seeding a group stage
    - Monitoring endpoint
