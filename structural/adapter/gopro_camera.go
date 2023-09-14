package adapter

import "fmt"

// HevcVideo This object act as a new service.
// It will be played by WindowsMoviePlayer, but our
// current player can't play this video because
// it has unsupported format.
// See to adapter.go file to see more details how
// we deal with current situation.
type HevcVideo struct{}

func (n *HevcVideo) PlayHevcVideo() {
	fmt.Println("Playing .HEVC video . . .")
}
