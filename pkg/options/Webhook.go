package options

import "time"

type WebhookDiscord struct {
	Content string                 `json:"content"`
	Embeds  []OptionsEmbedsDiscord `json:"embeds"`
}

type OptionsEmbedsDiscord struct {
	Title       string                     `json:"title"`
	Description string                     `json:"description"`
	Url         string                     `json:"url"`
	Color       int                        `json:"color"`
	Fields      []FieldEmbedsDiscord       `json:"fields"`
	Author      AuthorOptionsEmbedsDiscord `json:"author"`
	Footer      FooterEmbedsDiscord        `json:"footer"`
	Timestamp   time.Time                  `json:"timestamp"`
	Image       ImageEmbedsDiscord         `json:"image"`
	Thumbnail   ThumbnailEmbedsDiscord     `json:"thumbnail"`
}

type AuthorOptionsEmbedsDiscord struct {
	Name    string `json:"name"`
	Url     string `json:"url"`
	IconUrl string `json:"icon_url"`
}

type FieldEmbedsDiscord struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

type FooterEmbedsDiscord struct {
	Text    string `json:"text"`
	IconUrl string `json:"icon_url"`
}

type ImageEmbedsDiscord struct {
	Url string `json:"url"`
}

type ThumbnailEmbedsDiscord struct {
	Url string `json:"url"`
}

// Struct slack object

type WebhookSlack struct {
	Text   string              `json:"text"`
	Blocks []OptionsBlockSlack `json:"blocks"`
}

type OptionsBlockSlack struct {
	Type     string                      `json:"type"`
	Text     *OptionsTextMrkSlack        `json:"text,omitempty"`
	Elements []OptionsElementsBlockSlack `json:"elements,omitempty"`
}

type OptionsTextMrkSlack struct {
	Type  string `json:"type"`
	Text  string `json:"text"`
	Emoji bool   `json:"emoji,omitempty"` // omitempty porque mrkdwn no usa emoji
}

type OptionsElementsBlockSlack struct {
	Type  string `json:"type"`
	Text  string `json:"text"`
	Emoji bool   `json:"emoji,omitempty"`
}
