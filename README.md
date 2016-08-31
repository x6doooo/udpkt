#udpkt

较长数据拆分/粘合，便于UDP通信使用

---


## Example


发送端

```
// 初始化udp客户端
remoteAddr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:8553")
client, _ := net.DialUDP("udp", nil, remoteAddr)

// 初始化encoder
encoder := udpkt.NewEncoder()

// 待发送的数据
message := []byte("[22/Aug/2016:17:36:58 +0800] fwf[-] www.xxx.com GET /api?Action=DescribeAddressesOfBws&Service=eip&Version=2016-03-04&Region=cn-shanghai-3&BwsId=eee6fc52-586e-44bb-ba21-ac06ef08787a&_=1471858571768 HTTP/1.1 200  0.768 0.768 680 1487 http://www.xxx.com/index.html Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.87 Safari/537.36 10.111.17.86 80 -")

// 将数据拆分成若干个分片
udp_packets := encoder.Do(message)

// 分片发送
for _, pack := range udp_packets {
	conn.Write(pack)
}
```

服务端

```
// 初始化udp服务端
ServerAddr, _ := net.ResolveUDPAddr("udp", ":8553")
ServerConn, err := net.ListenUDP("udp", ServerAddr)

// 初始化decoder
decoder := udpkt.NewDecoder()

// 接收数据
for {
	// 初始化buffer
	buf := make([]byte, 1024)
	// 接收udp数据
	n, addr, err := udp.ReadFromUDP(buf)
	// 粘包
	fullMsg, done := Decoder.Do(buf)
	// done为真，表示这个数据包的所有分片都已接收到
	if done {
		fmt.Println(string(fulMsg))
	}
}

```


