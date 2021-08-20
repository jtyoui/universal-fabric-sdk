package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"log"
	"strings"
)

// CurrencyChaincode 通用智能合约
type CurrencyChaincode struct {
	stub shim.ChaincodeStubInterface
	Key  string
}

func (t *CurrencyChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success([]byte("初始化成功"))
}

func (t *CurrencyChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	function, args := stub.GetFunctionAndParameters()
	cc := &CurrencyChaincode{
		stub: stub,
		Key:  args[0],
	}
	switch function {
	case "set":
		return cc.set(args[1])
	case "update":
		return cc.update(args[1])
	case "get":
		return cc.get()
	case "delete":
		return cc.delete()
	default:
		return shim.Error("无效请求！只能输入set、get、update、delete函数")
	}
}

// hashKey 判断是否存在该建
func (t *CurrencyChaincode) hashKey() bool {
	log.Println("输入的主键为：" + t.Key)
	value, err := t.stub.GetState(t.Key)
	if err != nil {
		return false
	}
	if value == nil {
		return false
	}
	if len(value) > 0 {
		return true
	}
	return false
}

func (t *CurrencyChaincode) update(value string) peer.Response {
	if !t.hashKey() {
		return shim.Error("输入的主键不存在")
	}
	if strings.Trim(value, " ") == "" {
		return shim.Error("插入的数据不能为空")
	}
	err := t.stub.PutState(t.Key, []byte(value))
	if err != nil {
		return shim.Error("更新失败： " + err.Error())
	}
	return shim.Success([]byte("更新成功！"))
}

func (t *CurrencyChaincode) set(value string) peer.Response {
	if strings.Trim(value, " ") == "" {
		return shim.Error("插入的数据不能为空")
	}
	if t.hashKey() {
		return shim.Error("输入的主键已存在")
	}
	err := t.stub.PutState(t.Key, []byte(value))
	if err != nil {
		return shim.Error("增加失败： " + err.Error())
	}
	return shim.Success([]byte("增加成功！"))
}

func (t *CurrencyChaincode) get() peer.Response {
	value, err := t.stub.GetState(t.Key)
	if err != nil {
		return shim.Error("查询失败，不存在主键ID：" + err.Error())
	}
	if value == nil {
		return shim.Error("查询失败： 因为值为空")
	}
	return shim.Success(value)
}

func (t *CurrencyChaincode) delete() peer.Response {
	if !t.hashKey() {
		return shim.Error("输入的主键不存在")
	}
	err := t.stub.DelState(t.Key)
	if err != nil {
		return shim.Error("删除失败：" + err.Error())
	}
	return shim.Success([]byte("删除成功！"))
}

func main() {
	if err := shim.Start(new(CurrencyChaincode)); err != nil {
		log.Fatalf("启动链表失败: %s", err)
	} else {
		log.Println("启动链表成功！")
	}
}
