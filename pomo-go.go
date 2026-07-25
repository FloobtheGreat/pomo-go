package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
	"bufio"

	"charm.land/huh/v2"
	"charm.land/huh/v2/spinner"
	//"charm.land/lipgloss/v2"
	//xstrings "github.com/charmbracelet/x/exp/strings"
)

type Timer int

const (
	Standard = iota + 1
	Short
	Long
)


func (t Timer) String() string {
	switch t {
	case Standard:
		return "Standard"
	case Short:
		return "Short"
	case Long:
		return "Long"
	default:
		return ""
	}
}

func main() {
	var (
		mode      int
		cycles    string
		confirm   bool
		cyclesInt int
		work int
		rest int
	)

	form := huh.NewForm(
		huh.NewGroup(huh.NewNote().
			Title("Pomo-Go").
			Description("Welcome to Pomo-Go.\nReady to focus?").
			Next(true).
			NextLabel("Next"),
		),

		huh.NewGroup(
			huh.NewSelect[int]().
				Title("Timer Type").
				Options(
					huh.NewOption("Standard", Standard).Selected(true),
					huh.NewOption("Short", Short),
					huh.NewOption("Long", Long),
				).
				Value(&mode),

			huh.NewInput().
				Title("How many Cycles?").
				Prompt("1-4?").
				Validate(func(s string) error {
					num, err := strconv.Atoi(s)
					if err != nil || num < 1 || num > 4 {
						return errors.New("Not valid input")
					}
					return nil
				}).
				Value(&cycles),
		),

		huh.NewGroup(
			huh.NewConfirm().
				Title("Ready to go?").
				Value(&confirm).
				Affirmative("Yes!").
				Negative("No."),
		),
	)

	err := form.Run()

	if err != nil {
		fmt.Println("Uh oh: ", err)
		os.Exit(1)
	}

	cyclesInt, _ = strconv.Atoi(cycles)

	prepTimer := func()  {
		// this really isn't necessary
		// I just wanted to use it
		work, rest = getMinutes(mode)
		time.Sleep(2 * time.Second)
	}

	_ = spinner.New().Title("Getting ready...").Action(prepTimer).Run()

	for i := 0; i < cyclesInt; i++ {
		//fmt.Printf("\033[F")
		fmt.Printf("Get to Work!...\n")
		runTimer(work)
		fmt.Printf("Time for a break!\n")
		runTimer(rest)
		fmt.Print("Hit ENTER to continue.")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}


}

func getMinutes(m int) (int, int) {
	switch m {
	case 1:
		return 25, 5
	case 2:
		return 1, 1
	case 3:
		return 50, 10
	default:
		return 0, 0
	}
}

func formatTime(totalSeconds int) string {
	minutes := totalSeconds / 60
	seconds := totalSeconds % 60

	return fmt.Sprintf("%02dm%02ds", minutes, seconds)
}

func runTimer(minutes int) {
	remaining := minutes * 30

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for remaining > 0 {
		<-ticker.C
		remaining--
		fmt.Println(formatTime(remaining))
		fmt.Printf("\033[F")
	}

}
