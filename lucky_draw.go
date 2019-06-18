package main

import (
	mysqlWrapper "./mysql"
	redisWrapper "./redis"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"math/rand"
	"strconv"
	"time"
)


func GetAllAwardBatch() []AwardBatch{

	var awardBatches []AwardBatch

	// 获取所有奖品数
	conn := redisWrapper.GetConn()

	if conn == nil {
		log.Println("conn is nil")
		return awardBatches
	}

	defer conn.Close()

	awardInfoKey := getAwardInfoKey()
	values , err := redis.Values(conn.Do("ZRANGE",awardInfoKey,0,-1,"WITHSCORES"))

	if err != nil || len(values) == 0 {
		log.Println("get all award redis error ", err)
	}

	for index, value := range values {

		if index % 2 == 0 {

			awardName , ok := value.([]byte)

			if !ok {
				log.Println("value type error : ", value)
				continue
			}
			awardBatches = append(awardBatches, AwardBatch{
				name:string(awardName),
			})

		} else {

			lastUpdateTimeStr , ok := value.([]byte)
			if !ok {
				log.Println("time type error : ", lastUpdateTimeStr)
				continue
			}

			lastUpdateTime ,err := strconv.ParseInt(string(lastUpdateTimeStr), 10, 64)

			if err != nil {
				log.Println("time type error", err)
				continue
			}

			awardBatches[index/2].SetUpdateTime(lastUpdateTime)
		}
	}

	// 填充 totalAmount
	for index, awardBatch := range awardBatches {
		awardBatches[index].SetTotalAmount(Conf.AwardMap[awardBatch.GetName()])
	}

	return awardBatches
}

func WinPrize(username string) *AwardBatch {

	awardBatch := GetAwardBatch()

	if awardBatch == nil {
		return awardBatch
	}

	// 更新 lastUpdateTime 和 balance， 如果更新不成功，视为抽奖失败
	conn := redisWrapper.GetConn()

	if conn == nil {
		log.Println("conn is nil")
		return nil
	}

	defer conn.Close()

	conn.Send("WATCH", getAwardBalanceKey())
	conn.Send("MULTI")
	conn.Send("ZADD",getAwardInfoKey(),time.Now().Unix(), awardBatch.GetName())
	conn.Send("HSET",getAwardBalanceKey(), awardBatch.GetName(), awardBatch.totalBalance - 1)
	conn.Send("EXEC")

	err := conn.Flush()
	if err != nil {
		log.Println("redis error, " , err)
		return nil
	}

	log.Println("congratulations , you won ", awardBatch.GetName() )

	// 保存用户中奖纪录
	awardTime := time.Unix(awardBatch.GetUpdateTime(), 0).Format("2006-01-02 15:04:05")
//	userName := req.Form.Get("user_name")
	mysqlWrapper.SaveRecords(awardBatch.GetName() , awardTime, username)

	return awardBatch

}


func GetAwardBatch() *AwardBatch {

	awardBatch := RandomGetAwardBatch()

	if awardBatch == nil {
		log.Println("sorry, you didn't win the prize.")
		return nil
	}

	// 判断是否到达奖品释放时间点
	startTime , _ := ParseStringToTime(Conf.Award.StartTime)
	endTime , _ := ParseStringToTime(Conf.Award.EndTime)
	totalAmount := awardBatch.totalAmount
	totalBalance := awardBatch.totalBalance
	lastUpdateTime := awardBatch.GetUpdateTime()
	random := rand.New(rand.NewSource(lastUpdateTime))

	detaTime := (endTime - startTime) / awardBatch.totalAmount

	// 计算下一个奖品释放的时间点
	releaseTime := startTime + (totalAmount - totalBalance) * detaTime + int64(random.Int()) % detaTime

	log.Println("relaseTime : ", time.Unix(releaseTime, 0).Format("2006-01-02 15:04:05"))

	if time.Now().Unix() < releaseTime {
		// 未到达奖品释放的时间点，即不中奖
		log.Println("sorry, you didn't win the prize")
		return nil
	}

	return awardBatch
}

func RandomGetAwardBatch() *AwardBatch {

	conn := redisWrapper.GetConn()

	if conn == nil {
		log.Println("conn is nil")
		return nil
	}

	defer conn.Close()

	retMap, err := redis.Int64Map(conn.Do("HGETALL", getAwardBalanceKey()))

	if err != nil || retMap == nil {
		log.Println("redis HGETALL award error", err)
		return nil
	}

	totalBalance := int64(0)
	for _, value := range retMap {
		totalBalance += value
	}

	fmt.Println("retMap : ", retMap)

	if totalBalance == 0 {
		log.Println("total balance is 0")
		return nil
	}

	log.Println("totalBalance :", totalBalance)

	awardBatches := GetAllAwardBatch()

	for index , awardBatch := range awardBatches {
		awardBatches[index].totalBalance = retMap[awardBatch.GetName()]
	}

	log.Println("awardBatches :", awardBatches)

	random := rand.New(rand.NewSource(totalBalance))

	num := random.Int63n(totalBalance)

	for _ , awardBatch := range awardBatches {

		// 奖品已经抽完
		if awardBatch.GetTotalBalance() <= 0 {
			log.Println("奖品已经抽完")
			continue
		}

		num = num - awardBatch.GetTotalBalance()
		if num < 0 {
			return &awardBatch
		}
	}

	return nil
}


func InitAwardPool() {

	conn := redisWrapper.GetConn()

	if conn == nil {
		log.Println("conn is nil")
		return
	}

	defer conn.Close()

	conn.Send("ZADD", getAwardInfoKey(), time.Now().Unix(), "A")
	conn.Send("ZADD", getAwardInfoKey(), time.Now().Unix(), "B")
	conn.Send("ZADD", getAwardInfoKey(), time.Now().Unix(), "C")

	conn.Send("HSET", getAwardBalanceKey(), "A", Conf.AwardMap["A"])
	conn.Send("HSET", getAwardBalanceKey(), "B", Conf.AwardMap["B"])
	conn.Send("HSET", getAwardBalanceKey(), "C", Conf.AwardMap["C"])
	conn.Flush()

	for i := 0 ; i < 6; i++ {
		_ , err := conn.Receive()

		if err != nil {
			log.Println("conn send error", err)
		}
	}


}





