package yingyan

import (
	"encoding/json"
	"strconv"
)

type CoordTypeOutput string

const (
	CoordDefault       CoordTypeOutput = `bd09ll`
	CoordGuoce         CoordTypeOutput = `gcj02`
	CoordWgs           CoordTypeOutput = `wgs84`
)
/**
停留点查询
 */
func (serv *s) StayPoint(entityName string, startTime, endTime int64,stayTime int, po *ProcessOption,coord_type_output CoordTypeOutput ) (r *CommonResp, err error) {
	if coord_type_output ==""{
		coord_type_output = CoordDefault
	}

	param := map[string]string{
		"entity_name":     entityName,
		"start_time":      strconv.FormatInt(startTime, 10),
		"end_time":        strconv.FormatInt(endTime, 10),
		"stay_time":       strconv.Itoa(stayTime),
		"process_option":  po.ToData(),
		"coord_type_output": string(coord_type_output),
	}

	respByte, err := serv.Get(analysisStayPoint, param)

	if err != nil {
		return nil, err
	}
	r = &CommonResp{}
	err = json.Unmarshal(respByte, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

/*
驾驶行为分析
 */
func (serv *s) DrivingBehaviour(entityName string, startTime, endTime int64, speedingThreshold float64, po *ProcessOption,coord_type_output CoordTypeOutput) (r *CommonResp, err error) {
	if coord_type_output ==""{
		coord_type_output = CoordDefault
	}

	param := map[string]string{
		"entity_name":     entityName,
		"start_time":      strconv.FormatInt(startTime, 10),
		"end_time":        strconv.FormatInt(endTime, 10),
		"process_option":  po.ToData(),
		"coord_type_output": string(coord_type_output),
	}

	respByte, err := serv.Get(analysisDrivingBehavior, param)

	if err != nil {
		return nil, err
	}
	r = &CommonResp{}
	err = json.Unmarshal(respByte, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}