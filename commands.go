package wrapper

import (
	"context"
	"fmt"
	"reflect"
	"time"
)

// Reach objective: verify minecraft version (Bedrock vs. Java and 1.XX) to not build / prevent using certain commands

type Command interface {
	Command() string
	Events() []Event
}

// Could be a great spot to use the github.com/densestvoid/postoffice package.
// It was designed to be able to send and receive on channels identified by interface addresses.
// Each event type could be registered as an address
// Event is for any console resposne, error is for command processing only
func (w *Wrapper) ExecuteCommand(ctx context.Context, cmd Command) (Event, error) {
	// TODO: get event addresses for each event type
	var types []interface{}
	for _, event := range cmd.Events() {
		types = append(types, reflect.TypeOf(event))
	}

	// TODO: write the command to the console
	if err := w.writeToConsole(cmd.Command()); err != nil {
		return nil, err
	}

	// TODO: wait to receive on one of the event channels, and return that event
	receiveCtxTimeout, _ := context.WithTimeout(ctx, time.Second)
	mail, ok := w.eventRouting.Receive(receiveCtxTimeout, types...)
	if !ok {
		return nil, fmt.Errorf("failed to receive event related to the command")
	}

	event, ok := mail.Contents.(Event)
	if !ok {
		return nil, fmt.Errorf("received contents were not an event")
	}

	return event, nil
}

/*
attribute
advancement
ban x
ban-ip
banlist x
bossbar
clear
clone
data (get)
datapack
debug
defaultgamemode x
deop x
difficulty x
effect
enchant
execute
experience (add,query)
fill
forceload
function
gamemode
gamerule
give x
help
kick x
kill
list x
locate
locatebiome
loot
me
msg
op
pardon
particle
playsound
publish
recipe
reload
save-all x
save-off x
save-on x
say x
schedule
scoreboard
seed
setblock
setidletimeout
setworldspawn
spawnpoint
spectate
spreadplayers
stop x
stopsound
summon
tag
team
teammsg
teleport
tell x
tellraw
time
title
tp
trigger
w
weather
whitelist
worldborder
xp
*/
