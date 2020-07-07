package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

type EnvData map[string]string

func ReadDir(path string) (EnvData, error) {
	envData := make(EnvData)
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	fileInfo, err := f.Readdir(-1)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	for _, file := range fileInfo {
		// println(file.Name(), filepath.Join(path,"/",file.Name()))
		fileTmp, err := os.Open(filepath.Join(path, "/", file.Name()))
		if err != nil {
			return nil, err
		}
		data, err := ioutil.ReadAll(fileTmp)
		if err != nil {
			return nil, err
		}
		envData[file.Name()] = string(data)

	}
	return envData, nil

}

func RunCmd(cmd []string, env map[string]string) int {
	count := 0
	for k, v := range env {
		if err := os.Setenv(k, v); err != nil {
			return 0
		}
	}
	for i := 0; i < len(cmd); i++ {
		cmd := exec.Command(cmd[i])
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		log.Printf("\nCommand finished with error: %v", err)
		if err == nil {
			count++
		}
	}

	return count
}

func main() {
	// var args []string
	defer println("Завершена работа приложения...")
	args := os.Args
	println("Running......")
	for i := 0; i < len(args); i++ {
		println("Argument", i, "--", args[i])
	}
	if len(args) < 3 {
		println("Неправильное количество аргументов...")
		return
	}

	dirName, err := filepath.Abs(args[1])
	if err != nil {
		println("Указан неправильный путь")
		return
	}
	// println(s)
	envs, err := ReadDir(dirName)
	if err != nil {
		println("Ошибка при чтении данных из директории", err)
	}
	println("Произведена загрузка данных о переменных окружения")
	// fmt.Printf("%v", envs)

	var scriptGo []string

	for i := 2; i < len(args); i++ {
		scriptGo = append(scriptGo, args[i])
	}
	fmt.Printf("Перечень скриптов для запуска - %v,\nПеременные  - %v\n", scriptGo, envs)
	count := RunCmd(scriptGo, envs)
	println("Выполнен запуск ", count, "скриптов")

}
