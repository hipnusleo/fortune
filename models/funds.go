package models

/*
fund_tb 记录基金基本信息
fundNetValue_tb 记录着基金公司每天的净值数据


*/

/*
Fund 表记录着基金的基本情况
*/
type Fund struct {
	Name             string
	Code             string
	Equity           int // unit = ¥
	EstablishedAt    int32
	Manager          string
	ManagerChangedAt int32
	Keyword          string
	NetValueRecordID int
}

/*
FundNetValueRecord 表记录着基金的净值变化
*/
type FundNetValueRecord struct {
	Code             string
	NetValue         float32
	NetValueRecordID int
	RecordedAt       int32
}
