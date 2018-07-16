package main

import (
  "net/http"
  "io/ioutil"
  "fmt"
  "encoding/json"
  "github.com/fatih/color"
)

func getDataHttp() []byte {
  resp, err := http.Get("https://api.tfl.gov.uk/line/mode/tube/status")

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

type LineStatus struct {
  id string
  name string
  status string
  reason string
}

func getLineColor(id string) *color.Color {
  switch id {
  case "central":
    return color.New(color.BgHiRed)
  case "circle":
    return color.New(color.BgYellow)
  case "piccadilly":
    return color.New(color.BgBlue)
  case "district":
    return color.New(color.BgGreen)
  case "metropolitan":
    return color.New(color.BgMagenta)
  case "waterloo-city":
    return color.New(color.BgCyan)
  case "victoria":
    return color.New(color.BgHiBlue)
  case "hammersmith-city":
    return color.New(color.BgHiMagenta)
  case "jubilee":
    return color.New(color.BgHiBlack)
  case "northern":
    return color.New(color.BgBlack)
  case "bakerloo":
    return color.New(color.BgRed)
  default:
    return color.New(color.BgWhite).Add(color.FgBlack)
  }
}

func getStatusColor(status string) *color.Color {
  switch status {
  case "Severe Delays":
    return color.New(color.FgRed)
  case "Minor Delays":
    return color.New(color.FgYellow)
  // "Good Service" is included as default
  default:
    return color.New(color.FgHiBlack)
  }
}

func printTubeLine(line LineStatus) {
  lineColor := getLineColor(line.id).Add(color.Bold).SprintFunc()
  statusColor := getStatusColor(line.status).SprintFunc()

  fmt.Printf("%v: %v\n", lineColor(line.name), statusColor(line.status))

  if line.reason != "" {
    fmt.Printf("%v\n\n", line.reason)
  }
}

func main() {
  //body := getDataHttp()
  body := getDataFile()

  var dat []interface{}

  err := json.Unmarshal(body, &dat)

  if err != nil {
    panic(err)
  }

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

      printTubeLine(l)
    }
  }
}
