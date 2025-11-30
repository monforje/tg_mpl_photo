package sticker

import (
	tele "gopkg.in/telebot.v4"
)

// reg.go
var (
	StickerRegSuccess = "CAACAgIAAxkBAAETs1BpDFNglWNcAQKAF-sAAZU1--4Sc_8AAgEBAAJWnb0KIr6fDrjC5jQ2BA"
	StickerRegFail    = "CAACAgIAAxkBAAET48BpFfaohvrTZruZN9L5XoKkYw2eWAAC9AADVp29ChFYsPXZ_VVJNgQ"
)

// photo.go
var (
	StickerPhotoUploadSuccess = "CAACAgIAAxkBAAET2KdpFF4zr8zi_Ze2dO_8NoJdkEZjtwAC_gADVp29CtoEYTAu-df_NgQ"
	StickerPhotoUploadFail    = "CAACAgIAAxkBAAEUTRhpK5A0vu4wZ8BfU6waPgUsyxdSHgAC8wADVp29Cmob68TH-pb-NgQ"
)

func SendSticker(c tele.Context, stickerID string) {
	c.Send(&tele.Sticker{
		File: tele.File{FileID: stickerID},
	})
}
