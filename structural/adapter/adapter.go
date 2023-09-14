package adapter

import "fmt"

// IphoneVideoAdapter This object has a purpose
// to act as an adapter from H264Video video format to Mp4Video format,
// so it can be played with our WindowsMoviePlayer.
type IphoneVideoAdapter struct {
	format *H264Video
}

func (a *IphoneVideoAdapter) Play() {
	fmt.Println("Converting H.264 video format to Mp4 format . . .")
	a.format.PlayH264Video()
	fmt.Println()
}

type CanonCameraVideoAdapter struct {
	format *MovVideo
}

func (c *CanonCameraVideoAdapter) Play() {
	fmt.Println("Converting .MOV video format to Mp4 format . . .")
	c.format.PlayMovVideo()
	fmt.Println()
}

type GoproCameraVideoAdapter struct {
	format *HevcVideo
}

func (g *GoproCameraVideoAdapter) Play() {
	fmt.Println("Converting .HEVC video format to Mp4 format . . .")
	g.format.PlayHevcVideo()
	fmt.Println()
}
