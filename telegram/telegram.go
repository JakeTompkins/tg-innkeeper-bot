package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

var apiKey = os.Getenv("API_KEY")
var botUrl = os.Getenv("BOT_URL")
var listenPort = os.Getenv("PORT")

type ITelegramBot interface {
	Send(request telegramBotRequest) telegramBotResponse
}
type telegramBot struct {
	root_url string

	updateParser *updateParser
}

func NewTelegramBot() *telegramBot {
	var telegramBot telegramBot

	telegramBot.root_url = fmt.Sprintf("https://api.telegram.org/bot%s/", apiKey)
	telegramBot.updateParser = newUpdateParser()

	return &telegramBot
}

type telegramBotResponse struct {
	Code  int    `json:"code,omitempty"`
	Error error  `json:"error,omitempty"`
	Data  string `json:"data,omitempty"`
}

func newErrorResponse(error error) telegramBotResponse {
	var response telegramBotResponse

	response.Code = 500
	response.Error = error
	response.Data = ""

	return response
}

func newSuccessResponse(data string) telegramBotResponse {
	var response telegramBotResponse

	response.Code = 200
	response.Error = nil
	response.Data = data

	return response
}

type telegramBotRequest struct {
	Command string      `json:"command,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}

func newRequest(command string, payload interface{}) telegramBotRequest {
	var request telegramBotRequest

	request.Command = command
	request.Payload = payload

	return request
}

type user struct {
	Id                      int    `json:"id,omitempty"`
	IsBot                   bool   `json:"is_bot,omitempty"`
	FirstName               string `json:"first_name,omitempty"`
	Username                string `json:"username,omitempty"`
	LanguageCode            string `json:"language_code,omitempty"`
	IsPremium               bool   `json:"is_premium,omitempty"`
	AddedToAttachmentMenu   bool   `json:"added_to_attachment_menu,omitempty"`
	CanJoinGroups           bool   `json:"can_join_groups,omitempty"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages,omitempty"`
	SupportsInlineQueries   bool   `json:"supports_inline_queries,omitempty"`
}

type chatPermissions struct {
	CanSendMessages       bool `json:"can_send_messages,omitempty"`
	CanSendAudios         bool `json:"can_send_audios,omitempty"`
	CanSendDocuments      bool `json:"can_send_documents,omitempty"`
	CanSendPhotos         bool `json:"can_send_photos,omitempty"`
	CanSendVideos         bool `json:"can_send_videos,omitempty"`
	CanSendVideoNotes     bool `json:"can_send_video_notes,omitempty"`
	CanSendVoiceNotes     bool `json:"can_send_voice_notes,omitempty"`
	CanSendPolls          bool `json:"can_send_polls,omitempty"`
	CanSendOtherMessages  bool `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`
	CaanChangeInfo        bool `json:"caan_change_info,omitempty"`
	CanInviteUsers        bool `json:"can_invite_users,omitempty"`
	CanPinMessages        bool `json:"can_pin_messages,omitempty"`
	CanManageTopics       bool `json:"can_manage_topics,omitempty"`
}

type location struct {
	Longitude            float32 `json:"longitude,omitempty"`
	Latitude             float32 `json:"latitude,omitempty"`
	HorizontalAccuracy   float32 `json:"horizontal_accuracy,omitempty"`
	LivePeriod           int     `json:"live_period,omitempty"`
	Heading              int     `json:"heading,omitempty"`
	ProximityAlertRadius int     `json:"proximity_alert_radius,omitempty"`
}

type chatLocation struct {
	Location *location `json:"location,omitempty"`
	address  string    `json:"address,omitempty"`
}

