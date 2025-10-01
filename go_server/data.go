package main

import (
	"bytes"
	"encoding/json"
	"time"
)

func PrepJson(structToMarshal any) (reader *bytes.Reader, err error) {
	jsonBytes, err := json.Marshal(structToMarshal)
	if err != nil {
		return nil, err
	}

	reader = bytes.NewReader(jsonBytes)
	return reader, nil
}

type LoginStruct struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

type SessionStruct struct {
	Did    string `json:"did"`
	DidDoc struct {
		Context            []string `json:"@context"`
		ID                 string   `json:"id"`
		AlsoKnownAs        []string `json:"alsoKnownAs"`
		VerificationMethod []struct {
			ID                 string `json:"id"`
			Type               string `json:"type"`
			Controller         string `json:"controller"`
			PublicKeyMultibase string `json:"publicKeyMultibase"`
		} `json:"verificationMethod"`
		Service []struct {
			ID              string `json:"id"`
			Type            string `json:"type"`
			ServiceEndpoint string `json:"serviceEndpoint"`
		} `json:"service"`
	} `json:"didDoc"`
	Handle          string `json:"handle"`
	Email           string `json:"email"`
	EmailConfirmed  bool   `json:"emailConfirmed"`
	EmailAuthFactor bool   `json:"emailAuthFactor"`
	AccessJwt       string `json:"accessJwt"`
	RefreshJwt      string `json:"refreshJwt"`
	Active          bool   `json:"active"`
}

type PostStruct struct {
	Repo       string       `json:"repo"`
	Collection string       `json:"collection"`
	Record     RecordStruct `json:"record"`
}

type RecordStruct struct {
	Text      string `json:"text"`
	CreatedAt string `json:"createdAt"`
}

type GetTimelineStruct struct {
	Limit string `json:"limit"`
}

type FeedResponse struct {
	Feed   []FeedItem `json:"feed"`
	Cursor string     `json:"cursor"`
}

type FeedItem struct {
	Post  Post       `json:"post"`
	Reply *ReplyInfo `json:"reply,omitempty"`
}

type ReplyInfo struct {
	Root   PostView `json:"root"`
	Parent PostView `json:"parent"`
}

type PostView struct {
	URI           string    `json:"uri"`
	Cid           string    `json:"cid"`
	Author        Author    `json:"author"`
	Record        Record    `json:"record"`
	BookmarkCount int       `json:"bookmarkCount"`
	ReplyCount    int       `json:"replyCount"`
	RepostCount   int       `json:"repostCount"`
	LikeCount     int       `json:"likeCount"`
	QuoteCount    int       `json:"quoteCount"`
	IndexedAt     time.Time `json:"indexedAt"`
	Viewer        Viewer    `json:"viewer"`
	Labels        []any     `json:"labels"`
	Type          string    `json:"$type"`
}

type Post struct {
	URI           string    `json:"uri"`
	Cid           string    `json:"cid"`
	Author        Author    `json:"author"`
	Record        Record    `json:"record"`
	Embed         *Embed    `json:"embed,omitempty"`
	BookmarkCount int       `json:"bookmarkCount"`
	ReplyCount    int       `json:"replyCount"`
	RepostCount   int       `json:"repostCount"`
	LikeCount     int       `json:"likeCount"`
	QuoteCount    int       `json:"quoteCount"`
	IndexedAt     time.Time `json:"indexedAt"`
	Viewer        Viewer    `json:"viewer"`
	Labels        []any     `json:"labels"`
}

type Author struct {
	Did         string       `json:"did"`
	Handle      string       `json:"handle"`
	DisplayName string       `json:"displayName"`
	Avatar      string       `json:"avatar"`
	Associated  *Associated  `json:"associated,omitempty"`
	Viewer      AuthorViewer `json:"viewer"`
	Labels      []any        `json:"labels"`
	CreatedAt   time.Time    `json:"createdAt"`
}

type Associated struct {
	Chat                 *Chat                 `json:"chat,omitempty"`
	ActivitySubscription *ActivitySubscription `json:"activitySubscription,omitempty"`
}

type Chat struct {
	AllowIncoming string `json:"allowIncoming"`
}

type ActivitySubscription struct {
	AllowSubscriptions string `json:"allowSubscriptions"`
}

type AuthorViewer struct {
	Muted     bool   `json:"muted"`
	BlockedBy bool   `json:"blockedBy"`
	Following string `json:"following"`
}

type Record struct {
	Type      string       `json:"$type"`
	CreatedAt time.Time    `json:"createdAt"`
	Embed     *RecordEmbed `json:"embed,omitempty"`
	Langs     []string     `json:"langs"`
	Text      string       `json:"text"`
	Reply     *ReplyRef    `json:"reply,omitempty"`
}

type RecordEmbed struct {
	Type     string    `json:"$type"`
	External *External `json:"external,omitempty"`
}

type ReplyRef struct {
	Parent PostRef `json:"parent"`
	Root   PostRef `json:"root"`
}

type PostRef struct {
	Cid string `json:"cid"`
	URI string `json:"uri"`
}

type External struct {
	Description string `json:"description"`
	Thumb       *Thumb `json:"thumb,omitempty"`
	Title       string `json:"title"`
	URI         string `json:"uri"`
}

type Thumb struct {
	Type     string   `json:"$type"`
	Ref      ThumbRef `json:"ref"`
	MimeType string   `json:"mimeType"`
	Size     int      `json:"size"`
}

type ThumbRef struct {
	Link string `json:"$link"`
}

type Embed struct {
	Type     string         `json:"$type"`
	External *EmbedExternal `json:"external,omitempty"`
}

type EmbedExternal struct {
	URI         string `json:"uri"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Thumb       string `json:"thumb"`
}

type Viewer struct {
	Bookmarked        bool `json:"bookmarked"`
	ThreadMuted       bool `json:"threadMuted"`
	EmbeddingDisabled bool `json:"embeddingDisabled"`
}
