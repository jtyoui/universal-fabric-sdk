package sdk

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
	"io/ioutil"
	"log"
)

// ConfigContract 连接的配置文件，configDir文件夹必须要用一个连接文件：connection-org.yaml
type ConfigContract struct {
	ConfigDir string // 配置文件目录，放在配置文件的文件夹
	CertPath  string // peerOrganizations/org1.example.com/users/User1@org1.example.com/msp/signcerts/xxcert.pem
	KeyPath   string // peerOrganizations/org1.example.com/users/User1@org1.example.com/msp/keystore/xx_sk
	ChinaCode string // 链码名字
	Channel   string // 通道名字
	MSPId     string // MSPid名字，比如：Org1MSP
}

// Contract SDK连接
func Contract(ct *ConfigContract) *gateway.Contract {
	wallet, err := gateway.NewFileSystemWallet(ct.ConfigDir)
	if err != nil {
		log.Panicf("创建wallet失败: %s\n", err)
	}
	if !wallet.Exists("appUser") {
		err = populateWallet(wallet, ct)
		if err != nil {
			log.Panicf("构建wallet链接失败: %s\n", err)
		}
	}
	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(ct.ConfigDir+"/connection-org.yaml")),
		gateway.WithIdentity(wallet, "appUser"),
	)
	if err != nil {
		log.Panicf("连接gateway失败: %s\n", err)
	}
	defer gw.Close()
	network, err := gw.GetNetwork(ct.Channel)
	if err != nil {
		log.Panicf("连接network失败: %s\n", err)
	}
	contract := network.GetContract(ct.ChinaCode)
	return contract
}

// populateWallet 构建Wallet
func populateWallet(wallet *gateway.Wallet, ct *ConfigContract) error {
	cert, err := ioutil.ReadFile(ct.ConfigDir + "/" + ct.CertPath)
	if err != nil {
		log.Panicf("读取CertPath文件失败:%s\n", err)
	}
	key, err := ioutil.ReadFile(ct.ConfigDir + "/" + ct.KeyPath)
	if err != nil {
		log.Panicf("读取KeyPath文件失败:%s\n", err)
	}
	identity := gateway.NewX509Identity(ct.MSPId, string(cert), string(key))
	err = wallet.Put("appUser", identity)
	if err != nil {
		log.Panicf("解析X509失败:%s\n", err)
	}
	return nil
}
