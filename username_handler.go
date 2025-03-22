package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/OferRavid/bloggregator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return errors.New("login command requires a single argument: username")
	}
	username := cmd.Args[0]
	_, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("can't login as %s. user isn't in database", username)
	}
	err = s.cfg.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Printf("Username was successfully set to %s\n", username)

	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return errors.New("login command requires a single argument: username")
	}

	username := cmd.Args[0]
	_, err := s.db.GetUser(context.Background(), username)
	if err == nil {
		log.Fatalf("User %s already exists", username)
	}

	id := uuid.New()
	timeNow := time.Now()
	user, err := s.db.CreateUser(
		context.Background(),
		database.CreateUserParams{
			ID:        id,
			CreatedAt: timeNow,
			UpdatedAt: timeNow,
			Name:      username,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to create user %s: %v", username, err)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Printf("User %s was created successfully\n", user.Name)
	fmt.Println(user)

	return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteData(context.Background())
	if err != nil {
		return fmt.Errorf("failed to delete users: %v", err)
	}
	fmt.Println("All users were deleted successfully.")
	return nil
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	for _, user := range users {
		msg := fmt.Sprintf("* %s", user.Name)
		if user.Name == s.cfg.CurrentUserName {
			msg += " (current)"
		}
		fmt.Println(msg)
	}
	return nil
}
