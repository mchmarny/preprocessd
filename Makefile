mod:
	go mod tidy
	go mod vendor

image: mod
	gcloud builds submit \
		--project cloudylabs-public \
		--tag gcr.io/cloudylabs-public/preprocessd:0.1.3

service:
	gcloud beta run deploy preprocessd \
		--image=gcr.io/cloudylabs-public/preprocessd:0.1.3 \
		--region=us-central1 \
		--concurrency=80 \
		--memory=256Mi

serviceless:
	gcloud beta run services delete preprocessd

