package tubes

import (
  "fmt"
  "strings"

  "github.com/aybabtme/rgbterm"
)

type RGBColor struct {
  R uint8
  G uint8
  B uint8
}

// http://content.tfl.gov.uk/tfl-colour-standards-issue04.pdf
var lineColors = map[string]RGBColor {
  "central": RGBColor{220, 36, 31},
  "circle": RGBColor{255, 211, 41},
  "piccadilly": RGBColor{0, 25, 168},
  "district": RGBColor{0, 125, 50},
  "metropolitan": RGBColor{155, 0, 88},
  "waterloo-city": RGBColor{147, 206, 186},
  "victoria": RGBColor{0, 152, 216},
  "hammersmith-city": RGBColor{244, 169, 190},
  "jubilee": RGBColor{161, 165, 167},
  "northern": RGBColor{0, 0, 0},
  "bakerloo": RGBColor{178, 99, 0},
  "london-overground": RGBColor{239, 123, 16},
  "elizabeth": RGBColor{147, 100, 204},
  "dlr": RGBColor{0, 175, 173},
}

func getLineColor(id string) RGBColor {
  c, ok := lineColors[id]

  if ok {
    return c
  }

  return RGBColor{0, 0, 0}
}

// https://api.tfl.gov.uk/Line/Meta/Severity
func getStatusColor(status string) RGBColor {
  switch status {
  case "Good Service":
    return RGBColor{0, 165, 50}
  case "Severe Delays", "Closed", "Suspended",
    "Planned Closure", "Part Closure", "Part Closed",
    "Not Running", "Service Closed":
    return RGBColor{255, 0, 0}
  default:
    return RGBColor{255, 255, 40}
  }
}

func boldString(s string) string {
  return fmt.Sprintf("\033[1m%v\033[0m", s)
}

func darken(c RGBColor) RGBColor {
  if c.R == 0 && c.G == 0 && c.B == 0 {
    return RGBColor{255, 255, 255}
  }

  var p uint8 = 3
  return RGBColor{c.R / p, c.G / p, c.B / p}
}

func padRight(s string, n int) string {
  return fmt.Sprintf("%v%v", s, strings.Repeat(" ", n - len([]rune(s))))
}

func printTubeLine(line LineStatus) {
  lineColor := getLineColor(line.Id)
  statusColor := getStatusColor(line.Status)

  if line.Reason != "" {
    fmt.Printf("\n")
  }

  paddedName := fmt.Sprintf(" %v", padRight(line.Name, 19))

  d := darken(lineColor)
  lineName := rgbterm.FgString(paddedName, d.R, d.G, d.B)

  lineString := rgbterm.BgString(boldString(lineName), lineColor.R, lineColor.G, lineColor.B)
  statusString := rgbterm.FgString(line.Status, statusColor.R, statusColor.G, statusColor.B)

  fmt.Printf("%v: %v\n", lineString, statusString)

  if line.Reason != "" {
    fmt.Printf("%v\n\n", line.Reason)
  }
}

func PrintLineStatus() {
  for _, s := range GetTubeStatus() {
    printTubeLine(s)
  }
}
