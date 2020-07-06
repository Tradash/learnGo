package main

import (
	"flag"
	"os"
)

var fSource, fDest string
var fOffset, fLimit int

func init() {
	flag.StringVar(&fSource, "from", "", "file to read from")
	flag.StringVar(&fDest, "to", "", "file to read to")
	flag.IntVar(&fOffset, "offset", 0, "offset in input file")
	flag.IntVar(&fLimit, "limit", 0, "limit in input file")
}

func copyFile(fSource, fDest string, fOffset, fLimit int) (string, error) {
	var err error
	var file *os.File
	var src, dst int
	// var fi FileInfo
	if file, err = os.Open(fSource); err != nil {
		return "Ошибка при открытии исходного файла", err
	}
	if fi, err := file.Stat(); err != nil {
		return "Ошибка при определении размера файла", err
	} else {
		if fi.Size() < int64(fOffset) {
			return "Указано смещение боольше размера файла", err
		}
		if fi.Size() < int64(fOffset+fLimit) {
			fLimit = int(fi.Size() - int64(fOffset))
		}
	}

	b := make([]byte, fLimit)
	if src, err = file.ReadAt(b, int64(fOffset)); err != nil {
		return "Ошибка при чтении из файла", err
	}
	println("Прочитано байт", src)
	if file, err = os.Create(fDest); err != nil {
		return "Ошибка при создании нового файла", err
	}
	if dst, err = file.Write(b); err != nil {
		return "Ошибка при записи в файл", err
	}

	println("Записано байт", dst)
	err = file.Close()
	return "Выполнено", err
}

func main() {
	flag.Parse()
	println(fSource, fDest, fOffset, fLimit)
	info, err := copyFile(fSource, fDest, fOffset, fLimit)
	if err != nil {
		println("Panic!!!!", info, err)
	} else {
		println(info)
	}

}
