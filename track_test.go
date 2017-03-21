package yingyan_test

import (
	"fmt"
	"github.com/cnjack/yingyan.sdk"
	"testing"
	"time"
)

func TestS_AddPoint(t *testing.T) {
	client := yingyan.NewClient("do0znBH8Du2YflrDtZ6osFGVXOEG3osi", "fOBSnqpZyppQ2ImXP7HfqFeeY17MouzZ", 135928)
	body, err := client.AddPoint(&yingyan.PointData{
		EntityName:     "run_1",
		Latitude:       39.848434,
		Longitude:      116.492943,
		CoordTypeInput: yingyan.GPSCoordType,
	})
	if err != nil {
		t.Error("EntityAdd err :", err.Error())
	}
	fmt.Println(body)
}

func TestS_AddPoints(t *testing.T) {
	client := yingyan.NewClient("do0znBH8Du2YflrDtZ6osFGVXOEG3osi", "fOBSnqpZyppQ2ImXP7HfqFeeY17MouzZ", 135928)
	params := []*yingyan.PointData{}
	params = append(params, &yingyan.PointData{
		EntityName:     "run_1",
		Latitude:       39.845434,
		LocTime:        time.Now().Unix(),
		Longitude:      116.442943,
		CoordTypeInput: yingyan.GPSCoordType,
	})
	params = append(params, &yingyan.PointData{
		EntityName:     "run_1",
		Latitude:       39.845534,
		LocTime:        time.Now().Unix() - 60,
		Longitude:      116.495343,
		CoordTypeInput: yingyan.GPSCoordType,
	})
	body, err := client.AddPoints(params)
	if err != nil {
		t.Error("EntityAdd err :", err.Error())
	}
	fmt.Println(body)
}