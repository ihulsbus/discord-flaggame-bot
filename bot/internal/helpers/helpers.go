/*
Copyright (c) 2023 Ian Hulsbus

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package helpers

import (
	"time"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

const (
	InsufficientPermissions string = "Insufficient permissions for the command"
	UnknownOption           string = "unknown option provided to command. Aborting."
)

func SetupLogLevels() map[string]log.Level {
	logLevels := make(map[string]log.Level, 7)

	logLevels["PANIC"] = log.PanicLevel
	logLevels["FATAL"] = log.FatalLevel
	logLevels["ERROR"] = log.ErrorLevel
	logLevels["WARN"] = log.WarnLevel
	logLevels["INFO"] = log.InfoLevel
	logLevels["DEBUG"] = log.DebugLevel
	logLevels["TRACE"] = log.TraceLevel

	return logLevels
}

func CreateEmbed(countryName string, imageURL string) *discordgo.MessageEmbed {
	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{},
		Color:  0x00ff00, // Green
		Image: &discordgo.MessageEmbedImage{
			URL: imageURL,
		},
		Thumbnail: nil,
		Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
		Title:     countryName,
	}

	return embed
}

func sendInteraction(s *discordgo.Session, i *discordgo.InteractionCreate, resp *discordgo.InteractionResponse) error {
	err := s.InteractionRespond(i.Interaction, resp)
	if err != nil {
		return err
	}
	return nil
}

func SendInteractionResponse(s *discordgo.Session, i *discordgo.InteractionCreate, message string) error {
	resp := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: message,
		},
	}

	return sendInteraction(s, i, &resp)
}

func SendInteractionEmbedResponse(s *discordgo.Session, i *discordgo.InteractionCreate, embeds []*discordgo.MessageEmbed) error {
	resp := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: embeds,
		},
	}

	return sendInteraction(s, i, &resp)
}

func SendInteractionAwaitResponse(s *discordgo.Session, i *discordgo.InteractionCreate, message string) error {
	resp := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: message,
		},
	}

	return sendInteraction(s, i, &resp)
}

func SendInteractionAwaitUpdate(s *discordgo.Session, i *discordgo.InteractionCreate, message string) error {
	resp := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredMessageUpdate,
		Data: &discordgo.InteractionResponseData{
			Content: message,
		},
	}

	return sendInteraction(s, i, &resp)
}

func SendInteractionPingResponse(s *discordgo.Session, i *discordgo.InteractionCreate) {
	resp := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponsePong,
	}

	sendInteraction(s, i, &resp)
}
