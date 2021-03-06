## App Design

### p2p, accounts, keys and nodes
- Refer to p2p2.md

### App Shell
- We are using https://gopkg.in/urfave/cli.v1 for the CLI app shell

### Runtime Metrics
- We plan using [go-merics](https://github.com/rcrowley/go-metrics) for system metrics

### Local DB
- We are using goleveldb for a key-value db

## Versioning
- This project uses [semantic versioning](http://semver.org/).

## Git workflow
- We follow [git workflow](http://nvie.com/posts/a-successful-git-branching-model/). 
- Feature branches should be created from `develop` and not from `master`.

### Debugging

#### GoLand IDEA interactive debugging
- Should be fully working - create an IDEA config to run main in the proj root source dir

#### Other ways to debug

### Logging
- We are using [go-logging](https://github.com/op/go-logging)
- Use /log/logger wrapper funcs for app-level logging
- You may add custom loggers (per package, to a file, etc...)
- For logging in tests use t.log() and not log.Info()

### Config
- We are using TOML files for config file and cli.v1 flags to modify any configurable param via the CLI.

### RPC
- Use gRPC and grpc rest-json gateway so we'll have both a gRPC and a json-rpc apis.
- Setup app config to start the rpc server for a node
- Make sure no private keys are exposed via rpc
- Setup app config to unlock account(s) via passphrase for session 

## Concurrency 
- We use channels for all concurrent flows.
- We use channels for all callbacks and select based event loops of channel messages processing.
- We do not use any thread-safe data structures beyond channels - all state is updated by internal serialized callbacks from event processing loops.
- [watch me](https://www.youtube.com/watch?v=QDDwwePbDtw)
