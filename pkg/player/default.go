package player

var defaultsAudioPlayer = map[string]interface{}{
	"no-config":  true,
	"no-video":   true,
	"no-sub":     true,
	"no-softvol": true,
}

func (p *Player) defaultProperties() error {
	for prop, val := range defaultsAudioPlayer {
		err := p.m.SetProperty(prop, val)
		if err != nil {
			return err
		}
	}
	return nil
}
