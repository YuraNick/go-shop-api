package verify

import "3-validation-api/pkg/jsonfile"

type VerifyRepository struct {
	Repo *jsonfile.JsonFile
}

func NewVerifyRepository(repo *jsonfile.JsonFile) *VerifyRepository {
	return &VerifyRepository{
		Repo: repo,
	}
}

func (repo *VerifyRepository) Create(email *EmailHash) (*EmailHash, error) {
	err := repo.Repo.WriteJSONByKey(email.Hash, email.Email)
	if err != nil {
		return nil, err
	}
	return email, nil
}

func (repo *VerifyRepository) GetByHash(hash string) (*EmailHash, error) {
	var email string
	err := repo.Repo.ReadJSONByKey(hash, &email)
	if err != nil {
		return nil, err
	}
	return &EmailHash{
		Email: email,
		Hash:  hash,
	}, nil
}
