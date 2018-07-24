package tubes

import (
  "net/http"
  "io/ioutil"
  "encoding/json"

  "github.com/gardym/tubes/internal/tfl"
)

type LineStatus struct {
  Id string
  Name string
  Status string
  Reason string
}

func getDataHttp() []byte {
  resp, err := http.Get(tfl.StatusApiEndpoint)

  if err != nil {
    panic(err)
  }

  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)

  if err != nil {
    panic(err)
  }

  return body
}

func getDataFile() []byte {
  data, err := ioutil.ReadFile("../../sample-data.json")

  if err != nil {
    panic(err)
  }

  return data
}

func GetTubeStatus() (lineStatuses []LineStatus) {

  body := getDataHttp()
  //body := getDataFile()

  var tflLines []tfl.Line

  err := json.Unmarshal(body, &tflLines)

  if err != nil {
    panic(err)
  }

  for _, l := range tflLines {
    for _, ls := range l.LineStatuses {
      lineStatuses = append(lineStatuses, LineStatus {
        Id: l.Id,
        Name: l.Name,
        Status: ls.Status,
        Reason: ls.Reason,
      })
    }
  }

  return
}
