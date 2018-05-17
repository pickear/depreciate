package model

type Ticket struct {

	Clstm string  //出发钟点
	Seats string  //票数
}

type Shifts struct {
	Pnm string  //首发站->尾站
	Rdc string   //经过的站点
	tickets map[string][]Ticket   //某一天的所有时间点的票
}