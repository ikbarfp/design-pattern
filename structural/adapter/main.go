package adapter

import "fmt"

func Execute() {
	// Initialize new player to play the only supported video format
	// which is Mp4Video
	player := &Player{Name: "Barsky"}
	mp4 := &Mp4Video{}

	// Play the video from Player
	player.PlayVideo(mp4)

	// Got new iPhone and we want to play the video
	// on our computers, but we need an adapter to
	// play this H.264 video format to .MP4
	h264 := &H264Video{}
	iphoneVideoAdapter := &IphoneVideoAdapter{format: h264}
	fmt.Println(">>> Injecting unsupported format (.H.264) video")
	iphoneVideoAdapter.Play()

	// Got new DSLR Camera and we want to play the video
	// on our computers, but we need an adapter to
	// play this .MOV video format to .MP4
	mov := &MovVideo{}
	canonVideoAdapter := &CanonCameraVideoAdapter{format: mov}
	fmt.Println(">>> Injecting unsupported format (.MOV) video")
	canonVideoAdapter.Play()

	// Got new GoPro Camera and we want to play the video
	// on our computers, but we need an adapter to
	// play this .HEVC video format to .MP4
	hevc := &HevcVideo{}
	goproVideoAdapter := &GoproCameraVideoAdapter{format: hevc}
	fmt.Println(">>> Injecting unsupported format (.HEVC) video")
	goproVideoAdapter.Play()
}
