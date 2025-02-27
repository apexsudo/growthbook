package growthbook

import (
	"context"
	"time"

	"github.com/apexsudo/analytics"
	"github.com/growthbook/growthbook-golang"
	"github.com/tomarrell/wrapcheck/v2/wrapcheck/testdata/ignore_pkg_errors/src/github.com/pkg/errors"
)

type Attributes struct {
	UserID *string
}

type Builder interface {
	GetExperimentClient(attributes Attributes) ExperimentClient
}
type client struct {
	client *growthbook.Client
}

type ExperimentClient interface {
	EvalFeature(context context.Context, key string) *growthbook.FeatureResult
}

func (c *client) GetExperimentClient(attributes Attributes) ExperimentClient {
	expClient, err := c.client.WithAttributes(growthbook.Attributes{
		"id": *attributes.UserID,
	})
	if err != nil {
		panic(errors.Wrap(err, "failed setting attributes, wrong config for experiment?"))
	}
	expClient, err = expClient.WithExtraData(attributes)
	if err != nil {
		panic(errors.Wrap(err, "failed setting extra data, wrong config for experiment?"))
	}

	return expClient
}

func New(clientKey string, analyticsClient analytics.Client) (Builder, error) {
	gbClient, err := growthbook.NewClient(
		context.Background(),
		growthbook.WithClientKey(clientKey),
		growthbook.WithPollDataSource(time.Second*30),
		growthbook.WithExperimentCallback(
			func(
				ctx context.Context,
				experiment *growthbook.Experiment,
				result *growthbook.ExperimentResult,
				extraData any,
			) {
				attrs, _ := extraData.(Attributes)
				_ = analyticsClient.
					Track(*attrs.UserID,
						"Experiment Viewed",
						map[string]any{
							"experimentId": experiment.Key,
							"variationId":  result.VariationId,
						},
					)
			}),
	)
	if err != nil {
		return nil, err
	}
	err = gbClient.EnsureLoaded(context.Background())
	if err != nil {
		return nil, err
	}

	return &client{client: gbClient}, nil
}
