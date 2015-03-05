package ugolibev3

import (
	"io/ioutil"
	"path/filepath"
)

//Reads value of n-th plugged sensor(not counting unpluged).
func ReadSensorValue(n int8) string {
	sensors := getSensorsList()
	dat, err := ioutil.ReadFile(sensors[n] + "/value0")
	errorcheck(err)
	return string(dat)
}

func getSensorsList() []string {
	sensorslist, err := filepath.Glob("/sys/class/lego-sensor/*")
	errorcheck(err)
	return sensorslist
}
func errorcheck(e error) {
	if e != nil {
		panic(e)
	}
}
