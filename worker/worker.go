package worker

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func Start() error {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		return err
	}

	defer c.Close()

	w := worker.New(c, "CYBER_TASK_QUEUE", worker.Options{})
	// w.RegisterWorkflow(workflows.)

	log.Println("Starting worker...")
	return w.Run(worker.InterruptCh())
}
