package localPool

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/testcontainers/testcontainers-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type ComposeSetupIntput struct {
	ReusableCompose bool
	ComposePath     string
}

type ComposeSetupOutput struct {
	Compose  *testcontainers.LocalDockerCompose
	Teardown func()
}

func SetupTestContainers(input ComposeSetupIntput) (*ComposeSetupOutput, error) {
	c := testcontainers.NewLocalDockerCompose(
		[]string{input.ComposePath},
		"containers-tobepaid",
	)

	expectedImages := []string{
		"mongotbp",
	}

	running, err := isRunning(expectedImages)
	if err != nil {
		return nil, err
	}

	if !running {
		stopRunningInstances(expectedImages)

		err := composeUp(c)
		if err != nil {
			return nil, err
		}
	}

	return &ComposeSetupOutput{
		Compose: c,
		Teardown: func() {
			if !input.ReusableCompose {
				c.Down()
			}
		},
	}, nil
}

func isRunning(expectedImages []string) (bool, error) {
	stdout, err := exec.Command("docker", "ps").Output()
	if err != nil {
		return false, err
	}

	ps := string(stdout)
	if err != nil {
		return false, err
	}

	running := true
	for _, image := range expectedImages {
		if !strings.Contains(ps, image) {
			running = false
			break
		}
	}
	return running, nil
}

func stopRunningInstances(expectedImages []string) {
	for _, image := range expectedImages {
		exec.Command("docker", "stop", image).Run()
		exec.Command("docker", "rm", image).Run()
	}
}

func composeUp(compose *testcontainers.LocalDockerCompose) error {
	err := compose.WithCommand([]string{"up", "-d"}).Invoke()
	if err.Error != nil {
		return err.Error
	}

	// Didmnot see any need to wait here, but if needed check this out:
	// https://github.com/vishnubob/wait-for-it

	return nil
}

func NewPoolDocDB() (*mongo.Database, error) {
	url := "mongodb://admin:admin@localhost:27017"
	opts := options.Client()

	opts.ApplyURI(url)
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		msg := fmt.Sprintf("invalid mongodb connection: %v", err)
		log.Println(msg)
		return nil, fmt.Errorf(msg)
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Printf("connection test to mongodb client failed: %v", err)
		return nil, fmt.Errorf(fmt.Sprintf("connection test to mongodb client failed: %v", err))
	}

	log.Println("connected successfully to mongodb")

	db := client.Database("tobepaid")

	return db, nil
}

func NewLocalDocDB() (*mongo.Database, error) {
	database, err := NewPoolDocDB()
	if err != nil {
		return nil, err
	}

	return database, nil
}
