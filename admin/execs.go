package main

import (
	"database/sql"
)

type TempStats struct {
	backType string
	times    int
	total    int
}

func GetAllResponses() (Stats, []dbResponse, error) {
	db, err := sql.Open("mysql", "root:root@tcp(172.18.0.2:3306)/benchmarks?parseTime=true&charset=utf8")
	if err != nil {
		return Stats{}, nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM benchmarks")
	if err != nil {
		return Stats{}, nil, err
	}
	// var stat Stats
	s := make(map[string]*TempStats)

	var resps []dbResponse
	for rows.Next() {
		var resp dbResponse
		err = rows.Scan(&resp.ID, &resp.BackType, &resp.ExecTime, &resp.CallDate)
		if err != nil {
			return Stats{}, nil, err
		}
		// s[resp.BackType].total += resp.ExecTime
		// s[resp.BackType].times += 1
		_, ok := s[resp.BackType]
		if !ok {
			s[resp.BackType] = &TempStats{}
		}
		temp := s[resp.BackType]
		temp.times += 1
		temp.total += resp.ExecTime
		s[resp.BackType] = temp
		resps = append(resps, resp)
	}

	var stat Stats

	stat.MeanTimes = make(map[string]int)

	for k, v := range s {
		stat.MeanTimes[k] = v.total / v.times
	}

	return stat, resps, nil
}
