package main


type AwardBatch struct {

	// 奖品名称
	name string

	// 剩余奖品总数
	totalBalance int64

	// 奖品总数
	totalAmount int64

	// 上次奖品发放时间
	updateTime int64

}

func (award *AwardBatch) GetName() string {
	return award.name
}

func (award *AwardBatch) GetTotalBalance() int64 {
	return award.totalBalance
}

func (award *AwardBatch) GetTotalAmount() int64 {
	return award.totalAmount
}

func (award *AwardBatch) GetUpdateTime() int64 {
	return award.updateTime
}

func (award *AwardBatch) SetTotalBalance(totalBalance int64)  {
	award.totalBalance = totalBalance
}

func (award *AwardBatch) SetTotalAmount(totalAmount int64)  {
	award.totalAmount = totalAmount
}

func (award *AwardBatch) SetUpdateTime(updateTime int64)  {
	award.updateTime = updateTime
}

func (award *AwardBatch) SetName(name string)  {
	award.name = name
}
