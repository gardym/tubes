package tubes

import (
  "fmt"
  "github.com/fatih/color"
)

var lineColors = map[string]*color.Color {
  "central": color.New(color.BgHiRed),
  "circle": color.New(color.BgYellow),
  "piccadilly": color.New(color.BgBlue),
  "district": color.New(color.BgGreen),
  "metropolitan": color.New(color.BgMagenta),
  "waterloo-city": color.New(color.BgCyan),
  "victoria": color.New(color.BgHiBlue),
  "hammersmith-city": color.New(color.BgHiMagenta),
  "jubilee": color.New(color.BgHiBlack),
  "northern": color.New(color.BgBlack),
  "bakerloo": color.New(color.BgRed),
}

func getLineColor(id string) *color.Color {
  c, ok := lineColors[id]

  if ok {
    return c
  }

  return color.New(color.BgWhite).Add(color.FgBlack)
}

func getStatusColor(status string) *color.Color {
  switch status {
  case "Severe Delays", "Part Suspended":
    return color.New(color.FgRed)
  case "Minor Delays":
    return color.New(color.FgYellow)
  // "Good Service" is included as default
  default:
    return color.New(color.FgGreen)
  }
}

func printTubeLine(line LineStatus) {
  lineColor := getLineColor(line.id).Add(color.Bold).SprintFunc()
  statusColor := getStatusColor(line.status).SprintFunc()

  if line.reason != "" {
    fmt.Printf("\n")
  }

  fmt.Printf("%v: %v\n", lineColor(line.name), statusColor(line.status))

  if line.reason != "" {
    fmt.Printf("%v\n\n", line.reason)
  }
}

func PrintLineStatus() {
  for _, s := range GetTubeStatus() {
    printTubeLine(s)
  }
}