type chat struct {
	Id                                int              `json:"id,omitempty"`
	Type                              string           `json:"type,omitempty"`
	Title                             string           `json:"title,omitempty"`
	Username                          string           `json:"username,omitempty"`
	FirstName                         string           `json:"first_name,omitempty"`
	LastName                          string           `json:"last_name,omitempty"`
	IsForum                           string           `json:"is_forum,omitempty"`
	ActiveUsernames                   []string         `json:"active_usernames,omitempty"`
	EmojistatusCustomEmojiId          string           `json:"emojistatus_custom_emoji_id,omitempty"`
	EmojiStatusExpirationDate         string           `json:"emoji_status_expiration_date,omitempty"`
	Bio                               string           `json:"bio,omitempty"`
	HasPrivateForwards                bool             `json:"has_private_forwards,omitempty"`
	HasRestrictedVoidAndVideoMessages bool             `json:"has_restricted_void_and_video_messages,omitempty"`
	JoinToSendMessages                bool             `json:"join_to_send_messages,omitempty"`
	JoinByRequest                     bool             `json:"join_by_request,omitempty"`
	Description                       string           `json:"description,omitempty"`
	InviteLink                        string           `json:"invite_link,omitempty"`
	PinnedMessage                     *message         `json:"pinned_message,omitempty"`
	Permissions                       *chatPermissions `json:"permissions,omitempty"`
	SlowModeDelay                     int              `json:"slow_mode_delay,omitempty"`
	MessageAutoDeleteTime             int              `json:"message_auto_delete_time,omitempty"`
	HasAggressiveAntiSpamEnabled      bool             `json:"has_aggressive_anti_spam_enabled,omitempty"`
	HasHiddenMembers                  bool             `json:"has_hidden_members,omitempty"`
	HasProtectedContent               bool             `json:"has_protected_content,omitempty"`
	StickerSetName                    string           `json:"sticker_set_name,omitempty"`
	CanSetStickerSet                  bool             `json:"can_set_sticker_set,omitempty"`
	LinkedChatId                      int              `json:"linked_chat_id,omitempty"`
	location                          *chatLocation    `json:"location,omitempty"`
}

type messageEntity struct {
	Type          string `json:"type,omitempty"`
	Offset        int    `json:"offset,omitempty"`
	Length        int    `json:"length,omitempty"`
	Url           string `json:"url,omitempty"`
	User          *user  `json:"user,omitempty"`
	Language      string `json:"language,omitempty"`
	CustomEmojiId string `json:"custom_emoji_id,omitempty"`
}

type message struct {
	MessageId            int              `json:"message_id,omitempty"`
	MessageThreadId      int              `json:"message_thread_id,omitempty"`
	From                 *user            `json:"from,omitempty"`
	SenderChat           *chat            `json:"sender_chat,omitempty"`
	Date                 int              `json:"date,omitempty"`
	Chat                 *chat            `json:"chat,omitempty"`
	ForwardFrom          *user            `json:"forward_from,omitempty"`
	ForwardFromChat      *chat            `json:"forward_from_chat,omitempty"`
	ForwardFromMessageId int              `json:"forward_from_message_id,omitempty"`
	ForwardSignature     string           `json:"forward_signature,omitempty"`
	ForwardSenderName    string           `json:"forward_sender_name,omitempty"`
	ForwardDate          int              `json:"forward_date,omitempty"`
	IsTopicMessage       bool             `json:"is_topic_message,omitempty"`
	IsAutomaticForward   bool             `json:"is_automatic_forward,omitempty"`
	ReplyToMessage       *message         `json:"reply_to_message,omitempty"`
	ViaBot               *user            `json:"via_bot,omitempty"`
	EditDate             int              `json:"edit_date,omitempty"`
	HasProtectedContent  bool             `json:"has_protected_content,omitempty"`
	MediaGroupId         string           `json:"media_group_id,omitempty"`
	AuthorSignature      string           `json:"author_signature,omitempty"`
	Text                 string           `json:"text,omitempty"`
	Entities             []*messageEntity `json:"entities,omitempty"`
	// NOTE: This is not complete, but I'm not sure which props I actually need
	// TODO: Update this as we go
}

