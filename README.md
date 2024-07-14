# todo-app-1
https://zenn.dev/uta8a/scraps/f94e8c53ae8d6e

local test

```bash
# prepare
gcloud emulators firestore start
## set environment variables
export FIRESTORE_EMULATOR_HOST=[::1]:****

# exec local server
cd cloud-functions
go run cmd/main.go
```
