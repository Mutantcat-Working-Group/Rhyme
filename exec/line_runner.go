package exec

import (
	"context"
	"fmt"
	"os/exec"
	"runtime"
)

func Run(ctx context.Context, cmd *exec.Cmd, line string, su bool) *exec.Cmd {
	if runtime.GOOS == "windows" {
		// Windows系统执行cmd指令
		cmd = exec.CommandContext(ctx, "cmd", "/C", line)
	} else if runtime.GOOS == "linux" {
		// Linux系统执行bash指令
		if !su {
			cmd = exec.CommandContext(ctx, "bash", "-c", line)
		} else {
			cmd = exec.CommandContext(ctx, "sudo", "bash", "-c", line)
		}
	} else {
		fmt.Println("诗 - 未知的操作系统，请反馈此问题。")
		return nil
	}

	return cmd
}
