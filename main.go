package main

import (
	"fmt"
	"time"
	"flag"
	"strings"
	"os"
	"os/exec"
	"runtime"
	"github.com/fatih/color"
)

var at = flag.String("a", "whale", "Sets displayed ascii art")

func main() {

	help := flag.Bool("h", false, "List of available commands")
	lists := flag.Bool("l", false, "List of colors/art available")
	credits := flag.Bool("c", false, "Finfetch's credits")
	realtime := flag.Bool("r", false, "Enables real-time changes")
	tc := flag.String("tc", "cyan", "Sets text color")

	flag.Parse()

	switch {
	//-h
	case *help:
		fmt.Println(helper)
		return
	//-c
	case *credits:
		fmt.Println(credit)
		return
	//-l
	case *lists:
		fmt.Println(list)
		return
	}

	//-tc
	txtcolor := strings.ToLower(*tc)
	switch txtcolor {
	case "red":
		color.Set(color.FgRed)
	case "green":
		color.Set(color.FgGreen)
	case "yellow":
		color.Set(color.FgYellow)
	case "blue":
		color.Set(color.FgBlue)
	case "magenta":
		color.Set(color.FgMagenta)
	case "cyan":
		color.Set(color.FgCyan)
	case "white":
		color.Set(color.FgWhite)
	case "black":
		color.Set(color.FgBlack)
	default:
		color.Set(color.FgRed)
		fmt.Println(fmt.Sprintf(`"%+v" is not an available color!`, *tc))
		color.Unset()
		return
		}

	//-r
	switch {
	case *realtime:
		for {
					cleared()
					ascii()
					time.Sleep(1 * time.Second)
				}
	default:
		ascii()
		color.Unset()
		return
	}
}

func ascii(){
	clock := time.Now()
	date := fmt.Sprintf("%+v %+v, %+v", clock.Month(), clock.Day(), clock.Year())
	time := fmt.Sprintf("%+v:%+v:%+v", clock.Hour(), clock.Minute(), clock.Second())

	asciiart := strings.ToLower(*at)
	switch asciiart {
	case "whale":
		fmt.Println(fmt.Sprintf(`     .-
'--./ /     _.---.
'-,  (__..--       \ 	Date: %+v
   \          .     |	Time: %+v
    \,.__.   ,__.--/
      '._/_.'_____/

`, date, time))

	case "cat":
		fmt.Println(fmt.Sprintf(`      /\_/\
 ____/ o o \	Date: %+v
/~____  =ø= /	Time: %+v
(______)__m_m)

 `, date, time))
 
	case "clock":
		fmt.Println(fmt.Sprintf(`  _______
 /  12   \
|    |    |	Date: %+v
|9   |   3|	Time: %+v
|     \   |
|         |
 \___6___/
		`, date, time))
		
	case "none":
		fmt.Println(fmt.Sprintf(`Date: %+v
Time: %+v
		`, date, time))
		
 	default:
 		color.Set(color.FgRed)
 		fmt.Println(fmt.Sprintf(`"%+v" is not an available art!`, asciiart))
 		color.Unset()
	}
	return
}

func cleared(){
	switch os := runtime.GOOS; os {
	case "windows":
		wincls()
	default:
		clr()
	}
}

func wincls() {
	cls := exec.Command("cmd", "/c", "cls")
	cls.Stdout = os.Stdout
	cls.Run()
}

func clr(){
	clear := exec.Command("clear")
	clear.Stdout = os.Stdout
	clear.Run()
}

var helper = `Finfetch 2.1
Usage: finfetch [prefix]

Options:
	-h			List of available commands
	-l			List of colors/art available
	-c			Finfetch's credits
	-tc [Color]		Sets text color
	-a [Name]		Sets displayed ascii art
	-r			Enables real-time changes
	
	`

var list = `Colors:
	-red
	-green
	-yellow
	-blue
	-magenta
	-cyan
	-white

Ascii:
	-whale
	-cat
	-clock
	-none

`

var credit = `Finfetch Credits
Created by arkizenty

Resources Used:
	asciiart.website
	github.com/fatih/color

Special Thanks:
	screenfetch
	neofetch
	ufetch

  ______ _        __     _       _
 |  ____(_)      / _|   | |     | |
 | |__   _ _ __ | |_ ___| |_ ___| |__
 |  __| | | '_ \|  _/ _ \ __/ __| '_ \
 | |    | | | | | ||  __/ || (__| | | |
 |_|    |_|_| |_|_| \___|\__\___|_| |_|

`
