//go:generate go tool mockgen -destination=./mocks/mock_analytics.go -package=mocks -typed github.com/apexsudo/analytics Client
package growthbook_test

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/apexsudo/growthbook"
	"github.com/apexsudo/growthbook/mocks"
)

const mockFeatures = `
{
  "status": 200,
  "features": {
    "test-feature": {
      "defaultValue": false,
      "rules": [
        {
          "id": "fr_19g61wm7ndjl9i",
          "coverage": 1,
          "hashAttribute": "id",
          "seed": "20c124c0-bd76-4e07-bb8f-dcf5aad25643",
          "hashVersion": 2,
          "variations": [
            false,
            true
          ],
          "weights": [
            0.1,
            0.9
          ],
          "key": "test-feature",
          "meta": [
            {
              "key": "0",
              "name": "Control"
            },
            {
              "key": "1",
              "name": "Variation 1"
            }
          ],
          "phase": "1",
          "name": "test-experiment"
        }
      ]
    }
  },
  "dateUpdated": "2025-02-27T13:33:57.134Z"
}
`

func TestNew(t *testing.T) {
	t.Parallel()
	mockAnalyticsClient := mocks.NewMockClient(gomock.NewController(t))
	mockAnalyticsClient.
		EXPECT().
		Track(
			"user_2",
			"Experiment Viewed",
			gomock.Any(),
		).
		DoAndReturn(func(_ string, _ string, m map[string]any) error {
			assert.Equal(t, "test-feature", m["experimentId"])
			assert.Equal(t, 1, m["variationId"])

			return nil
		})
	defer mockGrowthBookEndpoints()()
	client, err := growthbook.New("mock_key", mockAnalyticsClient)
	assert.NoError(t, err)
	userID := "user_2"
	expClient := client.GetExperimentClient(growthbook.Attributes{UserID: &userID})
	assert.NoError(t, err)
	assert.NotNil(t, client)
	assert.True(t, expClient.EvalFeature(t.Context(), "test-feature").On)
}

func mockGrowthBookEndpoints() func() {
	httpmock.Activate()
	httpmock.RegisterResponder(
		"GET",
		"https://cdn.growthbook.io/api/features/mock_key",
		httpmock.NewStringResponder(200, mockFeatures),
	)

	return func() {
		httpmock.DeactivateAndReset()
	}
}
