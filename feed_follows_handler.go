package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/OferRavid/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <feed_url>", cmd.Name)
	}

	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("failed to select feed by url: %w", err)
	}

	feed_follow, err := s.db.CreateFeedFollow(
		context.Background(),
		database.CreateFeedFollowParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			UserID:    user.ID,
			FeedID:    feed.ID,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to insert feed follow: %w", err)
	}

	fmt.Println(feed_follow.FeedName, feed_follow.UserName)
	return nil
}

func handlerListFeedFollows(s *state, cmd command, user database.User) error {
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("failed to get feeds for user - %s : %w", s.cfg.CurrentUserName, err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}

	for _, feed := range feeds {
		fmt.Println(feed.FeedName)
	}

	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <feed_url>", cmd.Name)
	}

	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("failed to select feed by url: %w", err)
	}

	err = s.db.RemoveFeedFollowForUser(
		context.Background(),
		database.RemoveFeedFollowForUserParams{
			UserID: user.ID,
			FeedID: feed.ID,
		},
	)
	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("user %s doesn't follow this feed: %s", user.Name, feed.Name)
	}

	if err != nil {
		return fmt.Errorf("failed to delete feed follow: %w", err)
	}

	fmt.Printf("User %s successfully unfollowed feed %s", user.Name, feed.Name)
	return nil
}
