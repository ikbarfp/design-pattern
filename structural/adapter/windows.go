package adapter

import "fmt"

// Mp4Video This object inherit behaviour from the WindowsMoviePlayer
// This object is the only object can play the movie from the legacy system.
type Mp4Video struct{}

func (o *Mp4Video) Play() {
	fmt.Println("Playing .MP4 video on OS Windows . . .")
	fmt.Println()
}
