package datas

import (
	"ProQueries/crypto/bls"
	"bytes"
)

type Data_All struct {
	data_content *Data_R `json:"dataR"` //值
	data_sign    *Data_S `json:"dataS"` //加密
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

func (data *Data_All) Show_sign() []byte {

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

func (data *Data_All) Signs() []byte {
	if data.data_sign == nil {
		data.data_sign = NewnulldataS()
		data.data_sign.Hashsign(data.data_content.Tobyte())
	} else {
		data.data_sign.Hashsign(data.data_content.Tobyte())
	}
	return data.data_sign.Showsign()
}

func (data *Data_All) Tobyte() []byte {
	var buffer bytes.Buffer //Buffer是一个实现了读写方法的可变大小的字节缓冲

	buffer.Write(data.data_content.Tobyte())

	buffer.Write(data.data_sign.Tobyte())

	b3 := buffer.Bytes() //得到了b1+b2的结果
	return b3
}

func (data *Data_All) ToStr() string {
	return string(data.Tobyte())
}

func (data *Data_All) Equals(data2 *Data_All) bool {
	a := data.data_content == data2.data_content
	b := data.data_sign == data2.data_sign
	return (a && b)
}
