package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"os"
	"syscall"

	"golang.org/x/crypto/bcrypt"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/ProgrammingLab/prolab-accounts/dao"
	"github.com/pkg/errors"

	"github.com/ProgrammingLab/prolab-accounts/app/config"
	"github.com/ProgrammingLab/prolab-accounts/app/di"
)

func main() {
	os.Exit(run())
}

func run() int {
	c, err := config.LoadConfig()
	if err != nil {
		fmt.Errorf("%+v", err)
		return 1
	}
	store, err := di.NewStoreComponent(c)
	if err != nil {
		fmt.Errorf("%+v", err)
		return 1
	}

	s := bufio.NewScanner(os.Stdin)
	u := &dao.User{
		Name:     readText(s, "user name >"),
		Email:    readText(s, "email >"),
		FullName: readText(s, "full name >"),
	}
	d, err := readPassword(s)
	if err != nil {
		fmt.Errorf("%+v", err)
		return 1
	}
	u.PasswordDigest = string(d)

	err = store.UserStore(context.Background()).CreateUser(u)
	if err != nil {
		fmt.Errorf("%+v", err)
		return 1
	}

	fmt.Printf("%+v\n", *u)

	return 0
}

func readText(s *bufio.Scanner, message string) string {
	fmt.Print(message)
	s.Scan()
	return s.Text()
}

func readPassword(s *bufio.Scanner) ([]byte, error) {
	for {
		fmt.Print("password >")
		p, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			return nil, errors.WithStack(err)
		}
		fmt.Print("\nrepeate password >")
		rp, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			return nil, errors.WithStack(err)
		}
		fmt.Println()

		if bytes.Equal(p, rp) {
			d, err := bcrypt.GenerateFromPassword(p, bcrypt.DefaultCost)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			return d, nil
		}

		fmt.Println("The passwords do not match. Try again.")
	}
}
