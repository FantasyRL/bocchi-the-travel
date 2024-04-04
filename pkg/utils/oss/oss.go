package aliyunoss

import (
	"bocchi/config"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func OSSBucketCreate() (*oss.Bucket, error) {
	endpoint := config.OSS.EndPoint
	accessKeyId := config.OSS.AccessKeyId
	accessKeySecret := config.OSS.AccessKeySecret
	bucketName := config.OSS.BucketName
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		return nil, err
	}
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return nil, err
	}
	return bucket, nil
}
