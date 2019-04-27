package storage

type Lister interface {
	List() ([]byte, error)
}

type Storer interface {
	Store(server string, config []byte)
}

type Retriever interface {
	Retrieve(server string) []byte
}
