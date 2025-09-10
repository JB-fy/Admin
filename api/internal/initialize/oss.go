package initialize

import (
	"context"
	// "github.com/aws/aws-sdk-go-v2/aws"
	// "github.com/aws/aws-sdk-go-v2/config"
	// "github.com/aws/aws-sdk-go-v2/credentials"
	// "github.com/aws/aws-sdk-go-v2/service/sts"
)

// 初始化可能有顺序要求，故统一到这里执行初始化函数
func initOss(ctx context.Context) {
	/* accessKeyID := `DRvjoX3qIsFzba9dypZc`
	accessKeySecret := `WRL5tq0KNd17Gx2mEYCQFO8vgSyz3sXbkaMPIBnV`
	endpoint := `http://192.168.0.200:9000`
	region := `us-west-2`
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyID, accessKeySecret, ``)),
		config.WithBaseEndpoint(endpoint),
		config.WithRegion(region),
	)
	if err != nil {
		gutil.Dump(err)
	}

	stsClient := sts.NewFromConfig(cfg)
	respSts, err := stsClient.AssumeRole(ctx, &sts.AssumeRoleInput{
		RoleArn:         aws.String(`arn:aws:iam::123456789012:role/dummy`),
		RoleSessionName: aws.String(`web-session`),
		DurationSeconds: aws.Int32(3600),
		Policy: aws.String(`{
	    "Version": "2012-10-17",
	    "Statement": [
	        {
	            "Effect": "Allow",
	            "Action": [
	                "s3:*"
	            ],
	            "Resource": [
	                "arn:aws:s3:::*"
	            ]
	        }
	    ]
	}`),
	})
	if err != nil {
		gutil.Dump(err.Error())
		return
	}
	gutil.Dump(respSts.Credentials)

	client := s3.NewFromConfig(cfg, func(o *s3.Options) { o.UsePathStyle = true })
	// 生成预签名直传地址。只能固定到key对应的位置，也就是只能传一个文件
	resp, err := s3.NewPresignClient(client).PresignPutObject(ctx,
		&s3.PutObjectInput{
			Bucket: aws.String(`test`),
			Key:    aws.String(`upload/` + time.Now().Format(`20060102-150405`) + `.txt`),
			// ContentType: aws.String(`image/jpeg`),
		},
		s3.WithPresignExpires(10*time.Minute),
	)
	if err != nil {
		gutil.Dump(err.Error())
		return
	}
	gutil.Dump(resp)

	// 查看桶
	resp, err := client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		gutil.Dump(err.Error())
		return
	}
	for _, b := range resp.Buckets {
		gutil.Dump(b.Name)
	}

	bucket := `test`
	_, err = client.CreateBucket(ctx, &s3.CreateBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		gutil.Dump(err.Error())
		return
	}

	_, err = client.DeleteBucket(ctx, &s3.DeleteBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		gutil.Dump(err.Error())
		return
	}

	// 查看对象
	respObj, err := client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		gutil.Dump(err.Error())
		return
	}
	for _, obj := range respObj.Contents {
		gutil.Dump(obj.Key)
	}

	objKey := `test.txt`
	_, err = client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(objKey),
		Body:   strings.NewReader(`测试`),
	})
	if err != nil {
		gutil.Dump(err.Error())
		return
	}

	_, err = client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(objKey),
	})
	if err != nil {
		gutil.Dump(err.Error())
		return
	}

	respObjGet, err := client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(objKey),
	})
	if err != nil {
		gutil.Dump(err.Error())
		return
	}
	defer respObjGet.Body.Close()
	data, err := io.ReadAll(respObjGet.Body)
	if err != nil {
		gutil.Dump(err.Error())
		return
	}
	gutil.Dump(data) */
}
