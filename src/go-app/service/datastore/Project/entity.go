package Project

type Entity struct {
	ID       string `datastore:"id"` // project
	URL      string `datastore:"url"`
	Method   string `datastore:"method,noindex"`
	JSON     string `datastore:"json"`
	Schedule int    `datastore:"schedule"` // seconds
}
