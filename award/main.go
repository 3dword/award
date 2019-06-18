package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

func main() {
	parse()
	awardBatches := InitAwardPool()
	fmt.Println(awardBatches)

	for {
		go drawPrize(awardBatches)
		time.Sleep(time.Duration(200)*time.Millisecond)
	}

}



func drawPrize(awardBatches []AwardBatch) {
	awardBatch , err := GetAward(awardBatches)


	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("很遗憾，您未能抽中奖品")
	}


	var mutex sync.Mutex

	defer mutex.Unlock()

	mutex.Lock()

	for i, award := range awardBatches {
		if strings.Compare(award.GetName(),awardBatch.GetName()) == 0 || awardBatches[i].totalBalance > 0 {
			awardBatches[i].totalBalance -= 1
			fmt.Println("恭喜您抽中奖品 : " + award.GetName())
		}
	}



}