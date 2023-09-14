package adapter

import "fmt"

func Execute() {
	player := &Player{Name: "Barsky"}
	mp4 := &Mp4Video{}

	player.PlayVideo(mp4)

	h264 := &H264Video{}
	iphoneVideoAdapter := &IphoneVideoAdapter{format: h264}
	fmt.Println(">>> Injecting unsupported format (.H.264) video")
	iphoneVideoAdapter.Play()

	mov := &MovVideo{}
	canonVideoAdapter := &CanonCameraVideoAdapter{format: mov}
	fmt.Println(">>> Injecting unsupported format (.MOV) video")
	canonVideoAdapter.Play()

	hevc := &HevcVideo{}
	goproVideoAdapter := &GoproCameraVideoAdapter{format: hevc}
	fmt.Println(">>> Injecting unsupported format (.HEVC) video")
	goproVideoAdapter.Play()
}
