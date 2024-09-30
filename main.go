package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"blog-gator/internal/config"
	"blog-gator/internal/database"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

func main() {
	// Create/read the current config
	cfg, err := config.ReadConfig()
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	// Connecting to Postgres
	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		fmt.Println("cannot connect to Postgres")
		os.Exit(1)
	}

	dbQueries := database.New(db)

	// create a new instance of a state struct
	s := &state{
		cfg: &cfg,
		db:  dbQueries,
	}

	// create a new instance of a commands struct
	commands := commands{
		cmds: make(map[string]func(*state, command) error),
	}

	// Registering commands, which should be basic
	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)

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

	userCommand := command{
		Name:     cmmd,
		Commands: args,
	}

	if err = commands.run(s, userCommand); err != nil {
		log.Printf("failed to run command %s: %v", cmmd, err)
		os.Exit(1)
	}
}
