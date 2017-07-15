package radio

import (
	"encoding/json"
	"net/http"
	"time"
)

const url string = "https://r-a-d.io/api"

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

func Now() (*Status, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	dec := json.NewDecoder(res.Body)
	defer res.Body.Close()

	var r = new(Result)
	if err := dec.Decode(r); err != nil {
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
		return &r.Status, nil
	}
	pos.Current = pos.Now.Sub(pos.Start)
	pos.CurrentPerc = (100 / float64(pos.Len)) * float64(pos.Current)
	r.Status.Pos = pos

	return &r.Status, nil
}
