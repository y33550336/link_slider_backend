# Unity gRPC Client

このフォルダには、Unity からこのバックエンドへ接続するための最小サンプルと、C# クライアント生成手順をまとめています。

## 前提

- このバックエンドは `Ping` と `MovePlayer` を持つ `GameService` を公開します。
- Unity 側は gRPC-Web クライアントを使う想定です。
- 生成された C# コードは `LinkSlider.Grpc` 名前空間になります。

## C# クライアントの生成

`protoc` と C# 用の gRPC プラグインがある環境で、次のように生成します。

```bash
mkdir -p unity/Generated
protoc -I proto \
  --csharp_out=unity/Generated \
  --grpc_out=unity/Generated \
  --plugin=protoc-gen-grpc=/path/to/grpc_csharp_plugin \
  proto/test.proto
```

### Windows 上の Unity プロジェクトへ直接出力する

WSL から Windows の Unity プロジェクトへ直接出力する場合は、出力先を `/mnt/c/.../Assets/Generated` にします。

PLUGIN=/tmp/grpc_tools/tools/linux_x64/grpc_csharp_plugin
OUT_DIR="/mnt/c/Users/u5yuu/UnityGames/link_slider/Assets/Scripts/Generated"

mkdir -p "$OUT_DIR"
chmod +x "$PLUGIN"

protoc -I proto \
  --csharp_out="$OUT_DIR" \
  --grpc_out="$OUT_DIR" \
  --plugin=protoc-gen-grpc="$PLUGIN" \
  proto/test.proto


```bash
OUT_DIR="/mnt/c/Users/u5yuu/UnityGames/link_slider/Assets/Scripts/Generated"
PLUGIN_EXE="/mnt/c/Users/<YourUser>/.nuget/packages/grpc.tools/<version>/tools/windows_x64/grpc_csharp_plugin.exe"

mkdir -p "$OUT_DIR"
protoc -I proto \
  --csharp_out="$OUT_DIR" \
  --grpc_out="$OUT_DIR" \
  --plugin=protoc-gen-grpc="$PLUGIN_EXE" \
  proto/test.proto
```

PowerShell で実行する場合は次の例を使ってください。

```powershell
$outDir = "C:\Users\u5yuu\UnityGames\link_slider\Assets\Scripts\Generated"
$plugin = "C:\Users\<YourUser>\.nuget\packages\grpc.tools\<version>\tools\windows_x64\grpc_csharp_plugin.exe"

New-Item -ItemType Directory -Force -Path $outDir | Out-Null
protoc -I proto `
  --csharp_out="$outDir" `
  --grpc_out="$outDir" `
  --plugin=protoc-gen-grpc="$plugin" `
  proto/test.proto
```

`<version>` は実際にインストールされている Grpc.Tools のバージョンに置き換えてください。

`grpc_csharp_plugin` が無い場合は、`Grpc.Tools` を使う .NET の補助プロジェクトで生成する方法に切り替えてください。

## Unity への取り込み

1. `unity/Generated` に出力した `.cs` ファイルを Unity プロジェクトの `Assets/Generated` にコピーします。
2. `Grpc.Net.Client`、`Grpc.Net.Client.Web`、`Google.Protobuf`、`Grpc.Core.Api` を Unity に入れます。
3. `GameGrpcClientExample.cs` を `Assets/Scripts` に置き、任意の GameObject にアタッチします。

## 接続確認

このサンプルは既定で `http://127.0.0.1:8080` に接続し、失敗した場合は `http://localhost:8080` も自動で試します。
バックエンドを起動したあと、Unity の Console に `pong: ...` が出れば接続成功です。

## WebGL の注意

WebGL でも同じく gRPC-Web を使ってください。
このバックエンドは gRPC-Web をラップしているので、ブラウザ経由の実行でも同じ接続方式で動作します。
