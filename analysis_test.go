package yingyan

import (
	"fmt"
	"testing"
	"github.com/cnjack/yingyan.sdk"
	"time"
)
/*
停留点测试
 */
func TestS_StayPoint(t *testing.T) {
	start := time.Now().Unix() - 6300
	client := yingyan.NewClient("do0znBH8Du2YflrDtZ6osFGVXOEG3osi", "fOBSnqpZyppQ2ImXP7HfqFeeY17MouzZ", 135928)
	body, err := client.StayPoint("run_1", start, time.Now().Unix() ,600, &yingyan.ProcessOption{}, yingyan.BaiDuCoordType)
	if err != nil {
		t.Error("EntityAdd err :", err.Error())
	}
	fmt.Println(body)
}

/*
驾驶行为测试
 */
func TestS_DrivingBehaviour(t *testing.T) {
	client := yingyan.NewClient("do0znBH8Du2YflrDtZ6osFGVXOEG3osi", "fOBSnqpZyppQ2ImXP7HfqFeeY17MouzZ", 135928)
	body, err := client.DrivingBehavior("run_1", start, time.Now().Unix() , &yingyan.ProcessOption{}, yingyan.BaiDuCoordType)
	if err != nil {
		t.Error("EntityAdd err :", err.Error())
	}
	fmt.Println(body)
}
