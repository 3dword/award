package main

import (
	logger "./log"
	"fmt"
	"log"
)

func main() {
	parse()
	logger.InitLogConf()
	InitAwardPool()

	//for {
	//	go WinPrize("diuge")
	//	time.Sleep(time.Duration(200)*time.Millisecond)
	//}

	Start()
}



func drawPrize(awardBatches []AwardBatch) {
	awardBatch := WinPrize("diuge")


	if awardBatch == nil {
		fmt.Println("很遗憾，您未能抽中奖品")
	}

	log.Println("恭喜您抽中:",awardBatch.GetName())

	//var mutex sync.Mutex
	//
	//defer mutex.Unlock()
	//
	//mutex.Lock()
	//
	//for i, award := range awardBatches {
	//	if strings.Compare(award.GetName(),awardBatch.GetName()) == 0 || awardBatches[i].totalBalance > 0 {
	//		awardBatches[i].totalBalance -= 1
	//		fmt.Println("恭喜您抽中奖品 : " + award.GetName())
	//	}
	//}



}