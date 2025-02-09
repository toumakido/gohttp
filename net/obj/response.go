package obj

type Response struct {
	Version string
	Status  Status
	Header  map[string]string
	Body    string
}

type Status string
