package main

import (
	"github.com/atyagi9006/interest-calculator-svc/app/services/interest-cal-api/start"
)

//these are the build flags 
var (
	build = "dev"
	serviceName="interest-calculator-svc"
)

func main() {
	start.Run(serviceName+":"+build)
}
