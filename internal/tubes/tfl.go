package tubes

import (
  "net/http"
  "io/ioutil"
  "encoding/json"
)

const statusApiEndpoint = "https://api.tfl.gov.uk/line/mode/tube/status"

type LineStatus struct {
  id string
  name string
  status string
  reason string
}

func getDataHttp() []byte {
  resp, err := http.Get(statusApiEndpoint)

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

func GetTubeStatus() []LineStatus {

  body := getDataHttp()
  //body := getDataFile()

  var dat []interface{}

  err := json.Unmarshal(body, &dat)

  if err != nil {
    panic(err)
  }

  var statuses []LineStatus

  for _, v := range dat {
    o := v.(map[string]interface{})

    for _, s := range o["lineStatuses"].([]interface{}) {
      st := s.(map[string]interface{})

      l := LineStatus {
        id: o["id"].(string),
        name: o["name"].(string),
        status: st["statusSeverityDescription"].(string),
      }

      reason, ok := st["reason"]
      if ok {
        l.reason = reason.(string)
      }

      statuses = append(statuses, l)
    }
  }

  return statuses
}
