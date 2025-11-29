package sticker

import (
	tele "gopkg.in/telebot.v4"
)

// reg.go
var (
	StickerRegSuccess = "CAACAgIAAxkBAAETs1BpDFNglWNcAQKAF-sAAZU1--4Sc_8AAgEBAAJWnb0KIr6fDrjC5jQ2BA"
	StickerRegFail    = "CAACAgIAAxkBAAET48BpFfaohvrTZruZN9L5XoKkYw2eWAAC9AADVp29ChFYsPXZ_VVJNgQ"
)

func SendSticker(c tele.Context, stickerID string) {
	c.Send(&tele.Sticker{
		File: tele.File{FileID: stickerID},
	})
}
