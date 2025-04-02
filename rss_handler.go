package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	rssFeed, err := fetchFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("failed to fetch feed from url: %v", feedURL)
	}

	fmt.Println(rssFeed)
	return nil
}
