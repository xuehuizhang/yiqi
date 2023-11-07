package viewModel

type Channel struct {
	Id     int64          `json:"id"`
	Name   string         `json:"name"`
	Childs []ChannelChild `json:"childs"`
}

type ChannelChild struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func ChannelMockData() []Channel {
	list := make([]Channel, 0)

	for i := 0; i < 2; i++ {
		c := Channel{
			Id:     int64(i + 1),
			Name:   "测试主标题",
			Childs: nil,
		}
		for j := 0; j < 5; j++ {
			c.Childs = append(c.Childs, ChannelChild{
				Id:   int64(j),
				Name: "测试副标题",
			})
		}
		list = append(list, c)
	}

	return list
}
