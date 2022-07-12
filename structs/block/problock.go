package block

type ProBlock struct {
	Header BasicHeader `json:"header"`
	Datas  BasicDatas  `json:"data"`
}

func NewProBlock(
	header BasicHeader,
	datas BasicDatas,
) *ProBlock {
	block := &ProBlock{
		Header: header,
		Datas:  datas,
	}
	return block
}

func (block *ProBlock) Sethash(hash []byte) {
	block.Header.Sethash(hash)
}

func (block *ProBlock) Show() {

}
