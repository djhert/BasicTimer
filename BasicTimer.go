package BasicTimer
import "fmt"
import "time"

type Timer struct {
	second int
	minute int
	hour int
	isPaused bool
	isRunning bool
}
func (t *Timer) Start() {	
	t.second = 1
	t.minute = 0
	t.hour = 0
	t.isPaused = false
	t.isRunning = true
	go func() {
		for t.isRunning {
			time.Sleep(time.Second)
			if(!t.isPaused) {			 
				t.addSecond()
			}
		}
    }()    
}

func (t *Timer) addSecond() {
	t.second++
	if t.second >= 60 {
		t.minute++
		t.second = 0
	}
	if t.minute >= 60 {
		t.hour ++
		t.minute = 0
	}
}

func (t *Timer) Pause(i bool) {
	t.isPaused = i
}

func (t *Timer) Stop() {
	t.isRunning = false
}

func (t *Timer) Get() (int, int, int) {
	return t.second,t.minute,t.hour
}

func (t *Timer) Print() {
	fmt.Printf("%d hour(s) %d minute(s) %d second(s)\n", t.hour,t.minute,t.second)
}
