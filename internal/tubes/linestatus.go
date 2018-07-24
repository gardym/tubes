package tubes

import (
  "fmt"

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
  "overground": RGBColor{239, 123, 16},
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

func getStatusColor(status string) RGBColor {
  switch status {
  case "Severe Delays", "Part Suspended":
    return RGBColor{255, 0, 0}
  case "Minor Delays":
    return RGBColor{255, 255, 40}
  // "Good Service" is included as default
  default:
    return RGBColor{0, 125, 50}
  }
}

func boldString(s string) string {
  return fmt.Sprintf("\033[1m%v\033[0m", s)
}
func printTubeLine(line LineStatus) {
  lineColor := getLineColor(line.Id)
  statusColor := getStatusColor(line.Status)

  if line.Reason != "" {
    fmt.Printf("\n")
  }

  lineString := rgbterm.BgString(boldString(line.Name), lineColor.R, lineColor.G, lineColor.B)
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
