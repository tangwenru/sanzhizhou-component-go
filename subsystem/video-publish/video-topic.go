package subsystemVideoPublish

import (
	"fmt"
	sanzhizhouComponentLib "github.com/tangwenru/sanzhizhou-component-go/lib"
	typeVideoTopic "github.com/tangwenru/sanzhizhou-component-go/subsystem/video-publish/type"
)

type VideoTopic struct {
}

func init() {

}

func (this *VideoTopic) TopicList(
	userId int64,
	query *typeVideoTopic.VideoTopicTopicListQuery,
) (error, *[]typeVideoTopic.VideoTopicTopicList) {
	result := typeVideoTopic.VideoTopicListResult{}

	bytesResult, err := sanzhizhouComponentLib.SubsystemVideoPublish(
		userId,
		"system/videoTopic/list",
		query,
		&result,
	)

	if err != nil {
		fmt.Println("VideoTopic list :", string(bytesResult))
		return err, nil
	}

	return nil, &result.Data
}
