package ffmpeg

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// Variant is HLS variant that gonna be use to generate HLS master playlist
type Variant struct {
	// URL indicate the location of the variant playlist.
	// If variant located on remote server, this url should
	// contain the full url
	URL string

	// Bandwidth is an integer that is the upper bound of
	// the overall bitrate for each media file, in bits per second
	Bandwidth string

	// Resolution is display size, in pixels, at which to display
	// all of the video in the playlist
	Resolution string

	// Codecs is quoted string containing a comma-separated list of formats,
	// where each format specifies a media sample type that's present
	// in a media segment in the playlist file. Valid format identifiers are
	// those in the ISO file format name space defined by RFC 6381
	Codecs string
}

// GenerateHLSVariant will generate variants info from the given resolutions.
// The available resolutions are: 360p, 480p, 720p and 1080p.
func generateHLSVariant(resOptions []string, locPrefix string) ([]*Variant, error) {
	if len(resOptions) == 0 {
		return nil, errors.New("Please give at least 1 resolutions.")
	}

	var variants []*Variant

	for _, r := range resOptions {
		c, err := getConfig(r)
		if err != nil {
			continue
		}

		url := fmt.Sprintf("%s.m3u8", c.PlaylistName)
		if locPrefix != "" {
			url = locPrefix + "/" + url
		}

		v := &Variant{
			URL:        url,
			Bandwidth:  c.Bandwidth,
			Resolution: c.Resolution,
		}

		variants = append(variants, v)
	}

	if len(variants) == 0 {
		return nil, errors.New("No valid resolutions found.")
	}

	return variants, nil
}

// GeneratePlaylist will generate playlist file from the given variants.
// Variant itself can be generate from GenerateHLSVariant() function of
// suplied by the caller
func generatePlaylist(variants []*Variant, targetPath, filename string) {
	// Set default filename
	if filename == "" {
		filename = "playlist.m3u8"
	}

	// M3U Header
	data := "#EXTM3U\n"
	data += "#EXT-X-VERSION:3\n"

	// Add M3U Info for each variant
	for _, v := range variants {
		// URL & bandwidth is required,
		// if not found we will excluded them from the playlist
		if v.URL == "" || v.Bandwidth == "" {
			continue
		}

		data += "#EXT-X-STREAM-INF:"
		data += fmt.Sprintf("BANDWIDTH=%s", v.Bandwidth)
		if v.Resolution != "" {
			data += fmt.Sprintf(",RESOLUTION=%s", v.Resolution)
		}
		if v.Codecs != "" {
			data += fmt.Sprintf(",CODECS=%s", v.Codecs)
		}

		data += fmt.Sprintf("\n%s\n", v.URL)
	}

	// Write everything to the file
	f, _ := os.Create(filepath.Join(targetPath, filename))
	defer f.Close()

	f.Write([]byte(data))
}
