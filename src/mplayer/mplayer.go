package main

//导入规范，最好分层，前面的为核心类库，中间为第三方类库，最后为自己的类库，中间空一行
import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"

	//疑问？：自己的库，如何在pkg里引用   pkg/mplayer/mp
	"mplayer/library"
	"mplayer/mp"
)

var lib *library.MusicManager
var id int =1
//作用？
var ctrl, singal chan int


func main(){

	fmt.Println(`
	Enter following commands to control the player:
	lib list -- View the existing music lib
	lib add <name><artist><source><type> -- Add a music to the music lib
	lib remove <name> -- Remove the specified music from the lib
	play <name> -- Play the specified music
	q or e -- quit
	`)

	lib = library.NewMusicManager()

	//接收标准输入
	r := bufio.NewReader(os.Stdin)

	for{
		fmt.Print("Enter command ->")

		rawLine,_,_ := r.ReadLine()

		line := string(rawLine)

		if line == "q" || line == "e"{
			break
		}

		tokens := strings.Split(line," ")

		if tokens[0] == "lib" {
			handleLibCommands(tokens)
		}else if tokens[0] == "play" {
			handlePlayCommand(tokens)
		}else{
			fmt.Println("Unrecognized command:",tokens[0])
		}
	}
}

//处理播放命令
func handlePlayCommand(tokens []string) {

	if len(tokens) != 2{
		fmt.Println("Usage: plary <name>")
		return
	}

	e := lib.Find(tokens[1])

	if e == nil{
		fmt.Println("The music",tokens[1],"does not exist.")
		return
	}

	mp.Play(e.Source,e.Type)
}

//处理管理命令
func handleLibCommands(tokens []string) {

	switch tokens[1] {
	case "list":
		for i := 0;i < lib.Len() ; i++  {
			e,_ := lib.Get(i)
			fmt.Println(i+1,":",e.Name,e.Artist,e.Source,e.Type)

		}
	case "add":
		if len(tokens) == 6 {
			id ++
			lib.Add(&library.MusicEntry{
				strconv.Itoa(id),
				tokens[2],
				tokens[3],
				tokens[4],
				tokens[5]})
		}else {
			fmt.Println("Usage: lib add <name><artist><source><type>")
		}
	case "remove":
		if len(tokens) == 3{
			lib.RemoveByName(tokens[2])
		}else {
			fmt.Println("USAGE: lib remove <name>")
		}
	default:
		fmt.Println("Unrecognized lib command :",tokens[1])
	}
}
