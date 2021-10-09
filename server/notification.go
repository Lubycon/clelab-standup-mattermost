package main

func SendNotification(p *Plugin) error {
	message := "안냐쎄여 저는 클랩의 스탠드업을 책임져줄 로봇이에요!"
	p.PostBotDM("ziba5knnofy9ucguaotfzwyz3h", message) //FIXME: 여기 바꿔야됨!

	return nil
}
