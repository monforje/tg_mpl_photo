package fileurl

func GetTgFileURL(token, filePath string) string {
	return "https://api.telegram.org/file/bot" + token + "/" + filePath
}
