package test_pkg_1

type Student struct {
	Age       int
	Name      string
	ScoreList []*Score
}

type Score struct {
	Name string
	Num  int
}
