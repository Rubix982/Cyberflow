package workflows

import (
	"github.com/Rubix982/cyberflow/models"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
	"time"
)

func IntelWorkflow(ctx workflow.Context) error {
	retryPolicy := temporal.RetryPolicy{
		InitialInterval:    time.Second * 5,
		BackoffCoefficient: 2.0,
		MaximumInterval:    time.Minute,
		MaximumAttempts:    5,
	}

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 30,
		RetryPolicy:         &retryPolicy,
	}

	ctx = workflow.WithActivityOptions(ctx, ao)

	var result models.IntelData

	if err := workflow.ExecuteActivity(ctx, "FetchShodan").Get(ctx, &result.Shodan); err != nil {
		return err
	}

	if err := workflow.ExecuteActivity(ctx, "FetchVirusTotal").Get(ctx, &result.VirusTotal); err != nil {
		return err
	}

	if err := workflow.ExecuteActivity(ctx, "FetchAbuseIPDB").Get(ctx, &result.AbuseIPDB); err != nil {
		return err
	}

	return nil
}
