# ugolibev3
Really micro (u) library for reading raw sensor values from ev3dev.

###How to use
__Simple__:

go get github.com/janekjan/ugolibev3

go install github.com/janekjan/ugolibev3

Import with: import "github.com/janekjan/ugolibev3"

####Usage
value := ugolibev3.ReadSensorValue(0)

Put 0 as an argument when you have only one sensor attached and this will work. If you have two sensors the last attached sensor will return from ReadSensorValue(1) and the first attached - from 0. IT'S NOT THE PORT NUMBER.
