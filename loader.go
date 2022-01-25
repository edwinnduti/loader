package loader

import (
	"fmt"
	"io"
	"sync"
	"time"
)

var (
	mu *sync.RWMutex = &sync.RWMutex{}
	done = make(chan struct{})
	spins = map[string][]string{
		"rs" : {"\\", "|", "/", "-"},
		"dc" : {"⠈⠁", "⠈⠑", "⠈⠱", "⠈⡱", "⢀⡱", "⢄⡱", "⢄⡱", "⢆⡱", "⢎⡱", "⢎⡰", "⢎⡠", "⢎⡀", "⢎⠁", "⠎⠁", "⠊⠁"},
		"rb" : {"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷" },
	}
)

type Loader struct {
	Name 			string
	Delay   		time.Duration
	Output 			io.Writer
	StartMessage	string
	EndMessage		string
	HideCursor		string
	ShowCursor		string
	FinalMessage 	string
	StopChan		chan struct{}
}

func New(w io.Writer, name string, d time.Duration, firstmsg , endmsg string) *Loader {
	return &Loader{
		Output: w,
		Name: name,
		Delay: d,
		StartMessage: firstmsg,
		EndMessage: endmsg,
		ShowCursor: "\033[?25h",
		HideCursor: "\033[?25l",
		StopChan: make(chan struct{}, 1),
	}
}

func (L *Loader) Initialize(){

	// hide cursor
	fmt.Fprint(L.Output, L.HideCursor)

	go func(){

		for{

			for _, spinItem := range spins[L.Name] {
				select {
				case <- L.StopChan:
					return
				default:
					outLine := fmt.Sprintf("\r%s%s%s",L.StartMessage, spinItem, L.EndMessage)
					fmt.Fprintf(L.Output, outLine)
					time.Sleep(L.Delay)
				}
			}
		}
	}()
}

func (L *Loader) End(){

	// bring back cursor
	fmt.Fprintf(L.Output, L.ShowCursor)

	// erase line
	fmt.Fprintf(L.Output, "\r\033[K")

	L.StopChan <- struct{}{}
	return
}