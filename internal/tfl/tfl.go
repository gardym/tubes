package tfl

const StatusApiEndpoint = "https://api.tfl.gov.uk/line/mode/tube,overground,dlr/status"

type Line struct {
  Id string `json:"id"`
  Name string `json:"name"`
  LineStatuses []LineStatus `json:"lineStatuses"`
}

type LineStatus struct {
  Status string `json:"statusSeverityDescription"`
  Reason string `json:"reason"`
}
