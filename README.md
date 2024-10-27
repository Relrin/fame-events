# fame-events

To update dependencies:
```
bazelisk run //:gazelle-update-repos && bazelisk run //:gazelle
```

Test all packages:
```
bazelisk run //pkg/tournament:tournament_test
```

Dapr deploy:
```
dapr run --app-id goapp --app-port 50001 --app-protocol grpc go run main.go
```