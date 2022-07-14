package datas

type Data_R struct {
	//data_content block.Account
	data_content Content `json:"dataR"` //值
}

func NewnullDataR() *Data_R {
	data_r := &Data_R{}
	return data_r
}

//func NewDataR(Account account.AccountKey, value int, pointer account.Pointer) *Data_R {
//
//	data_r := &Data_R{
//		data_content: *account.NewAcconunt(Account, value, pointer),
//	}
//	return data_r
//}

func NewDataR(text Content) *Data_R {

	data_r := &Data_R{
		data_content: text,
	}
	return data_r
}

func (data *Data_R) Showdata() Content {
	return data.data_content
}

func (data *Data_R) Copy() Data_R {
	data_r := Data_R{
		data_content: data.data_content,
	}
	return data_r
}

func (data *Data_R) Islower(data_2 Data_R) bool {
	//if data.data_content <= data_2.data_content {
	//	return true
	//}
	//return false
	return data.data_content.Islower(data_2.data_content)
}

func (data *Data_R) Ishigher(data_2 Data_R) bool {
	//if data.data_content >= data_2.data_content {
	//	return true
	//}
	//return false

	return data.data_content.Ishigher(data_2.data_content)
}

func (data *Data_R) Equals(data_2 Data_R) bool {
	//if data.data_content == data_2.data_content {
	//	return true
	//}
	//return false

	return data.data_content.Equals(data_2.data_content)
}

func (data *Data_R) Tobyte() []byte {
	return data.data_content.Tobyte()
}
