package mp

import "fmt"

type WAVPlayer struct {

}

func (p *WAVPlayer) Play(source string){

	fmt.Println("Playing WAV music ",source)
}
