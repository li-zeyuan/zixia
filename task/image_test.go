package task

import "testing"

func TestImg2txt(t *testing.T) {
	img2txt("/Users/zeyuan.li/Desktop/workspace/code/src/github.com/li-zeyuan/sun/zixia/zixia.png",
		150, []string{"*", "%", "+", ",", ".", " "}, "\n", "./zixia")

}
