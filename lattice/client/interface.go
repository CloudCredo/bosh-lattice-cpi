package client

type App struct {
	Id int
}

type Client interface {
	CreateApp() (App, error)
	DeleteApp(int) (bool, error)
	FindApp(int) (App, error)
	SetTags(int, []string) (bool, error)
}