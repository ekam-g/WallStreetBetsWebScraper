package Reddit

type read struct{}

//
//func (read) init() (err error, client *reddit.Client) {
//	posts, _, err := client.Subreddit.TopPosts(context.Background(), "golang", &reddit.ListPostOptions{
//		ListOptions: reddit.ListOptions{
//			Limit: 5,
//		},
//		Time: "all",
//	})
//	if err != nil {
//		return err
//	}
//	fmt.Printf("Received %d posts.\n", len(posts))
//}
