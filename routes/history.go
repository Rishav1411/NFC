package routes

import (
	"encoding/json"
	"net/http"
	"nfc/m/database"
	"nfc/m/database/operations"
)

func history(w http.ResponseWriter, r *http.Request) {
	db := database.SQLConnection()
	if db == nil {
		ServerError(w)
		return
	}
	defer db.Close()
	val := operations.CheckWallet(r.Context().Value(userContextKey).(int64), db)
	if val == -2 {
		ServerError(w)
		return
	}
	if val == -1 {
		jsonData, _ := json.Marshal(map[string]interface{}{
			"details": "wallet is not associated with user",
		})
		WriteJson(w, jsonData, 400)
		return
	}
	var logs []HistoryRow
	rows := operations.GetHistory(val, db)
	if rows == nil {
		ServerError(w)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var log HistoryRow
		err := rows.Scan(&log.Sender_id, &log.Receiver_id, &log.Amount, &log.Date, &log.Time)
		if err != nil {
			ServerError(w)
			return
		}
		logs = append(logs, log)
	}
	if rows.Err() != nil {
		ServerError(w)
		return
	}
	jsonData, _ := json.Marshal(map[string]interface{}{
		"details": logs,
	})
	WriteJson(w, jsonData, 200)
}
