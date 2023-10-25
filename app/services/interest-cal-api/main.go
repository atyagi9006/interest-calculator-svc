package main

import (
	"github.com/atyagi9006/interest-calculator-svc/app/services/interest-cal-api/start"
)

var build = "dev"

func main() {
	start.Run(build)
}
