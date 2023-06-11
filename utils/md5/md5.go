package md5

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5 MD5是一种哈希算法，它将任意长度的消息作为输入，经过一系列复杂的运算，输出一个固定长度的消息摘要。
func Md5(src []byte) string {
	m := md5.New()
	m.Write(src)
	res := hex.EncodeToString(m.Sum(nil))
	return res
}

// 但是由于 MD5 的输出长度固定为 128 位（32 个十六进制字符），而输入数据的长度可以是任意的，
//所以在理论上是可能存在多个不同的输入数据生成相同的 MD5 值的情况。‘

// 概率极低。所以目前暂且用这个。
// 也可以使用 SHA-256、SHA-512 来减小碰撞概率。
