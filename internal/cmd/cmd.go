package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"

	"test/internal/controller/hello"
	"test/internal/controller/users"
	"test/internal/middleware"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
)

// 基础认证中间件：OpenAPI接口的前置认证
func openApiBasicAuth(r *ghttp.Request) {
	if !r.BasicAuth("OpenApiAuthUserName", "OpenApiAuthPass", "Restricted") {
		r.ExitAll()
		return
	}
}

// 定义主命令（对外暴露）
var Main = &gcmd.Command{
	Name:        "main",
	Usage:       "main [command] [options]",
	Brief:       "start http server",
	Description: "这是基于GF框架的自定义命令行程序，支持启动服务、查看版本、启动HTTP服务器",
}

// 初始化：在init中添加子命令（程序启动时自动执行）
func init() {
	Main.AddCommand(startCmd(), versionCmd())
}

// 定义version子命令（原逻辑不变）
func versionCmd() *gcmd.Command {
	return &gcmd.Command{
		Name:  "version",
		Brief: "查看版本",
		Usage: "main version",
		Func: func(ctx context.Context, parser *gcmd.Parser) error {
			fmt.Println("v1.0.0")
			return nil
		},
	}
}

func startCmd() *gcmd.Command {
	return &gcmd.Command{
		Name:  "start",
		Brief: "启动HTTP服务器（包含SSE、用户REST接口、OpenAPI认证）",
		Usage: "main start [--port=8080]",
		Func: func(ctx context.Context, parser *gcmd.Parser) error {
			// 修复3：只定义一次isDev，避免覆盖
			isDev := gcmd.GetOptWithEnv("dev", "false").Bool()
			pidFile := "/tmp/main.pid"

			if !isDev && gfile.Exists(pidFile) {
				oldPid := gfile.GetContents(pidFile)
				return fmt.Errorf("服务已运行，PID：%s（请先执行 main stop 停止）", oldPid)
			}

			if !isDev {
				pid := os.Getpid()
				if err := gfile.PutContents(pidFile, fmt.Sprintf("%d", pid)); err != nil {
					return fmt.Errorf("写入PID文件失败：%v", err)
				}
				defer gfile.Remove(pidFile)
			}

			port := parser.GetOpt("port", 8080).Int()
			fmt.Printf("HTTP服务器启动中，端口：%d，PID：%d\n", port, os.Getpid())

			s := g.Server()
			s.SetPort(port)
			s.SetGraceful(true)

			// 修复2：SSE接口从根路径 `/` 改为 `/sse`，避免路由冲突
			s.BindHandler("/sse", func(r *ghttp.Request) {
				r.Response.Header().Set("Content-Type", "text/event-stream")
				r.Response.Header().Set("Cache-Control", "no-cache")
				r.Response.Header().Set("Connection", "keep-alive")

				for i := 0; i < 100; i++ {
					r.Response.Writefln("data: %d", i)
					r.Response.Flush()
					time.Sleep(time.Millisecond * 200)
				}
			})

			// 关键：全局跨域中间件（替代仅/user分组的中间件，确保所有接口都支持跨域）
			s.Use(middleware.MiddlewareCORS)

			// /user 分组：无需重复绑定跨域中间件
			s.Group("/user", func(group *ghttp.RouterGroup) {
				group.Bind(new(users.User))
			})

			s.Group("/hello", func(group *ghttp.RouterGroup) {
				group.Bind(new(hello.ControllerV1))
			})

			// 修复3：使用已定义的isDev，避免重复定义
			if !isDev {
				s.BindHookHandler(s.GetOpenApiPath(), ghttp.HookBeforeServe, openApiBasicAuth)
			}

			// 监听系统信号
			sigChan := make(chan os.Signal, 1)
			signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
			go func() {
				sig := <-sigChan
				fmt.Printf("\n收到信号：%v，正在优雅关闭服务器...\n", sig)
				if err := s.Shutdown(); err != nil {
					fmt.Printf("服务器关闭失败：%v\n", err)
				} else {
					fmt.Println("服务器已成功停止")
				}
				os.Exit(0)
			}()

			s.Run()
			return nil
		},
	}
}
