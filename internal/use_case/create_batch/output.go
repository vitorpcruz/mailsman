package create_batch

type OutputInterface interface {
	OnSucess(message string)
	OnFail(message string)
}
