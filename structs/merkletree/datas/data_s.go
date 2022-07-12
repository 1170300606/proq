package datas

import (
	bls2 "ProQueries/crypto/bls"
	"crypto/sha1"
	"io"
)

type Data_S struct {
	data_content []byte
}

func (data Data_S) Showsign() []byte {
	return data.data_content
}

func (data *Data_S) Sign(msg []byte, privKey bls2.PrivKey) error {
	sig, err := privKey.Sign(msg)

	if err == nil {
		data.data_content = sig
		return nil
	}
	return err
}

func NewnulldataS() *Data_S {
	data_r := &Data_S{
		data_content: []byte(""),
	}
	return data_r
}

func NewdataS(sign []byte) *Data_S {
	data_r := &Data_S{
		data_content: sign,
	}
	return data_r
}

func (data *Data_S) Hashsign(msg []byte) []byte {
	h := sha1.New()
	io.WriteString(h, string(msg))
	data.data_content = h.Sum(nil)
	return data.data_content
}
