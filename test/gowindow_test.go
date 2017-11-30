package gowindow_test

import (
	"errors"
	"fmt"
	"regexp"
	"testing"

	"github.com/dsolymosi/gowindow"
)

func TestAllWindows(t *testing.T) {
	r, _ := regexp.Compile(".*")
	fmt.Println("---listing of all windows below---")
	for _, s := range gowindow.FindWindow(r) {
		fmt.Println(s)
	}
	fmt.Println("---listing of all windows above---")
}

func TestNoWindows(t *testing.T) {
	r, _ := regexp.Compile("&^")
	if len(gowindow.FindWindow(r)) > 0 {
		t.Fail()
	}
}

func TestInputWindow(t *testing.T) {
	var inputRegexpString string
	var r *regexp.Regexp
	err := errors.New("Undefined")

	for err != nil {
		fmt.Print("Enter a regex corresponding to some windows: ")
		fmt.Scanln(&inputRegexpString)
		r, err = regexp.Compile(inputRegexpString)
	}

	fmt.Println("---listing of matching windows below---")
	for _, s := range gowindow.FindWindow(r) {
		fmt.Println(s)
	}
	fmt.Println("---listing of matching windows above---")

	yesNo := ""

	for yesNo == "" {
		fmt.Print("Was this the expected output (y/n)?: ")
		fmt.Scanln(&yesNo)
	}

	if yesNo == "n" || yesNo == "no" || yesNo == "N" || yesNo == "No" || yesNo == "NO" {
		t.Fail()
	}

}

func TestMain(m *testing.M) {

	m.Run()

	//wait for user input to close window
	fmt.Print("Press enter to continue...")
	fmt.Scanln()
}
