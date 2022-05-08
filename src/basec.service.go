package main

type BaseCService struct {
	config     *Config
	repository *BaseCRepository
}

// func (service *BaseCService) UpdatePersonalEventByCPF(cpf CPF)

func (service *BaseCService) FindAll() *[]interface{} {
	resp, err := service.repository.FindAll()

	if err != nil {
		return nil
	}

	return resp
}

func NewBaseCService(config *Config, repository *BaseCRepository) *BaseCService {
	return &BaseCService{config: config, repository: repository}
}
