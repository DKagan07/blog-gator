package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"

	"blog-gator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Commands) == 0 {
		return fmt.Errorf("expect: %s <name>", cmd.Name)
	}

	newName := cmd.Commands[0]
	_, err := s.db.Getuser(context.Background(), newName)
	if err != nil {
		return fmt.Errorf("name: %s doesn't exist in database", newName)
	}

	err = s.cfg.SetUser(newName)
	if err != nil {
		return fmt.Errorf("could not set user, error: %+v", err)
	}

	fmt.Printf("User %s has logged in", newName)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Commands) < 1 && len(cmd.Commands) > 1 {
		return fmt.Errorf("Usage should be: %s <name>", cmd.Name)
	}

	if cmd.Commands[0] == "unknown" {
		return errors.New("cannot have name 'unknown'")
	}

	uuid, err := uuid.NewV6()
	if err != nil {
		return fmt.Errorf("with creating a UUID: %v", err)
	}

	now := sql.NullTime{
		Time:  time.Now().UTC(),
		Valid: true,
	}

	userParams := database.CreateUserParams{
		ID:        uuid,
		CreatedAt: now,
		UpdatedAt: now,
		Name:      cmd.Commands[0],
	}

	user, err := s.db.CreateUser(context.Background(), userParams)
	if err != nil {
		return fmt.Errorf("creating user: %v", err)
	}

	if err = s.cfg.SetUser(user.Name); err != nil {
		return fmt.Errorf("setting user in handleRegister: %+v", err)
	}

	fmt.Printf("User %s has been registered\n", user.Name)

	return nil
}
