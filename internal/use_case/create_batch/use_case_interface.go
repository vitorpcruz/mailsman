package create_batch

type UseCaseInterface interface {
	Execute(input Input, output OutputInterface)
}
