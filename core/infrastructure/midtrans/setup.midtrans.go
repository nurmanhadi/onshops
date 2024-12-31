package midtrans

import (
	"os"

	"github.com/midtrans/midtrans-go"
)

func SetupMidtrans() {
	midtrans.ServerKey = os.Getenv("SERVER_KEY")
	midtrans.Environment = midtrans.Sandbox
}
