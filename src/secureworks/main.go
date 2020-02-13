package main

import (
  "net/http"
  "secureworks/constants"
  "encoding/json"
  "io"
  "io/ioutil"
  "fmt"
)

type RequestJSON struct {
  UserName string `json:"username"`
  UnixTimestamp uint `json:"unix_timestamp"`
  EventUUID string `json:"event_uuid"`
  IPAddress string `json:"ip_address"`
}

func handler(w http.ResponseWriter, req *http.Request) {
  if req.Method != http.MethodPost {
    http.Error(
      w, fmt.Sprintf("%s not supported", req.Method), 
      http.StatusNotImplemented,
    )
    return
  }
  body, err := ioutil.ReadAll(req.Body)
  defer req.Body.Close()
  if err != nil {
    http.Error(w, "could not process body", http.StatusInternalServerError)
    return
  }
  rJSON := RequestJSON{}
  err = json.Unmarshal(body, &rJSON)
  if err != nil {
    http.Error(w, "could not process body", http.StatusInternalServerError)
    return
  }
  w.WriteHeader(http.StatusOK)
  io.WriteString(w, string(body))
}

//go:generate /root/generate.sh
func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(constants.PORT_STR, nil)
}
