package mp

import (
	"fmt"
)

//定义播放接口
type Player interface {
	Play(source string)
}

func Play(source,mtype string) {

	var p Player

	switch mtype {
	case "MP3":
		p = &MP3Player{}
	case "WAV":
		p = &WAVPlayer{}
	default:
		fmt.Println("Unsupported music type ",mtype)
		return
	}
	p.Play(source)
}