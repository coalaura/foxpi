package main

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"net/http"
	"os"
	"time"
)

const (
	MaxMessageLength = 4 * 1024 * 1024
)

var (
	receiver = make(chan Response)
)

type Request struct {
	Method  string            `json:"method"`
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}

type Response struct {
	Headers map[string]string `json:"headers"`
	Body    *string           `json:"body"`
}

func read() {
	var (
		length uint32
	)

	for {
		err := binary.Read(os.Stdin, binary.LittleEndian, &length)
		if err != nil {
			log.Log("error reading length: %s", err.Error())

			return
		}

		if length > MaxMessageLength {
			log.Log("message too large: %d", length)

			return
		}

		buffer := make([]byte, length)

		_, err = os.Stdin.Read(buffer)
		if err != nil {
			log.Log("error reading message: %s", err.Error())

			return
		}

		var response Response

		err = json.Unmarshal(buffer, &response)
		if err != nil {
			log.Log("error unmarshalling message: %s", err.Error())

			return
		}

		receiver <- response
	}
}

func request(req Request) (*Response, error) {
	message, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer

	err = binary.Write(&buf, binary.LittleEndian, uint32(len(message)))
	if err != nil {
		return nil, err
	}

	buf.Write(message)

	os.Stdout.Write(buf.Bytes())

	select {
	case response := <-receiver:
		return &response, nil
	case <-time.After(2 * time.Second):
		return nil, nil
	}
}

func (r *Response) Forward(w http.ResponseWriter) {
	if r.Headers != nil {
		for key, value := range r.Headers {
			w.Header().Set(key, value)
		}
	}

	var (
		body []byte
		err  error
	)

	if r.Body != nil {
		body, err = base64.StdEncoding.DecodeString(*r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}
	}

	w.Write(body)
}
