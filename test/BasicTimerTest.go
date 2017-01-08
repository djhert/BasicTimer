package main

//Import the Timer
import BasicTimer "github.com/markedhero/BasicTimer"
import "time" //Used for time.Sleep(), not needed otherwise

func main() {
	//Create the object
	basicTimer := new(BasicTimer.Timer)
	//Start the timer
	basicTimer.Start()
	//Wait for 1 minute
	time.Sleep(time.Second * 60)
	//Print the current time accumulated
	basicTimer.Print()
	//Pause the timer
	basicTimer.Pause(true)
	//Wait 10 seconds while paused
	time.Sleep(time.Second * 10)
	//Show no time has been accumulated
	basicTimer.Print()
	//Unpause timer
	basicTimer.Pause(false)
	//Wait another 10 seconds
	time.Sleep(time.Second * 10)
	//Show time is being accumulated after unpausing
	basicTimer.Print()
	//End the timer
	basicTimer.Stop()	
	/* Once "Stop()" has been run, you can reset the timer by using Start() */
}
