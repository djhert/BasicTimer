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
	go t.tick()
}


func (t *Timer) tick() {
	for {
		if !t.isRunning {
			return
		}
		time.Sleep(time.Second)
		if !t.isPaused {			 
			t.addSecond()
		}
	}
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
	time.Sleep(time.Second)
}

func (t *Timer) Get() (int, int, int) {
	return t.second,t.minute,t.hour
}

func (t *Timer) Print() {
	endString := ""
	if t.hour > 0 {
		endString = fmt.Sprintf("%s %d hour(s)", endString, t.hour)
	}
	if t.minute > 0 {
		endString = fmt.Sprintf("%s %d minute(s)", endString, t.minute)
	}
	endString = fmt.Sprintf("%s %d second(s)", endString, t.second)
	fmt.Printf("%s\n",endString)
}
