package service

type PhotoService struct{}

func NewPhotoService() *PhotoService {
	return &PhotoService{}
}

func (p *PhotoService) UploadRawPhoto(
/*
 */
) {
	/*
		загружает фото из телеграмма
		то бишь
		генерирует uuid id
		tg id
		file id
		file unique id
		размеры
		формат
		делать url по которому можно скачать
		пишет description null пока
		пишет created at
		и upload date

		потом она формирует json из этой информации
		и отправляет в kafka топик raw_photo_topic
		всё
	*/
}

func (p *PhotoService) UploadPhoto(
/*
 */
) {
	/*
		читает кафка топик result_photo_topic
		парсит json
		дописывает updated at
		и записывает в бд
	*/
}
