package jpmcpvpgo

var ResponceStatusCode map[int]string

func init() {
	ResponceStatusCode = map[int]string{
		200: "200 - OK - リクエストは正常に成功しました。",
		204: "204 - No Content - リクエストは正常に成功しました。レスポンスはありません。",
		400: "400 - Bad Request - 不正なフォーマットでリクエストが行われました。",
		401: "401 - Unauthorized - Client IDまたはアクセストークンがリクエストに含まれていないか有効ではありません。",
		403: "403 - Forbidden - 認可されたスコープ範囲外へのリクエストです。",
		404: "404 - Not Found - リクエストされたリソースが見つかりませんでした。",
		429: "429 - Too Many Requests - レート制限が適用されました。",
		500: "500 - Internal Server Error - APIサーバで問題が発生しています。",
		502: "502 - Bad Gateway - 一時的にAPIサーバに接続出来なくなっています。",
		503: "503 - Server Unavailable - 一時的にAPIサーバに接続出来なくなっています。",
	}
}
