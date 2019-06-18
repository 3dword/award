package log

import (
	"fmt"
	"log"
	"os"
	"time"
)

func InitLogConf() {
	y, m, d := time.Now().Date()
	fileName := fmt.Sprintf("%d-%d-%d_award.log", y, m, d)
	f , err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		log.Fatal("file open error, %v", err)
	}

	// 显示行号
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.SetOutput(f)
}