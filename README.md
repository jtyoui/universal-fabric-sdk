# universal-fabric-sdk

一个简单的万能链码调用

## 前提

    创建一个目录放置配置
    mkdir config
    下载配置文件： https://github.com/jtyoui/universal-fabric-sdk/releases/download/v1.1/connection-org.yaml
    然后将下载的文件复制到目录 cp connection-org.yaml ./config

## 使用方法

    go get -u github.com/jtyoui/universal-fabric-sdk

```go
package main

import (
	"fmt"
	"github.com/jtyoui/universal-fabric-sdk"
)

func main() {
	config := &ConfigContract{
		ConfigDir: "./config", // 配置文件目录，放在配置文件的文件夹,改文件夹目录必须包含：connection-org.yaml
		CertPath:  "cert.pem", // peerOrganizations/org1.example.com/users/User1@org1.example.com/msp/signcerts/xxcert.pem
		KeyPath:   "key_sk",   // peerOrganizations/org1.example.com/users/User1@org1.example.com/msp/keystore/xx_sk
		ChinaCode: "basic",
		Channel:   "channel",
		MSPId:     "Org1MSP",
	}
	contract := Contract(config)
	transaction, _ := contract.EvaluateTransaction("get", "1")
	fmt.Println(string(transaction))
}
```