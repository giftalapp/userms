package src

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

func InitSNS() (*sns.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		return nil, err
	}

	sc := sns.NewFromConfig(cfg)

	topicPages := sns.NewListTopicsPaginator(sc, &sns.ListTopicsInput{})

	for topicPages.HasMorePages() {
		topicPage, err := topicPages.NextPage(context.TODO())

		if err != nil {
			return nil, err
		}

		for _, topic := range topicPage.Topics {
			fmt.Printf("[INFO] (SNS) Available Topic: %s\n", *topic.TopicArn)
		}
	}

	return sc, nil
}
