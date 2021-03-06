package wcbot

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func TextHandle(c *gin.Context) {
	//消息处理
	if err := handleTextMsg(c); err != nil {
		c.Status(http.StatusBadRequest)
		_, _ = c.Writer.Write(nil)
	}

	c.Status(http.StatusOK)
	_, _ = c.Writer.Write(nil)
}

func handleTextMsg(c *gin.Context) error {
	to := c.Query("to")
	word := c.Query("word")

	if to == "" && word == "" {
		logrus.Error("param error")
		return errors.New("param error")
	}

	if to == "" {
		if ok := WechatBot.SendMsgByUid(word, "filehelper"); !ok {
			logrus.Error("send msg error")
			return errors.New("send msg error")
		}
	} else {
		if ok := WechatBot.SendMsg(to, word, false); !ok {
			logrus.Error("send msg error")
			return errors.New("send msg error")
		}
	}
	return nil
}
