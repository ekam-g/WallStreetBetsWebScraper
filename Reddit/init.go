package Reddit

import (
	"fmt"
	"github.com/vartanbeno/go-reddit/v2/reddit"
	"golang.org/x/net/context"
)

type Scrape struct {
	Read read
}

func (Scrape) Init() (err error, client *reddit.Client) {
	credentials := reddit.Credentials{ID: "id", Secret: "MtR7jkXVNykBxeivt0181FOk40k3jA", Username: "username", Password: "password"}
	client, err = reddit.NewClient(credentials)
	if err != nil {
		panic(err)
	}
	posts, _, err := client.Subreddit.TopPosts(context.Background(), "golang", &reddit.ListPostOptions{
		ListOptions: reddit.ListOptions{
			Limit: 5,
		},
		Time: "all",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Received %d posts.\n", len(posts))
	return err, client
}
