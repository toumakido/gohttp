# gohttp
netパッケージでhttpサーバを実装（net/httpは使用しない）

netパッケージは以下を提供
network I/O, 
* TCP/IP, UDP, 
* domain name resolution,
* Unix domain sockets

curlでhttpリクエスト
http://localhost:8080

httpの役割について整理
header,bodyをつけて送っている

サーバー
method: GET(パラメータなし)
tcpのlistenしておく
acceptしたら、決められたbodyとheaderをつけて
送信

tcpの