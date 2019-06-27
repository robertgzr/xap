package stew

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
	"time"
)

const baseURL = "https://radio.stew.moe"

var (
	stream       = baseURL + "/stream/stream256.opus"
	api          = baseURL + "/api"
	apiPlaying   = api + "/playing"
	apiListeners = api + "/listeners"
)

func StreamURL() string {
	return stream
}

type Result struct {
	Title     string `json:"title"`
	Album     string `json:"album"`
	Artist    string `json:"artist"`
	DurationS int    `json:"duration"`
	Duration  time.Duration
	Listeners int `json:"num_listeners"`
	Queue     []QItem
}

type QItem struct {
	Title string
}

func apiRequest() (*Result, error) {
	resp, err := http.Get(apiPlaying)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var r Result
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return nil, err
	}

	r.Duration = time.Duration(r.DurationS) * time.Second
	return &r, nil
}

const tmpl = `stew.moe:
| {{ .Title }} by {{ .Artist }}
{{- with .Album }}{{ . }}{{ end -}}
|
| Listeners: {{ .Listeners }}
|
| Queue:{{ range .Queue }}
|   * {{ .Title }}{{ end }}
`

func PrintStatus() error {
	now, err := apiRequest()
	if err != nil {
		return err
	}
	t := template.Must(template.New("now").Parse(tmpl))
	return t.Execute(os.Stdout, now)
}
