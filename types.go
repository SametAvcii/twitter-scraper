package twitterscraper

import "time"

type (
	// Mention type.
	Mention struct {
		ID       string
		Username string
		Name     string
	}

	// Photo type.
	Photo struct {
		ID  string
		URL string
	}

	// Video type.
	Video struct {
		ID      string
		Preview string
		URL     string
	}

	// GIF type.
	GIF struct {
		ID      string
		Preview string
		URL     string
	}

	// Tweet type.
	Tweet struct {
		ConversationID    string
		GIFs              []GIF
		Hashtags          []string
		HTML              string
		ID                string
		InReplyToStatus   *Tweet
		InReplyToStatusID string
		IsQuoted          bool
		IsPin             bool
		IsReply           bool
		IsRetweet         bool
		IsSelfThread      bool
		Likes             int
		Name              string
		Mentions          []Mention
		PermanentURL      string
		Photos            []Photo
		Place             *Place
		QuotedStatus      *Tweet
		QuotedStatusID    string
		Replies           int
		Retweets          int
		RetweetedStatus   *Tweet
		RetweetedStatusID string
		Text              string
		Thread            []*Tweet
		TimeParsed        time.Time
		Timestamp         int64
		URLs              []string
		UserID            string
		Username          string
		Videos            []Video
		Views             int
		SensitiveContent  bool
	}

	// ProfileResult of scrapping.
	ProfileResult struct {
		Profile
		Error error
	}

	// TweetResult of scrapping.
	TweetResult struct {
		Tweet
		Error error
	}

	legacyTweet struct {
		ConversationIDStr string `json:"conversation_id_str"`
		CreatedAt         string `json:"created_at"`
		FavoriteCount     int    `json:"favorite_count"`
		FullText          string `json:"full_text"`
		Entities          struct {
			Hashtags []struct {
				Text string `json:"text"`
			} `json:"hashtags"`
			Media []struct {
				MediaURLHttps string `json:"media_url_https"`
				Type          string `json:"type"`
				URL           string `json:"url"`
			} `json:"media"`
			URLs []struct {
				ExpandedURL string `json:"expanded_url"`
				URL         string `json:"url"`
			} `json:"urls"`
			UserMentions []struct {
				IDStr      string `json:"id_str"`
				Name       string `json:"name"`
				ScreenName string `json:"screen_name"`
			} `json:"user_mentions"`
		} `json:"entities"`
		ExtendedEntities struct {
			Media []struct {
				IDStr                    string `json:"id_str"`
				MediaURLHttps            string `json:"media_url_https"`
				ExtSensitiveMediaWarning struct {
					AdultContent    bool `json:"adult_content"`
					GraphicViolence bool `json:"graphic_violence"`
					Other           bool `json:"other"`
				} `json:"ext_sensitive_media_warning"`
				Type      string `json:"type"`
				URL       string `json:"url"`
				VideoInfo struct {
					Variants []struct {
						Bitrate int    `json:"bitrate"`
						URL     string `json:"url"`
					} `json:"variants"`
				} `json:"video_info"`
			} `json:"media"`
		} `json:"extended_entities"`
		IDStr                 string `json:"id_str"`
		InReplyToStatusIDStr  string `json:"in_reply_to_status_id_str"`
		Place                 Place  `json:"place"`
		ReplyCount            int    `json:"reply_count"`
		RetweetCount          int    `json:"retweet_count"`
		RetweetedStatusIDStr  string `json:"retweeted_status_id_str"`
		RetweetedStatusResult struct {
			Result *result `json:"result"`
		} `json:"retweeted_status_result"`
		QuotedStatusIDStr string `json:"quoted_status_id_str"`
		SelfThread        struct {
			IDStr string `json:"id_str"`
		} `json:"self_thread"`
		Time      time.Time `json:"time"`
		UserIDStr string    `json:"user_id_str"`
		Views     struct {
			State string `json:"state"`
			Count string `json:"count"`
		} `json:"ext_views"`
	}

	legacyUser struct {
		CreatedAt   string `json:"created_at"`
		Description string `json:"description"`
		Entities    struct {
			URL struct {
				Urls []struct {
					ExpandedURL string `json:"expanded_url"`
				} `json:"urls"`
			} `json:"url"`
		} `json:"entities"`
		FavouritesCount      int      `json:"favourites_count"`
		FollowersCount       int      `json:"followers_count"`
		FriendsCount         int      `json:"friends_count"`
		IDStr                string   `json:"id_str"`
		ListedCount          int      `json:"listed_count"`
		Name                 string   `json:"name"`
		Location             string   `json:"location"`
		PinnedTweetIdsStr    []string `json:"pinned_tweet_ids_str"`
		ProfileBannerURL     string   `json:"profile_banner_url"`
		ProfileImageURLHTTPS string   `json:"profile_image_url_https"`
		Protected            bool     `json:"protected"`
		ScreenName           string   `json:"screen_name"`
		StatusesCount        int      `json:"statuses_count"`
		Verified             bool     `json:"verified"`
	}

	Place struct {
		ID          string `json:"id"`
		PlaceType   string `json:"place_type"`
		Name        string `json:"name"`
		FullName    string `json:"full_name"`
		CountryCode string `json:"country_code"`
		Country     string `json:"country"`
		BoundingBox struct {
			Type        string        `json:"type"`
			Coordinates [][][]float64 `json:"coordinates"`
		} `json:"bounding_box"`
	}

	fetchProfileFunc func(query string, maxProfilesNbr int, cursor string) ([]*Profile, string, error)
	fetchTweetFunc   func(query string, maxTweetsNbr int, cursor string) ([]*Tweet, string, error)
)

