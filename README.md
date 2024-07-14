# todo-app-1
https://zenn.dev/uta8a/scraps/f94e8c53ae8d6e

setup

```bash
export PROJECT="Google Cloud Project ID"
export FUNCTION_TARGET="todoapp"
export GCLOUD_TOKEN=$(gcloud auth print-identity-token)
export URL="Cloud Functions URL"
```

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

frontend

```bash
$ deno run --allow-all main.ts -- post "宿題やる"

$ deno run --allow-all main.ts -- put 2900821b-ac99-457f-8afc-a7491dd34356 true "ご飯食べる"

$ deno run --allow-all main.ts -- done be8430a4-f628-4113-a047-fc3cb2e36db1

$ deno run --allow-all main.ts -- del c83a3ea4-030b-4a7f-9a49-334b2afbc033

$ deno run --allow-all main.ts -- show
⬜ コード書く (id: 112a16a4-afe1-4b21-8cf7-6bac54187a91)
✅ ご飯食べる (id: 2900821b-ac99-457f-8afc-a7491dd34356)
✅ 宿題やる (id: be8430a4-f628-4113-a047-fc3cb2e36db1)
```
