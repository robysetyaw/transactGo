package repository

import (
	"encoding/json"
	"errors"
	"os"
	"transactgo/app/model"
)

type MerchantRepository interface {
	GetByID(id int) (*model.Merchant, error)
	Save(merchant *model.Merchant) error
	Update(merchant *model.Merchant) error
	Delete(id int) error
	FindAll() ([]model.Merchant, error)
}

type merchantRepository struct {
	merchants []model.Merchant
}

func NewMerchantRepository() (MerchantRepository, error) {
	repo := &merchantRepository{}

	// Open the JSON file
	file, err := os.Open("data/merchants.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Decode the file into the merchants slice
	err = json.NewDecoder(file).Decode(&repo.merchants)
	if err != nil {
		return nil, err
	}

	return repo, nil
}

func (r *merchantRepository) GetByID(id int) (*model.Merchant, error) {
	for _, merchant := range r.merchants {
		if merchant.ID == id {
			return &merchant, nil
		}
	}
	return nil, errors.New("merchant not found")
}

func (r *merchantRepository) Save(merchant *model.Merchant) error {
	r.merchants = append(r.merchants, *merchant)

	file, err := os.OpenFile("data/merchants.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(r.merchants)
}

func (r *merchantRepository) Update(merchant *model.Merchant) error {
	for i, m := range r.merchants {
		if m.ID == merchant.ID {
			r.merchants[i] = *merchant
			break
		}
	}

	file, err := os.OpenFile("data/merchants.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(r.merchants)
}

func (r *merchantRepository) Delete(id int) error {
	for i, m := range r.merchants {
		if m.ID == id {
			r.merchants = append(r.merchants[:i], r.merchants[i+1:]...)
			break
		}
	}

	file, err := os.OpenFile("data/merchants.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(r.merchants)
}

func (r *merchantRepository) FindAll() ([]model.Merchant, error) {
	return r.merchants, nil
}
