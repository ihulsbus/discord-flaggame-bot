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

package flaggame

import (
	"context"
	c "flaggame/internal/config"
	h "flaggame/internal/handlers"

	"github.com/bwmarrin/discordgo"
)

var (
	falseBool = false

	started bool = false

	// less typing by referencing
	discordClient = c.Configuration.Discord.Client
	logger        = c.Configuration.Global.Logger

	// All commands and options must have a description
	// Commands/options without description will fail the registration
	// of the command.
	commands = []*discordgo.ApplicationCommand{
		{
			Name:         "flag",
			Description:  "Retrieve a random flag",
			DMPermission: &falseBool,
		},
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"flag": h.GetFlag,
	}
)

func Flaggame(ctx context.Context) {
	logger.Info("starting flaggame")

	// flaggameCtx, flaggameCancel := context.WithCancel(ctx)
	// defer flaggameCancel()

	// Register handler for incoming commands
	discordClient.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if handler, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			handler(s, i)
		}
	})

	discordClient.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		logger.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	err := discordClient.Open()
	if err != nil {
		logger.Fatalf("Error opening discord session: %v", err)
	}

	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := discordClient.ApplicationCommandCreate(discordClient.State.User.ID, "", v)
		if err != nil {
			logger.Errorf("Cannot create command '%v'. Error: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}
	defer discordClient.Close()

	started = true
	logger.Infoln("Bot is now running. Press CTRL-C to exit.")

	// wait for the context to report done and then do a cleanup
	<-ctx.Done()

	/*
		We need to fetch the commands, since deleting requires the command ID.
		We are doing this from the returned commands on line 70, because using
		this will delete all the commands, which might not be desirable, so we
		are deleting only the commands that we added.
	*/
	logger.Info("deregistering commands")

	registeredCommands, err = discordClient.ApplicationCommands(discordClient.State.User.ID, "")
	if err != nil {
		logger.Fatalf("Could not fetch registered commands: %v", err)
	}

	for _, v := range registeredCommands {
		err := discordClient.ApplicationCommandDelete(discordClient.State.User.ID, "", v.ID)
		if err != nil {
			logger.Panicf("Cannot delete '%v' command: %v", v.Name, err)
		}
	}
}
