package hackernews

import "fmt"

// Stories, comments, jobs, Ask HNs and even polls are just items.
type Item struct {
	ID          int    `json:"id"`          // The item's unique id.
	Deleted     bool   `json:"deleted"`     // true if the item is deleted.
	Type        string `json:"type"`        // The type of item. One of "job", "story", "comment", "poll", or "pollopt".
	By          string `json:"by"`          // The username of the item's author.
	Time        int    `json:"time"`        // Creation date of the item, in Unix Time.
	Text        string `json:"text"`        // The comment, story or poll text. HTML.
	Dead        bool   `json:"dead"`        // true if the item is dead.
	Parent      int    `json:"parent"`      // The comment's parent: either another comment or the relevant story.
	Poll        int    `json:"poll"`        // The pollopt's associated poll.
	Kids        []int  `json:"kids"`        // The ids of the item's comments, in ranked display order.
	Url         string `json:"url"`         // The URL of the story.
	Score       int    `json:"score"`       // The story's score, or the votes for a pollopt.
	Title       string `json:"title"`       // The title of the story, poll or job.
	Parts       []int  `json:"parts"`       // A list of related pollopts, in display order.
	Descendants int    `json:"descendants"` // In the case of stories or polls, the total comment count.
}

func (i Item) String() string {
	if i.Url != "" {
		return fmt.Sprintf("#%d %s %s", i.ID, i.Title, i.Url)
	}
	return fmt.Sprintf("#%d %s", i.ID, i.Title)
}
