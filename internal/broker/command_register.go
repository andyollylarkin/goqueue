package broker

import (
	"goqueue/pkg/protocol_commands"
	"net/rpc"
)

func RegisterCommands() error {
	connectCommand := &protocol_commands.ConnectCommand{}
	err := rpc.Register(connectCommand)
	if err != nil {
		return err
	}
	return nil
}
