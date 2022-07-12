package datas

type Content interface {
	Islower(data_2 Content) bool

	Ishigher(data_2 Content) bool

	Equals(data_2 Content) bool

	Tobyte() []byte
	//Islower(content *Content) bool
}
