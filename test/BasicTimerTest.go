package main
import BasicTimer "github.com/markedhero/BasicTimer"
import "time"

func main() {
	t := new(BasicTimer.Timer)
	t.Start()
	time.Sleep(time.Second * 60)
	t.Print()
	t.Pause(true)
	time.Sleep(time.Second * 10)
	t.Print()
	t.Pause(false)
	time.Sleep(time.Second * 10)
	t.Print()
	t.Stop()	
}
