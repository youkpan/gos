package player

import (
	"goslib/gen_server"
	"goslib/session_utils"
	"goslib/logger"
	"github.com/kataras/iris/core/errors"
)

const SERVER = "__player_manager_server__"

/*
   GenServer Callbacks
*/
type PlayerManager struct {
}

func StartPlayerManager() {
	gen_server.Start(SERVER, new(PlayerManager))
}

func StartPlayer(accountId string) error {
	session, err := session_utils.Find(accountId)
	if err != nil {
		logger.ERR("StartPlayer failed: ", err)
		return err
	}
	if session.GameAppId == CurrentGameAppId {
		_, err = gen_server.Call(SERVER, "StartPlayer", accountId)
		return err
	} else {
		logger.ERR("StartPlayer failed: player not belongs to this server!")
		return errors.New("player not belongs to this server!")
	}
}

func (self *PlayerManager) Init(args []interface{}) (err error) {
	return nil
}

func (self *PlayerManager) HandleCast(args []interface{}) {
}

func (self *PlayerManager) HandleCall(args []interface{}) (interface{}, error) {
	handle := args[0].(string)
	if handle == "StartPlayer" {
		accountId := args[1].(string)
		if !gen_server.Exists(accountId) {
			gen_server.Start(accountId, new(Player), accountId)
		}
	}
	return nil, nil
}

func (self *PlayerManager) Terminate(reason string) (err error) {
	return nil
}