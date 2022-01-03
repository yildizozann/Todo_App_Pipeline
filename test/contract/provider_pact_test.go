// +build pact_test

package contract_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"todo_app/config"
	"todo_app/internal/todo"

	"github.com/joho/godotenv"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"github.com/stretchr/testify/assert"
)

func RunApp() error{
	

	
	app := todo.CreateApp()
	database := todo.NewDatabase()
	service := todo.NewService(database)
	handler := todo.NewHandler(service)
	
	handler.RegisterRoutes(app)

	return app.Listen(":3000")
}


func TestProvider(t *testing.T) {

	
	errLoad := godotenv.Load("../../.env")
	if errLoad != nil {
		log.Fatal("Error loading .env file")
	}
	
	config.ConnectDB()
	go RunApp()
	port := "3000"

	var consumerSelector = types.ConsumerVersionSelector{
		Tag: "master",
		Latest: true,
	}
	var stateHandlers = types.StateHandlers{
		"todo's exist": func() error {
			return nil
		},
		"todo's created": func() error {
			return nil
		},
	}
	pact := dsl.Pact{
		Provider: "todo-provider",
		Consumer: "todo-consumer",
		Host: "127.0.0.1",
		DisableToolValidityCheck: true,
	}

	_, err := pact.VerifyProvider(t, types.VerifyRequest{

		ProviderBaseURL:    fmt.Sprintf("http://127.0.0.1:%s", port),
		Tags:               []string{"master"},
		FailIfNoPactsFound: false,
		Verbose: 			false,
		PactURLs:                   []string{os.Getenv("PACT_BROKER_BASE_URL")},
		
		BrokerToken:             	os.Getenv("PACT_BROKER_TOKEN"),
		PublishVerificationResults: true,
		ProviderVersion:            "1.0.0",
		StateHandlers:              stateHandlers,
		ConsumerVersionSelectors: 	[]types.ConsumerVersionSelector{
			consumerSelector,
		},
		
  })
  assert.NoError(t, err)
} 