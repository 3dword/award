package main

import (
	"time"
)

func ParseStringToTime(tm string) (int64, error)  {

	loc , _ := time.LoadLocation("Local")
	resultTime , err := time.ParseInLocation("2006-01-02 15:04:05", tm, loc)

	return resultTime.Unix() , err

}