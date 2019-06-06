package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// requestHandler handles the HTTP request
func requestHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")

	var c pushedContent
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		logger.Printf("Error decoding message: " + err.Error())
		writeResp(w, http.StatusBadRequest, "Invalid Content")
		return
	}
	//logger.Printf("Raw pubsub pushed content: %v", c)

	var m mockedEvent
	if err := json.Unmarshal(c.Message.Data, &m); err != nil {
		logger.Printf("Error decoding message data: " + err.Error())
		writeResp(w, http.StatusBadRequest, "Invalid Content")
		return
	}
	logger.Printf("Message content: %v", m)

	d, err := process(&m)
	if err != nil {
		logger.Printf("Error processing data: " + err.Error())
		writeResp(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	logger.Printf("Processed data: %v", d)

	// err = que.push(r.Context(), data)
	// if err != nil {
	// 	logger.Printf("Error storing data: " + err.Error())
	// 	writeResp(w, http.StatusBadRequest, "Internal Error")
	// 	return
	// }

	writeResp(w, http.StatusOK, "OK")
	return
}

func writeResp(w http.ResponseWriter, status int, msg string) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(msg)
}

type pushedContent struct {
	Message struct {
		Attributes map[string]string
		Data       []byte
		ID         string `json:"message_id"`
	}
	Subscription string
}

type mockedEvent struct {
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
}