// Failed Tweet
type TweetResults struct {
	Result struct {
		Tweet struct {
			RestID            string `json:"rest_id"`
			HasBirdwatchNotes bool   `json:"has_birdwatch_notes"`
			IsTranslatable    bool   `json:"is_translatable"`
			Views             struct {
				State string `json:"state"`
			} `json:"views"`
			Source        string   `json:"source"`
			AwardEligible bool     `json:"award_eligible"`
			GrantedAwards struct{} `json:"granted_awards"`
			Legacy        struct {
				BookmarkCount       int    `json:"bookmark_count"`
				Bookmarked          bool   `json:"bookmarked"`
				CreatedAt           string `json:"created_at"`
				ConversationControl struct {
					Policy                   string `json:"policy"`
					ConversationOwnerResults struct {
						Result struct {
							Type   string `json:"__typename"`
							Legacy struct {
								ScreenName string `json:"screen_name"`
							} `json:"legacy"`
						} `json:"result"`
					} `json:"conversation_owner_results"`
				} `json:"conversation_control"`
				ConversationIDStr string `json:"conversation_id_str"`
				DisplayTextRange  []int  `json:"display_text_range"`
				Entities          struct {
					UserMentions []interface{} `json:"user_mentions"`
					Urls         []interface{} `json:"urls"`
					Hashtags     []interface{} `json:"hashtags"`
					Symbols      []interface{} `json:"symbols"`
				} `json:"entities"`
				FavoriteCount  int    `json:"favorite_count"`
				Favorited      bool   `json:"favorited"`
				FullText       string `json:"full_text"`
				IsQuoteStatus  bool   `json:"is_quote_status"`
				Lang           string `json:"lang"`
				LimitedActions string `json:"limited_actions"`
				QuoteCount     int    `json:"quote_count"`
				ReplyCount     int    `json:"reply_count"`
				RetweetCount   int    `json:"retweet_count"`
				Retweeted      bool   `json:"retweeted"`
				UserIDStr      string `json:"user_id_str"`
				IDStr          string `json:"id_str"`
			} `json:"legacy"`
			QuickPromoteEligibility struct {
				Eligibility string `json:"eligibility"`
			} `json:"quick_promote_eligibility"`
		} `json:"tweet"`
		LimitedActionResults struct {
			LimitedActions []struct {
				Action string `json:"action"`
				Prompt struct {
					Type     string `json:"__typename"`
					CtaType  string `json:"cta_type"`
					Headline struct {
						Text     string        `json:"text"`
						Entities []interface{} `json:"entities"`
					} `json:"headline"`
					Subtext struct {
						Text     string        `json:"text"`
						Entities []interface{} `json:"entities"`
					} `json:"subtext"`
				} `json:"prompt"`
			} `json:"limited_actions"`
		} `json:"limitedActionResults"`
	} `json:"result"`
}

type ItemContent struct {
	TweetResults        TweetResults `json:"tweet_results"`
	TweetDisplayType    string       `json:"tweetDisplayType"`
	HasModeratedReplies bool         `json:"hasModeratedReplies"`
}

type EntryContent struct {
	ItemType    string      `json:"itemType"`
	ItemContent ItemContent `json:"itemContent"`
}

type Entry struct {
	EntryID   string       `json:"entryId"`
	SortIndex string       `json:"sortIndex"`
	Content   EntryContent `json:"content"`
}

type ThreadedConversation struct {
	Instructions []struct {
		Type    string  `json:"type"`
		Entries []Entry `json:"entries"`
	} `json:"instructions"`
}

type FetchFailed struct {
	Data struct {
		ThreadedConversationWithInjectionsV2 ThreadedConversation `json:"threaded_conversation_with_injections_v2"`
	} `json:"data"`
}

func (t *Tweet) FailMapping(fail FetchFailed) {
	t.ConversationID = fail.Data.ThreadedConversationWithInjectionsV2.Instructions[0].Entries[0].Content.ItemContent.TweetResults.Result.Tweet.Legacy.ConversationIDStr
	t.GIFs = nil
	//t.Hashtags=fail.Data.ThreadedConversationWithInjectionsV2.Instructions[0].Entries[0].Content.ItemContent.TweetResults.Result.Tweet.Legacy.Entities.Hashtags
	t.Text = fail.Data.ThreadedConversationWithInjectionsV2.Instructions[0].Entries[0].Content.ItemContent.TweetResults.Result.Tweet.Legacy.FullText
}
