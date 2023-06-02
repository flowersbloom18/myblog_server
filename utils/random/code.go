package random

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

func Code(length int) string {
	rand.Seed(time.Now().UnixNano())

	// 10的几次方，10^4=10000，表示四位验证码
	power := int(math.Pow(10, float64(length)))
	// 生成对应长度的验证码
	return fmt.Sprintf("%"+(strconv.Itoa(length))+"v", rand.Intn(power))
}
