package cmd

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"syscall"
	"time"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"
)

// 初始化：在init中添加子命令（程序启动时自动执行）
func init() {
	// 注册子命令：start/version/server（修复：serverCmd改为返回*gcmd.Command）
	// Main.AddCommand(stopCmd())
}

func stopCmd() *gcmd.Command {
	return &gcmd.Command{
		Name:  "stop",
		Brief: "停止HTTP服务器",
		Usage: "main stop",
		Func: func(ctx context.Context, parser *gcmd.Parser) error {
			// 1. 定义PID文件路径（和serverCmd保持一致）
			pidFile := "/tmp/main.pid"
			if !gfile.Exists(pidFile) {
				return fmt.Errorf("未找到运行的服务（PID文件不存在）")
			}

			// 2. 读取PID并转换为整数
			pidStr := gfile.GetContents(pidFile)
			pid, err := strconv.Atoi(pidStr)
			if err != nil {
				return fmt.Errorf("解析PID失败：%v", err)
			}

			// 3. 查找进程并发送停止信号（SIGTERM）
			process, err := os.FindProcess(pid)
			if err != nil {
				// 进程不存在，删除无效PID文件
				gfile.Remove(pidFile)
				return fmt.Errorf("服务进程不存在（PID：%d）", pid)
			}

			// 发送SIGTERM信号（触发serverCmd中的信号监听）
			if err := process.Signal(syscall.SIGTERM); err != nil {
				return fmt.Errorf("发送停止信号失败：%v", err)
			}

			// 4. 验证进程是否停止（可选）
			fmt.Printf("已向服务进程（PID：%d）发送停止信号，正在验证...\n", pid)
			time.Sleep(time.Second * 1) // 等待1秒
			// 再次尝试发送信号，判断进程是否存活
			if err := process.Signal(syscall.Signal(0)); err == nil {
				return fmt.Errorf("服务进程（PID：%d）未停止，可能需要强制终止", pid)
			}

			// 5. 删除PID文件
			gfile.Remove(pidFile)
			fmt.Printf("服务进程（PID：%d）已成功停止\n", pid)
			return nil
		},
	}
}
