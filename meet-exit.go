package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func timeVaridate(target string) (int, error) {
	const min int = 0

	targetVal, err := strconv.Atoi(target)
	if err != nil {
		return targetVal, errors.New("Value error.")
	}

	if targetVal < min {
		return targetVal, errors.New("Number below the range.")
	}

	return targetVal, nil
}

func main() {
	flag.Parse()

	var leftTime time.Duration
	var leftTimeMin time.Duration

	nowTime := time.Now()

	if len(flag.Args()) == 4 {
		hour, err := timeVaridate(flag.Args()[1])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		min, err := timeVaridate(flag.Args()[2])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		sec, err := timeVaridate(flag.Args()[3])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		leftTimeMin = time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), hour, min, sec, 0, time.Local).Sub(nowTime)
	}

	if len(flag.Args()) == 1 {
		specifyTime, err := timeVaridate(flag.Args()[0])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		leftTime = time.Duration(specifyTime)
		leftTimeMin = time.Minute * leftTime
	}

	if len(flag.Args()) != 1 && len(flag.Args()) != 4 {
		fmt.Println("Argument Error!")
		os.Exit(1)
	}

	limitTime := nowTime.Add(leftTimeMin)
	timer := time.NewTimer(leftTimeMin)

	fmt.Printf("Meet ends at %s.\n", limitTime.Format("2006-01-02 15:04:05"))

	<-timer.C
	_, err := exec.Command("osascript", "closeTab.scpt").Output()
	if err != nil {
		fmt.Println("something wrong!")
		os.Exit(1)
	}
	fmt.Println("Meet has ended.")
}
