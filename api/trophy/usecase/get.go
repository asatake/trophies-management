package usecase

import (
	"github.com/asatake/trophies-management/api/trophy/domain"
)

func (interactor *ContentInteractor) ShowContent(id int) (content domain.Content, err error) {
	content, err = interactor.ContentRepository.FindContentBy(id)
	return
}
