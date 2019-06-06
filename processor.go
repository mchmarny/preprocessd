package main

import (
	"time"
)

type processedEvent struct {

	// from mocked
	SourceID     string    `json:"source_id"`
	EventID      string    `json:"event_id"`
	EventTs      time.Time `json:"event_ts"`
	Label        string    `json:"label"`
	MemUsed      float64   `json:"mem_used"`
	CPUUsed      float64   `json:"cpu_used"`
	Load1        float64   `json:"load_1"`
	Load5        float64   `json:"load_5"`
	Load15       float64   `json:"load_15"`
	RandomMetric float64   `json:"random_metric"`

	// new
	CPULoadBucket string  `json:"cpu_load_bucket"`
	MemLoadBucket string  `json:"mem_load_bucket"`
	UtilBias      string  `json:"util_bias"`
	CombinedUtil  float64 `json:"combined_util"`
	LoadTrend     int32   `json:"load_trend"`
}

func process(m *mockedEvent) (p *processedEvent, err error) {

	logger.Print("Init Processor...")

	// first reload process event from mocked
	d := &processedEvent{
		SourceID:     m.SourceID,
		EventID:      m.EventID,
		EventTs:      m.EventTs,
		Label:        m.Label,
		MemUsed:      m.MemUsed,
		CPUUsed:      m.CPUUsed,
		Load1:        m.Load1,
		Load5:        m.Load5,
		Load15:       m.Load15,
		RandomMetric: m.RandomMetric,
	}

	// add bucket sizing for cpu and mem
	d.CPULoadBucket = bucketSize100(m.CPUUsed)
	d.MemLoadBucket = bucketSize100(m.MemUsed)

	// calc combined util
	d.CombinedUtil = (m.CPUUsed + m.MemUsed + m.Load1) / 3

	// utilization bias
	if m.CPUUsed > m.MemUsed {
		d.UtilBias = "cpu"
	} else {
		d.UtilBias = "ram"
	}

	// load trending
	if m.Load1 < m.Load5 && m.Load5 < m.Load15 {
		d.LoadTrend = -1
	} else if m.Load1 > m.Load5 && m.Load15 > m.Load5 {
		d.LoadTrend = 1
	} else {
		d.LoadTrend = 0
	}

	return d, nil
}

func bucketSize100(v float64) string {
	if v < 10.00 {
		return "small"
	} else if v > 10.00 && v < 75.00 {
		return "medium"
	} else if v > 75.00 {
		return "large"
	} else {
		logger.Fatal("Check if clauses on float")
	}
	return "na"
}
