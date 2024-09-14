package main

import (
	"net/http"
	"strconv"

	"github.com/skip2/go-qrcode"
)

func main() {
	http.HandleFunc("/qr", qrHandler)
	http.ListenAndServe(":8080", nil)
}

func qrHandler(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("text")
	if text == "" {
		http.Error(w, "Missing 'text' query parameter", http.StatusBadRequest)
		return
	}

	sizeParam := r.URL.Query().Get("size")
	size := 256
	if sizeParam != "" {
		var err error
		size, err = strconv.Atoi(sizeParam)
		if err != nil || size <= 0 {
			http.Error(w, "Invalid 'size' parameter", http.StatusBadRequest)
			return
		}
	}

	levelParam := r.URL.Query().Get("level")
	level := qrcode.Medium
	switch levelParam {
	case "low":
		level = qrcode.Low
	case "medium":
		level = qrcode.Medium
	case "high":
		level = qrcode.High
	case "highest":
		level = qrcode.Highest
	case "":

	default:
		http.Error(w, "Invalid 'level' parameter", http.StatusBadRequest)
		return
	}

	png, err := qrcode.Encode(text, level, size)
	if err != nil {
		http.Error(w, "Failed to generate QR code", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "image/png")
	w.Write(png)
}