type inlineQuery struct {
	Id       string    `json:"id,omitempty"`
	From     *user     `json:"from,omitempty"`
	Query    string    `json:"query,omitempty"`
	Offset   string    `json:"offset,omitempty"`
	ChatType string    `json:"chat_type,omitempty"`
	Location *location `json:"location,omitempty"`
}

type chosenInlineResult struct {
	ResultId        string    `json:"result_id,omitempty"`
	From            *user     `json:"from,omitempty"`
	Location        *location `json:"location,omitempty"`
	InlineMessageId string    `json:"inline_message_id,omitempty"`
	Query           string    `json:"query,omitempty"`
}

type callbackQuery struct {
	Id              string   `json:"id,omitempty"`
	From            *user    `json:"from,omitempty"`
	Message         *message `json:"message,omitempty"`
	InlineMessageId string   `json:"inline_message_id,omitempty"`
	ChatInstance    string   `json:"chat_instance,omitempty"`
	Data            string   `json:"data,omitempty"`
	GameShortName   string   `json:"game_short_name,omitempty"`
}

type Poll struct {
	Id                    string `json:"id,omitempty"`
	Question              string `json:"question,omitempty"`
	TotalVoterCount       int    `json:"total_voter_count,omitempty"`
	IsClosed              bool   `json:"is_closed,omitempty"`
	IsAnonymous           bool   `json:"is_anonymous,omitempty"`
	Type                  string `json:"type,omitempty"`
	AllowsMultipleAnswers bool   `json:"allows_multiple_answers,omitempty"`
	CorrectOptionId       int    `json:"correct_option_id,omitempty"`
	Explanation           string `json:"explanation,omitempty"`
	OpenPeriod            int    `json:"open_period,omitempty"`
	CloseDate             int    `json:"close_date,omitempty"`
}

type PollAnswer struct {
	PollId    string `json:"poll_id,omitempty"`
	VoterChat *chat  `json:"voter_chat,omitempty"`
	User      *user  `json:"user,omitempty"`
	OptionIds []int  `json:"option_ids,omitempty"`
}

type update struct {
	UpdateId          int                 `json:"update_id,omitempty"`
	Message           *message            `json:"message,omitempty"`
	EditedMessage     *message            `json:"edited_message,omitempty"`
	ChannelPost       *message            `json:"channel_post,omitempty"`
	EditedChannelPost *message            `json:"edited_channel_post,omitempty"`
	InlineQuery       *inlineQuery        `json:"inline_query,omitempty"`
	ChoseInlineResult *chosenInlineResult `json:"chose_inline_result,omitempty"`
	CallbackQuery     *callbackQuery      `json:"callback_query,omitempty"`
	Poll              *Poll               `json:"poll,omitempty"`
	PollAnser         *PollAnswer         `json:"poll_anser,omitempty"`
}

type setWebhookPayload struct {
	Url string
}

func (t *telegramBot) Send(request telegramBotRequest) telegramBotResponse {
	url := t.root_url + request.Command

	bodyJson, err := json.Marshal(request.Payload)

	if err != nil {
		return newErrorResponse(err)
	}

	res, err := http.Post(url, "application/json", bytes.NewBuffer(bodyJson))

	if err != nil {
		return newErrorResponse(err)
	}

	resBody, err := io.ReadAll(res.Body)

	return newSuccessResponse(string(resBody))
}

func (t *telegramBot) RegisterWebhook() telegramBotResponse {
	request := newRequest("setWebhook", setWebhookPayload{
		Url: botUrl + "/webhook",
	})

	response := t.Send(request)

	return response
}

func HandleUpdate(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		return
	}

	var update *update

	err = json.Unmarshal(body, update)

	if err != nil {
		return
	}

}

func (t *telegramBot) Listen() error {
	mux := http.NewServeMux()
	// mux.HandleFunc("/webhook", HandleUpdate)
	mux.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Pong!")
	})

	response := t.RegisterWebhook()

	if response.Error != nil {
		fmt.Println(response.Error)
	}

	return nil
}
