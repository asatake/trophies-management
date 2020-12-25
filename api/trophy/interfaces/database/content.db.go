package database

import "github.com/asatake/trophies-management/api/trophy/domain"

type ContentRepository struct {
	SqlHandler
}

func (repo *ContentRepository) FindContentBy(id int) (content domain.Content, err error) {
	result, err := repo.Query(
		"SELECT c.id, c.title_ja, c.title_en FROM contents c WHERE c.id = ? LIMIT 1;",
		id,
	)
	defer result.Close()
	if err != nil {
		return
	}

	for result.Next() {
		if err = result.Scan(
			&content.Id,
			&content.Title.Ja,
			&content.Title.En,
		); err != nil {
			return
		}
	}
	if content.Achievements, err = repo.FindAchievementsBy(id); err != nil {
		return
	}
	return
}

func (repo *ContentRepository) FindAchievementsBy(contentId int) (achivements []domain.Achievement, err error) {
	result, err := repo.Query(
		"SELECT a.id, a.title_ja, a.title_en FROM achievements a WHERE a.content_id = ?;",
		contentId,
	)
	defer result.Close()
	if err != nil {
		return
	}

	index := 0
	for result.Next() {
		if err = result.Scan(
			&achivements[index].Id,
			&achivements[index].Title.Ja,
			&achivements[index].Title.En,
		); err != nil {
			return
		}
		if achivements[index].Checks, err = repo.FindChecksBy(achivements[index].Id); err != nil {
			return
		}
		index++
	}
	return
}

func (repo *ContentRepository) FindChecksBy(achievementId int) (check domain.Check, err error) {
	result, err := repo.Query(
		"SELECT ck.id, ck.name, ck.completed FROM checks ck WHERE ck.achievement_id = ?;",
		achievementId,
	)
	defer result.Close()
	if err != nil {
		return
	}

	for result.Next() {
		var name string
		var completed bool
		if err = result.Scan(
			&name,
			&completed,
		); err != nil {
			return
		}
		check[name] = completed
	}
	return
}
