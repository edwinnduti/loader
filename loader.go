/*
A loader library that shows spinner animation
*/
package loader

import (
	"fmt"
	"io"
	"time"
)

// spin types
var (
	spins = map[string][]string{
		"rs" : {"\\", "|", "/", "-"},
		"dc" : {"⠈⠁", "⠈⠑", "⠈⠱", "⠈⡱", "⢀⡱", "⢄⡱", "⢄⡱", "⢆⡱", "⢎⡱", "⢎⡰", "⢎⡠", "⢎⡀", "⢎⠁", "⠎⠁", "⠊⠁"},
		"rb" : {"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷" },
	}
)

// loader contents
type Loader struct {
	Name 			string
	Delay   		time.Duration
	Output 			io.Writer
	StartMessage	string
	EndMessage		string
	HideCursor		string
	ShowCursor		string
	StopChan		chan struct{}
}

// create new spinner
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

// run the spinner
func (L *Loader) Initialize(){
	// hide cursor
	fmt.Fprint(L.Output, L.HideCursor)

	go func(){

		for{
			// get the current spin type
			for _, spinItem := range spins[L.Name] {
				select {
					// stop the spinner
				case <- L.StopChan:
					return
				default:
					// print the spinner
					outLine := fmt.Sprintf("\r%s%s%s",L.StartMessage, spinItem, L.EndMessage)
					fmt.Fprint(L.Output, outLine)
					time.Sleep(L.Delay)
				}
			}
		}
	}()
}

// stop the spinner
func (L *Loader) End(finalMessage string){

	// bring back cursor
	fmt.Fprintf(L.Output, L.ShowCursor)

	// erase line
	fmt.Fprint(L.Output, "\r\033[K", finalMessage)

	// close the stop channel
	L.StopChan <- struct{}{}
}