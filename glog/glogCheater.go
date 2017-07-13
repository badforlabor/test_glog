package glog

var dir string = ""

func Cheat(tmpdir string) {
	dir = tmpdir
	// 设置输出路径
	logDir = &dir
}
