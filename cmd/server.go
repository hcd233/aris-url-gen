package cmd

import (
	"fmt"

	"github.com/gofiber/contrib/fiberzap"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/hcd233/Aris-url-gen/internal/api/router"
	"github.com/hcd233/Aris-url-gen/internal/config"
	"github.com/hcd233/Aris-url-gen/internal/cron"
	"github.com/hcd233/Aris-url-gen/internal/logger"
	"github.com/hcd233/Aris-url-gen/internal/resource/cache"
	"github.com/hcd233/Aris-url-gen/internal/resource/database"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "服务器命令组",
	Long:  `包含服务器相关操作的命令组`,
}

var startServerCmd = &cobra.Command{
	Use:   "start",
	Short: "启动API服务器",
	Long:  `启动并运行API服务器并监听指定的主机和端口`,
	Run: func(cmd *cobra.Command, args []string) {
		host, port := lo.Must1(cmd.Flags().GetString("host")), lo.Must1(cmd.Flags().GetString("port"))

		database.InitDatabase()
		cache.InitCache()
		cron.InitCronJobs()

		app := fiber.New(fiber.Config{
			AppName:           "Aris-url-gen",
			ReadTimeout:       config.ReadTimeout,
			WriteTimeout:      config.WriteTimeout,
			Concurrency:       config.Concurrency,
			EnablePrintRoutes: config.APIMode == config.ModeDev,
		})

		app.Use(
			cors.New(),
			recover.New(recover.Config{
				EnableStackTrace: true,
			}),
			fiberzap.New(fiberzap.Config{
				Logger: logger.Logger,
			}),
		)

		router.RegisterRouter(app)

		lo.Must0(app.Listen(fmt.Sprintf("%s:%s", host, port)))
	},
}

func init() {
	serverCmd.AddCommand(startServerCmd)
	rootCmd.AddCommand(serverCmd)

	startServerCmd.Flags().StringP("host", "", "localhost", "监听的主机")
	startServerCmd.Flags().StringP("port", "p", "8080", "监听的端口")
}
