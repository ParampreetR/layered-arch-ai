package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	common_config "github.com/parampreetr/layered_arch_ai/internal/common/config"
	common_di "github.com/parampreetr/layered_arch_ai/internal/common/di"
	erm_transfer_bank_di "github.com/parampreetr/layered_arch_ai/internal/erm-transfer-bank/di"
)

func InitModules(router fiber.Router, env *common_config.Config) {
	common_di.Init(env, router)
	erm_transfer_bank_di.Init(router)
}

func main() {
	env := common_config.Load()

	app := fiber.New()
	router := app.Group("/api/v1")

	InitModules(router, env)

	if err := app.Listen(fmt.Sprintf("0.0.0.0:%s", env.Port)); err != nil {
		fmt.Println(err)
	}
}
