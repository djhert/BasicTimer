# BasicTimer
This is a basic timer that runs on another thread.  Used to see how long operations take in your Go Applications.
It's so basic, I did not include error checking.  

## Installation:
```sh
$ go get github.com/markedhero/BasicTimer/
```

## Usage: 

Include the Header:
```go
import BasicTimer "github.com/markedhero/BasicTimer"
```
Create the BasicTimer Object:

```go
timer := new(BasicTimer.Timer)
```
Then you can use the new BasicTimer Object in your code to control the timer.  See an example of the application in BasicTimerTest.go
```
Start() -- Sets the Timer object to "zero" and starts the timer thread.  This must be run first.
You can run this to recycle your Timers as well.  

Pause(bool) -- Pauses the timer, set to True to pause, false to unpause

Stop() -- Stops the timer

Get() (hour int, minute int, second int) -- Returns hour, minute, and second of the timer as ints

Print() -- Prints out the Hours, Minutes, and Seconds of the timer to the console
```
