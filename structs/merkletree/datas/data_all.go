package datas

import "ProQueries/crypto/bls"

type Data_All struct {
	data_content *Data_R
	data_sign    *Data_S
}

func NewDataAll(s *Data_R) *Data_All {
	data_a := &Data_All{
		data_content: s,
	}
	return data_a
}

func (data *Data_All) Showdata() *Data_R {
	return data.data_content
}

func (data *Data_All) Showsign() []byte {
	if data.data_sign == nil {
		data.data_sign = NewnulldataS()
		data.data_sign.Hashsign(data.data_content.Tobyte())
	}
	return data.data_sign.Showsign()
}

func (data *Data_All) Getds() Data_S {
	return *data.data_sign
}

func (data *Data_All) Getr() *Data_R {
	return data.data_content
}

func NewnullDataAll() *Data_All {
	data_a := &Data_All{
		data_content: NewnullDataR(),
	}
	return data_a
}

func (data *Data_All) Sign(msg []byte, privKey bls.PrivKey) {
	if data.data_sign == nil {
		data.data_sign = NewnulldataS()
	}

	data.data_sign.Sign(msg, privKey)
}
