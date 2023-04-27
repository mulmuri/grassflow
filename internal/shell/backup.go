package shell

import "context"

var (
	origin = ".git"
	backup = ".git.swp"
	temporary = ".git.tmp"
)



func Backup(ctx context.Context) error {
	return CopyDir(ctx, origin, backup)
}


func Rollback(ctx context.Context) error {

	if err := ChangeDir(ctx, origin, temporary); err != nil {
		return err
	}

	if err := ChangeDir(ctx, backup, origin); err != nil {
		return err
	}

	if err := DeleteDir(ctx, temporary); err != nil {
		return err
	}

	return nil
}