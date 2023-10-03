package hw1

import (
	"fmt"
	"os"

	"time"

	ntp "github.com/beevik/ntp"
)

func CurrentTime() string {
	ntptime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	}
	//ntptime := time.Now().Add(response.ClockOffset)
	return fmt.Sprintf("Current time: %s\nExact time: %s", time.Now().Round(time.Second), ntptime.Round(time.Second))
}
