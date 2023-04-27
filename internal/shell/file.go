package shell

import (
	"context"
	"fmt"
	"path/filepath"
)



func ChangeDir(ctx context.Context, from string, to string) error {
	cmd := fmt.Sprintf("mv %s %s", from, to)
	return run(ctx, cmd)
}


func CreateFile(ctx context.Context, fileName string, filePath string) error {
	newFile := filepath.Join(filePath, fileName)
	cmd := fmt.Sprintf("touch %s", newFile)
	return run(ctx, cmd)
}


func CopyDir(ctx context.Context, from string, to string) error {
	cmd := fmt.Sprintf("cp -r %s %s", from, to)
	return run(ctx, cmd)
}


func DeleteDir(ctx context.Context, file string) error {
	cmd := fmt.Sprintf("rn -rf %s", file)
	return run(ctx, cmd)
}