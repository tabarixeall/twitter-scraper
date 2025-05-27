package twitterscraper

import "time"

type (
	// Mention type.
	Mention struct {
		ID       string
		Username string
		Name     string
	}

	// Url represents a URL with display, expanded, and index data.
	Url struct {
		DisplayURL  string `json:"display_url"`
		ExpandedURL string `json:"expanded_url"`
		URL         string `json:"url"`
		Indices     []int  `json:"indices"`
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
		HLSURL  string
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

	ScheduledTweet struct {
		ID        string
		State     string
		ExecuteAt time.Time
		Text      string
		Videos    []Video
		Photos    []Photo
		GIFs      []GIF
	}

	ExtendedMedia struct {
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
				Type    string `json:"content_type"`
				Bitrate int    `json:"bitrate"`
				URL     string `json:"url"`
			} `json:"variants"`
		} `json:"video_info"`
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
			URLs         []Url `json:"urls"`
			UserMentions []struct {
				IDStr      string `json:"id_str"`
				Name       string `json:"name"`
				ScreenName string `json:"screen_name"`
			} `json:"user_mentions"`
		} `json:"entities"`
		ExtendedEntities struct {
			Media []ExtendedMedia `json:"media"`
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
		FavouritesCount         int      `json:"favourites_count"`
		FollowersCount          int      `json:"followers_count"`
		FriendsCount            int      `json:"friends_count"`
		IDStr                   string   `json:"id_str"`
		ListedCount             int      `json:"listed_count"`
		Name                    string   `json:"name"`
		Location                string   `json:"location"`
		PinnedTweetIdsStr       []string `json:"pinned_tweet_ids_str"`
		ProfileBannerURL        string   `json:"profile_banner_url"`
		ProfileImageURLHTTPS    string   `json:"profile_image_url_https"`
		Protected               bool     `json:"protected"`
		ScreenName              string   `json:"screen_name"`
		StatusesCount           int      `json:"statuses_count"`
		Verified                bool     `json:"verified"`
		FollowedBy              bool     `json:"followed_by"`
		Following               bool     `json:"following"`
		CanDm                   bool     `json:"can_dm"`
		CanMediaTag             bool     `json:"can_media_tag"`
		DefaultProfile          bool     `json:"default_profile"`
		DefaultProfileImage     bool     `json:"default_profile_image"`
		FastFollowersCount      int      `json:"fast_followers_count"`
		HasCustomTimelines      bool     `json:"has_custom_timelines"`
		IsTranslator            bool     `json:"is_translator"`
		MediaCount              int      `json:"media_count"`
		NeedsPhoneVerification  bool     `json:"needs_phone_verification"`
		NormalFollowersCount    int      `json:"normal_followers_count"`
		PossiblySensitive       bool     `json:"possibly_sensitive"`
		ProfileInterstitialType string   `json:"profile_interstitial_type"`
		TranslatorType          string   `json:"translator_type"`
		WantRetweets            bool     `json:"want_retweets"`
		WithheldInCountries     []string `json:"withheld_in_countries"`
	}

	legacyUserV2 struct {
		FollowedBy          bool   `json:"followed_by"`
		Following           bool   `json:"following"`
		CanDm               bool   `json:"can_dm"`
		CanMediaTag         bool   `json:"can_media_tag"`
		CreatedAt           string `json:"created_at"`
		DefaultProfile      bool   `json:"default_profile"`
		DefaultProfileImage bool   `json:"default_profile_image"`
		Description         string `json:"description"`
		Entities            struct {
			Description struct {
				Urls []Url `json:"urls"`
			} `json:"description"`
			URL struct {
				Urls []Url `json:"urls"`
			} `json:"url"`
		} `json:"entities"`
		FastFollowersCount      int           `json:"fast_followers_count"`
		FavouritesCount         int           `json:"favourites_count"`
		FollowersCount          int           `json:"followers_count"`
		FriendsCount            int           `json:"friends_count"`
		HasCustomTimelines      bool          `json:"has_custom_timelines"`
		IsTranslator            bool          `json:"is_translator"`
		ListedCount             int           `json:"listed_count"`
		Location                string        `json:"location"`
		MediaCount              int           `json:"media_count"`
		Name                    string        `json:"name"`
		NormalFollowersCount    int           `json:"normal_followers_count"`
		PinnedTweetIdsStr       []string      `json:"pinned_tweet_ids_str"`
		PossiblySensitive       bool          `json:"possibly_sensitive"`
		ProfileBannerURL        string        `json:"profile_banner_url"`
		ProfileImageURLHTTPS    string        `json:"profile_image_url_https"`
		ProfileInterstitialType string        `json:"profile_interstitial_type"`
		ScreenName              string        `json:"screen_name"`
		StatusesCount           int           `json:"statuses_count"`
		TranslatorType          string        `json:"translator_type"`
		URL                     string        `json:"url"`
		Verified                bool          `json:"verified"`
		WantRetweets            bool          `json:"want_retweets"`
		WithheldInCountries     []interface{} `json:"withheld_in_countries"`
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

	legacyExtendedProfile struct {
		Birthdate struct {
			Day            int    `json:"day"`
			Month          int    `json:"month"`
			Year           int    `json:"year"`
			Visibility     string `json:"visibility"`
			YearVisibility string `json:"year_visibility"`
		} `json:"birthdate"`
	}

	verificationInfo struct {
		IsIdentityVerified bool `json:"is_identity_verified"`
		Reason             struct {
			Description struct {
				Text     string `json:"text"`
				Entities []struct {
					FromIndex int `json:"from_index"`
					ToIndex   int `json:"to_index"`
					Ref       struct {
						URL     string `json:"url"`
						URLType string `json:"url_type"`
					} `json:"ref"`
				} `json:"entities"`
			} `json:"description"`
			VerifiedSinceMsec string `json:"verified_since_msec"`
		} `json:"reason"`
	}

	highlightsInfo struct {
		CanHighlightTweets bool   `json:"can_highlight_tweets"`
		HighlightedTweets  string `json:"highlighted_tweets"`
	}
)
