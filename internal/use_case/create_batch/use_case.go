package create_batch

type UseCaseInterface interface {
	Execute(input Input, output OutputInterface)
}

type UseCase struct{}

func InstantiateUseCase() *UseCase {
	return &UseCase{}
}

func (u *UseCase) Execute(input Input, output OutputInterface) {
	output.OnSucess("this is a test")
}
