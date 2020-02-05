package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/Rhymen/go-whatsapp"
	"github.com/bregydoc/gtranslate"
	"golang.org/x/text/language"
)

type messageHandler struct{}

var defaultLangFrom = language.German
var defaultLangTo = language.English

var roomLangsTo map[string]language.Tag
var roomLangsFrom map[string]language.Tag

func langInit() {
	roomLangsFrom = make(map[string]language.Tag)
	roomLangsTo = make(map[string]language.Tag)
}

func (messageHandler) HandleTextMessage(message whatsapp.TextMessage) {
	//	Example reaction
	messageText := message.Text
	if message.Info.Timestamp > startTime {
		if strings.HasPrefix(messageText, "/il") || strings.HasPrefix(messageText, "!il") {
			initEmpty(message.Info.RemoteJid)
			str := roomLangsFrom[message.Info.RemoteJid].String() + " -> " + roomLangsTo[message.Info.RemoteJid].String()
			conn.Send(whatsapp.TextMessage{
				Info: whatsapp.MessageInfo{
					RemoteJid: message.Info.RemoteJid,
				},
				Text: str,
			})
		} else if (strings.HasPrefix(messageText, "/sl") || strings.HasPrefix(messageText, "!sl")) && len(strings.Split(messageText, " ")) == 2 && message.Info.FromMe {
			languag := strings.Split(messageText, " ")[1]
			if strings.HasPrefix(messageText, "/slt") || strings.HasPrefix(messageText, "!slt") {
				retMsg := ""
				switch languag {
				case "en", "english", "englisch":
					{
						roomLangsTo[message.Info.RemoteJid] = language.English
						retMsg = "Translate to english"
					}
				case "de", "deutsch", "german":
					{
						roomLangsTo[message.Info.RemoteJid] = language.German
						retMsg = "Translate to german"
					}
				case "pl", "polish", "polnisch":
					{
						roomLangsTo[message.Info.RemoteJid] = language.Polish
						retMsg = "Translate to polish"
					}
				default:
					{
						retMsg = "Can't find language '" + languag + "'!"
					}
				}
				conn.Send(whatsapp.TextMessage{
					Info: whatsapp.MessageInfo{
						RemoteJid: message.Info.RemoteJid,
					},
					Text: retMsg,
				})
			} else if strings.HasPrefix(messageText, "/slf") || strings.HasPrefix(messageText, "!slf") {
				retMsg := ""
				switch languag {
				case "en", "english", "englisch":
					{
						roomLangsFrom[message.Info.RemoteJid] = language.English
						retMsg = "Translate from english"
					}
				case "de", "deutsch", "german":
					{
						roomLangsFrom[message.Info.RemoteJid] = language.German
						retMsg = "Translate from german"
					}
				case "pl", "polish", "polnisch":
					{
						roomLangsFrom[message.Info.RemoteJid] = language.Polish
						retMsg = "Translate from polish"
					}
				default:
					{
						retMsg = "Can't find language '" + languag + "'!"
					}
				}
				conn.Send(whatsapp.TextMessage{
					Info: whatsapp.MessageInfo{
						RemoteJid: message.Info.RemoteJid,
					},
					Text: retMsg,
				})
			}
		} else if strings.HasPrefix(messageText, "/t") || strings.HasPrefix(messageText, "!t") && len(messageText) > 1 {
			initEmpty(message.Info.RemoteJid)
			var txt string
			if message.ContextInfo.QuotedMessage != nil {
				txt = message.ContextInfo.QuotedMessage.GetConversation()
			} else if len(messageText) > 2 {
				txt = messageText[2:]
			}
			if len(txt) > 0 {
				tl, _ := gtranslate.Translate(txt, roomLangsFrom[message.Info.RemoteJid], roomLangsTo[message.Info.RemoteJid])
				msg := whatsapp.TextMessage{
					Info: whatsapp.MessageInfo{
						RemoteJid: message.Info.RemoteJid,
					},
					Text: tl,
				}
				time.Sleep((time.Duration)(rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(5)) * time.Second)
				_, err := conn.Send(msg)
				if err != nil {
					fmt.Println(err.Error())
				}
			}
		}
	}

	//fmt.Println(message)
	fmt.Printf("%s:\t%s\n", jidToName(message.Info.RemoteJid), message.Text)
}

func initEmpty(roomID string) {
	_, ok := roomLangsTo[roomID]
	if !ok {
		roomLangsTo[roomID] = defaultLangTo
	}
	_, ok = roomLangsFrom[roomID]
	if !ok {
		roomLangsFrom[roomID] = defaultLangFrom
	}
}

func (messageHandler) HandleImageMessage(message whatsapp.ImageMessage) {
	//fmt.Println(message)
}

func (messageHandler) HandleDocumentMessage(message whatsapp.DocumentMessage) {
	//fmt.Println(message)
}

func (messageHandler) HandleVideoMessage(message whatsapp.VideoMessage) {
	//fmt.Println(message)
}

func (messageHandler) HandleAudioMessage(message whatsapp.AudioMessage) {
	//fmt.Println(message)
}

func (messageHandler) HandleJSONMessage(message string) {
	//	fmt.Println(message)
}

func (messageHandler) HandleContactMessage(message whatsapp.ContactMessage) {
	//fmt.Println(message)
}

func (messageHandler) HandleError(err error) {
	//fmt.Fprintf(os.Stderr, "%v", err)
}
