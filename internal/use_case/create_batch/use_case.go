package create_batch

type UseCase struct{}

func InstantiateUseCase() *UseCase {
	return &UseCase{}
}

func (u *UseCase) Execute(input Input, output OutputInterface) {
	output.OnSucess("OK")
}
