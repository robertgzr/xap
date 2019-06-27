package radio

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
	"time"
)

const base string = "https://r-a-d.io"

var (
	url    = base + "/api"
	stream = base + "/main.mp3"
)

func StreamURL() string {
	return stream
}

type Result struct {
	Status Status `json:"main"`
}

type Status struct {
	NowPlaying string `json:"np"`
	Listeners  int    `json:"listeners"`
	IsAfk      bool   `json:"isafkstream"`
	CurrentPos int64  `json:"current"`
	StartPos   int64  `json:"start_time"`
	EndPos     int64  `json:"end_time"`
	Pos        Position
	Dj         DJ      `json:"dj"`
	Queue      []QItem `json:"queue"`
	LastPlayed []QItem `json:"lp"`
}

type DJ struct {
	ID   int    `json:"id"`
	Name string `json:"djname"`
}

type QItem struct {
	Meta      string `json:"meta"`
	Time      string `json:"-"`
	Timestamp int64  `json:"timestamp"`
}

type Position struct {
	Ok          bool
	Start       time.Time
	End         time.Time
	Now         time.Time
	Len         time.Duration
	Current     time.Duration
	CurrentPerc float64
}

func apiRequest() (*Result, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var r Result
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, err
	}

	pos := Position{
		Ok:    true,
		Start: time.Unix(r.Status.StartPos, 0),
		End:   time.Unix(r.Status.EndPos, 0),
		Now:   time.Unix(r.Status.CurrentPos, 0),
	}
	pos.Len = pos.End.Sub(pos.Start)
	if pos.Len.Nanoseconds() < 1 {
		pos.Ok = false
		r.Status.Pos = pos
		return &r, nil
	}
	pos.Current = pos.Now.Sub(pos.Start)
	pos.CurrentPerc = (100 / float64(pos.Len)) * float64(pos.Current)
	r.Status.Pos = pos
	return &r, nil
}

const tmpl = `r-a-d.io:
| {{ .NowPlaying }} {{ if .Pos.Ok }}
| {{ .Pos.Current }} / {{ .Pos.Len }} ({{ printf "%.2f%%" .Pos.CurrentPerc }}){{ end }}
|
| {{ .Dj.Name }} {{ if .IsAfk }}(afk){{ end }}
| Listeners: {{ .Listeners }}
|
| Last Played:{{ range .LastPlayed }}
|   * {{ .Meta }}{{ end }}
| Queue:{{ range .Queue }}
|   * {{ .Meta }}{{ end }}
`

// | Last: {{ with index .LastPlayed 0 }}{{ .Meta }}{{ end }}
// | Next: {{ with index .Queue 0 }}{{ .Meta }}{{ end }}

func PrintStatus() error {
	r, err := apiRequest()
	if err != nil {
		return err
	}
	t := template.Must(template.New("now").Parse(tmpl))
	return t.Execute(os.Stdout, r.Status)
}
