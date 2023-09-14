package adapter

import "fmt"

// WindowsMoviePlayer Act as an abstraction that have
// a behaviour to play a video with exclusive windows
// video format. This abstraction means the legacy
// system in this case.
type WindowsMoviePlayer interface {
	Play()
}

type Player struct {
	Name string
}

func (p *Player) PlayVideo(windows WindowsMoviePlayer) {
	fmt.Println(p.Name, "try to plays some video from a windows computer")
	fmt.Println()
	fmt.Println(">>> Injecting supported format (.MP4) video")
	windows.Play()
}
