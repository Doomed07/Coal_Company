package miners

type MinersInfo struct {
	Junior JuniorMiner
	Middle MiddleMiner
	Senior SeniorMiner
}

var DefaultMinersInfo = MinersInfo{
	Junior: *NewJuniorMiner(),
	Middle: *NewMiddleMiner(),
	Senior: *NewSeniorMiner(),
}
