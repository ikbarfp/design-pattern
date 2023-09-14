package adapter

import "fmt"

// H264Video This object act as a new service.
// It will be played by WindowsMoviePlayer, but our
// current player can't play this video because
// it has unsupported format.
// See to adapter.go file to see more details how
// we deal with current situation.
type H264Video struct{}

func (n *H264Video) PlayH264Video() {
	fmt.Println("Playing H.264 video . . .")
}
