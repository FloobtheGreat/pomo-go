package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

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
				Placeholder("1").
				Validate(func(s string) error {
					num, err := strconv.Atoi(s)
					if err != nil || num < 0 || num > 4 {
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

	prepTimer := func() {
		// this really isn't necessary
		// I just wanted to use it
		time.Sleep(2 * time.Second)
	}

	_ = spinner.New().Title("Getting ready...").Action(prepTimer).Run()

	fmt.Printf("%v\n%v\n%v\n", mode, cyclesInt, confirm)

	// c := time.Tick(5 * time.Second)
	// for next := range c {
	// 	fmt.Printf("%v %s\n", next, statusUpdate())
	// }

}
