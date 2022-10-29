package Checker

import (
	"crypto/md5"
	"fmt"
	"os"
)

func FileCheker(path string) {
	dat, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("ERR READ FILE")
		os.Exit(0)
	}

	correctHash := [16]byte{172, 133, 232, 49, 39, 228, 158, 196, 36, 135, 242, 114, 217, 185, 219, 139}
	hash := md5.Sum([]byte(string(dat)))
	if hash != correctHash {
		fmt.Println("Error! file has been modified!")
		os.Exit(0)
	}
}
