package shell_test

import (
	"context"
	"os"
	"os/exec"
	"testing"

	"github.com/mulmuri/grassflow/internal/task/shell"
)


func TestMain(m *testing.M) {

	if err := os.Chdir("../../../test"); err != nil {
		panic(err)
	}

	cmd := exec.Command("git", "init")
	cmd.Run()
	
	code := m.Run()
	os.Exit(code)
}


func TestBackup(t *testing.T) {

	if err := shell.Backup(context.TODO()); err != nil {
		t.Errorf("error: %v", err)
		return
	}

	path := ".git.swp"
	info, err := os.Stat(path);
	
	if os.IsNotExist(err) {
		t.Errorf(".git.swp file doesn't exist on %s", path)
		return
	}

	if !info.IsDir() {
		t.Errorf("expected dir, file exist")
		return
	}
}



func TestRollback(t *testing.T) {

	if err := shell.Rollback(context.TODO()); err != nil {
		t.Errorf("error: %v", err)
		return
	}

	path := ".git.swp"
	_, err := os.Stat(path);
	
	if !os.IsNotExist(err) {
		t.Errorf(".git.swp file is not deleted %s", path)
		return
	}
}
