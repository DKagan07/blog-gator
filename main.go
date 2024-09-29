package main

import (
	"fmt"
	"log"
	"os"

	"blog-gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	// Create/read the current config
	cfg, err := config.ReadConfig()
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	// create a new instance of a state struct
	s := &state{
		cfg: &cfg,
	}

	// create a new instance of a commands struct
	commands := commands{
		cmds: make(map[string]func(*state, command) error),
	}

	// Registering the "login" command, which should be basic
	commands.register("login", handlerLogin)

	// Now parsing the args
	if len(os.Args) < 2 {
		fmt.Println("need more args to use cli")
		os.Exit(1)
	}

	fmt.Printf("args: %+v", os.Args)
	// Register the login command
	cmmd := os.Args[1]
	args := []string{}
	if len(os.Args) > 2 {
		args = os.Args[2:]
	}

	fmt.Println("args blah: ", args)

	userCommand := command{
		Name:     cmmd,
		Commands: args,
	}

	if err = commands.run(s, userCommand); err != nil {
		log.Fatalf("failed to run command %s: %v", cmmd, err)
	}
}
