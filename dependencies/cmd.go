package dependencies

import (
	"fmt"
	"os/exec"
)

func GoModInit(folderName string) error {
	// 4. Inisialisasi modul Go
	cmd := exec.Command("go", "mod", "init", folderName)
	cmd.Dir = folderName
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("go mod init failed: %s\n%s", err, output)
	}
	return nil
}

func GoModTidy(folderName string) error {
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = folderName
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("go mod tidy failed: %s\n%s", err, output)
	}
	return nil
}

func OpenInVSCode(folderName string) error {
	cmd := exec.Command("code", folderName)
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start VSCode: %w", err)
	}
	return nil
}
