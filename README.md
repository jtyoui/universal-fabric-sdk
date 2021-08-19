# universal-fabric-sdk

一个简单的万能链码调用

## 使用返回

    go get -u github.com/jtyoui/universal-fabric-sdk

```go
package main

import (
	"fmt"
	"github.com/jtyoui/universal-fabric-sdk"
)

func main() {
	config := &ConfigContract{
		ConfigDir: "config",   // 配置文件目录，放在配置文件的文件夹
		CertPath:  "cert.pem", // peerOrganizations/org1.example.com/users/User1@org1.example.com/msp/signcerts/xxcert.pem
		KeyPath:   "key_sk",   // peerOrganizations/org1.example.com/users/User1@org1.example.com/msp/keystore/xx_sk
		ChinaCode: "basic",
		Channel:   "mychannel",
		MSPId:     "Org1MSP",
	}
	contract := Contract(config)
	transaction, _ := contract.EvaluateTransaction("get", "1")
	fmt.Println(string(transaction))
}
```