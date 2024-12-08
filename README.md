# toybox-server-go

## フォルダ構成
### main.go
最初に呼び出されるentrypoint

### router.go
ginの各エンドポイントの定義

### controllers/
各エンドポイントでどういった処理をするかの定義

### cruds/
controllersで使用するDBの操作処理の定義

### db/
DBの接続情報などの定義
