package postgres

type CommandEntity struct {
	Id         uint64
	TelegramId string
	Command    string
}
