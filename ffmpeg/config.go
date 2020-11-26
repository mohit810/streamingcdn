package ffmpeg

import "errors"

type config struct {
	Name         string
	VideoBitrate string
	Maxrate      string
	BufSize      string
	AudioBitrate string
	Resolution   string
	Bandwidth    string
	PlaylistName string
}

var preset = map[string]*config{
	"360p": {
		Name:         "360p",
		VideoBitrate: "800k",
		Maxrate:      "856k",
		BufSize:      "1200k",
		AudioBitrate: "96k",
		Resolution:   "640x360",
		Bandwidth:    "800000",
		PlaylistName: "quality_0",
	},
	"480p": {
		Name:         "480p",
		VideoBitrate: "1400k",
		Maxrate:      "1498k",
		BufSize:      "2100k",
		AudioBitrate: "128k",
		Resolution:   "842x480",
		Bandwidth:    "1400000",
		PlaylistName: "quality_1",
	},
	"720p": {
		Name:         "720p",
		VideoBitrate: "2800k",
		Maxrate:      "2996k",
		BufSize:      "4200k",
		AudioBitrate: "128k",
		Resolution:   "1280x720",
		Bandwidth:    "2800000",
		PlaylistName: "quality_2",
	},
	"1080p": {
		Name:         "1080p",
		VideoBitrate: "5000k",
		Maxrate:      "5350k",
		BufSize:      "7500k",
		AudioBitrate: "192k",
		Resolution:   "1920x1080",
		Bandwidth:    "5000000",
		PlaylistName: "quality_3",
	},
}

// getConfig return config from the available preset
func getConfig(res string) (*config, error) {
	cfg, ok := preset[res]
	if !ok {
		return nil, errors.New("Preset not found")
	}

	return cfg, nil
}
