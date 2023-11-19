package twitterscraper

type newTimelineV2 struct {
	Data struct {
		User struct {
			Result struct {
				TimelineV2 struct {
					Timeline struct {
						Instructions []struct {
							Entries []newEntry `json:"entries"`
							Entry   newEntry   `json:"entry"`
							Type    string     `json:"type"`
						} `json:"instructions"`
					} `json:"timeline"`
				} `json:"timeline_v2"`
			} `json:"result"`
		} `json:"user"`
	} `json:"data"`
}

type newEntry struct {
	Content struct {
		ItemContent struct {
			TweetDisplayType string `json:"tweetDisplayType"`
			TweetResults     struct {
				Result newResult `json:"result"`
			} `json:"tweet_results"`
		} `json:"itemContent"`
	} `json:"content"`
}

type newResult struct {
	Typename string `json:"__typename"`

	QuotedStatusResult struct {
		Result *result `json:"result"`
	} `json:"quoted_status_result"`
	Legacy newLegacyTweet `json:"legacy"`
}

type newLegacyTweet struct {
	ConversationIDStr string `json:"conversation_id_str"`
	CreatedAt         string `json:"created_at"`
	FullText          string `json:"full_text"`
	IDStr             string `json:"id_str"`
}

func (n *newTimelineV2) GetIDS() []string {
	var ids []string

	for _, instruction := range n.Data.User.Result.TimelineV2.Timeline.Instructions {
		for _, entry := range instruction.Entries {
			ids = append(ids, entry.Content.ItemContent.TweetResults.Result.Legacy.ConversationIDStr)
		}
	}
	return ids
}
