package create_batch

type UseCaseInterface interface {
	Execute(input Input) error
}

type UseCase struct{}

func InstantiateUseCase() *UseCase {
	return &UseCase{}
}

func (u *UseCase) Execute(input Input) error {
	return nil
}
