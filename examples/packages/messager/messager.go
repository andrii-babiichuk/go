package messager

import (
	"fmt"
	"time"
)

var initTime int64

func init() {
	initTime = time.Now().Unix()
}

func Hello() string {
	return fmt.Sprintf(msg, initTime)
}
