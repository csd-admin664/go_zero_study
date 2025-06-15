package main

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
	"log"
)

/**
创建钱包 ：
1. 创建私钥   256 位随机数， 私钥
2. 获取公钥 使用私钥通过椭圆曲线乘法计算出的点（x, y）坐标
3. 获取地址  对公钥进行 Keccak-256 哈希，取最后 20 字节，转换为十六进制字符串
*/

func main() {
	//生成私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalln("Failed to generate private key: %v", err)
	}
	//转换为自字节 32 位 调用这个会自动往高位补0
	privateKeyByte := crypto.FromECDSA(privateKey)
	// 转带有0x的16进制 截取到ox 的到私钥地址
	fmt.Println("私钥地址", hexutil.Encode(privateKeyByte)[2:])

	//利用私钥生成公钥
	publicKey := privateKey.Public()
	//断言
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	//转换为字节
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	//转16进制 去掉前四位
	fmt.Println("from pubKey:", hexutil.Encode(publicKeyBytes)[4:]) // 去掉'0x04'

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("地址是", address)
	//end
	//公共地址其实就是公钥的Keccak-256哈希，然后我们取最后40个字符（20个字节）并用“0x”作为前缀。 以下是使用 golang.org/x/crypto/sha3 的 Keccak256函数手动完成的方法。
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println("full:", hexutil.Encode(hash.Sum(nil)[:]))
	fmt.Println("create By hash", hexutil.Encode(hash.Sum(nil)[12:]))

}
