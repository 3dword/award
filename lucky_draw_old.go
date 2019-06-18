package main


/**
	废弃版本，仅用于演示思路二

	通过 toml 配置文件来实现 mysql 效果

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var awardBatches []AwardBatch


func GetAllAward(awardBatches []AwardBatch) (AwardBatch, error) {

	startTime , _ := ParseStringToTime(conf.Award.StartTime)
	endTime , _ := ParseStringToTime(conf.Award.EndTime)

	fmt.Println("startTime : ", startTime)
	fmt.Println("endTime : ", endTime)

	award , err := RandomGetAwardBatch(awardBatches)
	if err != nil {
		return AwardBatch{}, err
	}


	totalAmount := award.GetTotalAmount()
	totalBalance := award.GetTotalBalance()
	updateTime := award.GetUpdateTime()

	detaTime := (endTime - startTime) / totalAmount
	currentTime := time.Now().Unix()

	random := rand.New(rand.NewSource(updateTime))
	// 计算下一个奖品的释放时间
	releaseTime := startTime + (totalAmount - totalBalance) * detaTime +  int64(random.Int()) % detaTime

	fmt.Println("releaseTime : " + fmt.Sprintf("%d", releaseTime) + " currentTime : " + fmt.Sprintf("%d",currentTime))

	if (currentTime < releaseTime) {
		return AwardBatch{} , errors.New(" currentTime not in award release period ")
	}

	return award, nil
}


func RandomGetAwardBatch(awardBatches []AwardBatch) ( AwardBatch , error ) {

	if len(awardBatches) == 0 {
		return AwardBatch{} , errors.New("empty param awardBatches")
	}

	totalBalance := int64(0)

	for _, awardBatch := range awardBatches {
		totalBalance += awardBatch.GetTotalBalance()
	}

	if totalBalance == 0 {
		return AwardBatch{}, errors.New("weight is 0")
	}

	r := rand.New(rand.NewSource(totalBalance))

	num := r.Int63n(totalBalance)

	for _, awardBatch := range awardBatches {
		num -= awardBatch.GetTotalBalance()

		if num < 0 {
			return awardBatch , nil
		}
	}

	return AwardBatch{}, errors.New("randomGetAwardBatch should shoot at least one batch")

}


func InitAwardPool() []AwardBatch {
	startTime , _ := ParseStringToTime(conf.Award.StartTime)
	awardBatches = make([]AwardBatch , 3)
	awardBatches[0].SetName("A")
	awardBatches[0].SetTotalBalance(conf.Award.A)
	awardBatches[0].SetTotalAmount(conf.Award.A)
	awardBatches[0].SetUpdateTime(startTime)

	awardBatches[1].SetName("B")
	awardBatches[1].SetTotalBalance(conf.Award.B)
	awardBatches[1].SetTotalAmount(conf.Award.B)
	awardBatches[1].SetUpdateTime(startTime)

	awardBatches[2].SetName("C")
	awardBatches[2].SetTotalBalance(conf.Award.C)
	awardBatches[2].SetTotalAmount(conf.Award.C)
	awardBatches[2].SetUpdateTime(startTime)

	return awardBatches
}

 */