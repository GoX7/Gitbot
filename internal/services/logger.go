package services

import (
	"fmt"
	"log/slog"
	"os"
	"time"
)

func NewLogger(name string, logger **slog.Logger) {
	os.MkdirAll("/logs/", 0755)
	filename := fmt.Sprintf("/logs/%s-%s.log", name, time.Now().Format(time.DateOnly))
	open, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("[-] logger.file: " + err.Error())
		return
	}

	fmt.Println("[+] logger.new: " + name)
	*logger = slog.New(slog.NewTextHandler(open, nil))
}
