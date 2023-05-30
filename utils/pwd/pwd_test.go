package pwd

import (
	"fmt"
	"testing"
)

func TestHashPwd(t *testing.T) {
	fmt.Println(HashPwd("1234"))
}

func TestCheckPwd(t *testing.T) {
	fmt.Println(CheckPwd("$2a$04$gp3pPvolVfo12ZQbiArA3eWGqr1Mb6Hm38VGrRE.6hayUsJUMmv4.", "1234"))
}
