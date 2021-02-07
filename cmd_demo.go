package main

import (
	// "bytes"
	"log"
	"os"
	"os/exec"
)

func main() {
	// 无返回结果
	// cmd := exec.Command("ls","-l", "/var/log")
	// err := cmd.Run()
	// if err != nil {
	// 	log.Fatalf("cmd.Run() failed with %s\n", err)
	// }

	// 有返回结果
	// cmd := exec.Command("ls", "-l", "/var/log")
	// out, err := cmd.CombinedOutput()
	// if err != nil {
	// 	log.Printf("combined out: \n %s \n", string(out))
	// 	log.Fatalf("cmd.Run()failed with %s \n", err)
	// }
	// log.Printf("combined out: \n %s \n", string(out))

	// 区分stdout,stderr
	// cmd := exec.Command("ls","-l", "/var/log/*.log")
	// var stdout, stderr bytes.Buffer
	// cmd.Stdout = &stdout
	// cmd.Stderr = &stderr
	// err:= cmd.Run()
	// outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())

	// log.Printf("out: \n %s \n err:\n%s\n",  outStr, errStr)
	// if err != nil {
	// 	log.Fatalf("cmd.Run() failed with %s\n", err)
	// }

	// 多命令组合,请使用管道
	c1 := exec.Command("grep", "ERROR", "/var/log/messages")
	c2 := exec.Command("wc", "-l")
	c2.Stdin, _ = c1.StdoutPipe()
	c2.Stdout = os.Stdout
	_ = c2.Start()
	_ = c1.Run()
	_ = c2.Wait()

	// 设置命令级别的环境变量
	os.Setenv("NAME", "wangbm")
	cmd := exec.Command("echo", os.ExpandEnv("$NAME"))
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	log.Printf("%s", out)

	cmd1 := exec.Command("bash", "/home/wangbm/demo.sh")
	ChangeYourCmdEnvironment(cmd1)

	out1, _ := cmd1.CombinedOutput()
	log.Printf("output: %s",  out1)

	cmd2 := exec.Command("bash", "/home/wangbm/demo.sh")
	out2, _ := cmd2.CombinedOutput()
	log.Printf("output: %s", out2)
}

func ChangeYourCmdEnvironment(cmd *exec.Cmd) error {
	env := os.Environ()
	cmdEnv := []string{}

	for _, e := range env {
		cmdEnv = append(cmdEnv, e)
	}

	cmdEnv = append(cmdEnv, "NAME=wangbm")
	cmd.Env = cmdEnv
	return nil
}
