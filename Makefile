mod:
	go mod tidy
	go mod vendor

auth:
	gcloud projects add-iam-policy-binding cloudylabs \
    	--member="serviceAccount:service-727951037194@gcp-sa-pubsub.iam.gserviceaccount.com" \
    	--role=roles/iam.serviceAccountTokenCreator

meta:
	PROJECT=$(gcloud config get-value project)
	PROJECT_NUM=$(gcloud projects list --filter="${PROJECT}" --format="value(PROJECT_NUMBER)")

image: mod
	gcloud builds submit \
		--project cloudylabs-public \
		--tag gcr.io/cloudylabs-public/preprocessd:0.1.5

service:
	gcloud beta run deploy preprocessd \
		--image=gcr.io/cloudylabs-public/preprocessd:0.1.5 \
		--region=us-central1 \
		--concurrency=80 \
		--memory=256Mi
connect:
	gcloud compute ssh eventmaker --zone us-central1-c

serviceless:
	gcloud beta run services delete preprocessd


sa:
	gcloud iam service-accounts create preprocessdinvoker \
    	--display-name "PreProcess Cloud Run Service Invoker"

	gcloud beta run services add-iam-policy-binding preprocessd \
		--member=serviceAccount:preprocessdinvoker@cloudylabs.iam.gserviceaccount.com \
		--role=roles/run.invoker

sub:
	gcloud beta pubsub subscriptions create preprocessdsub \
		--topic eventmaker \
		--push-endpoint=https://preprocessd-2gtouos2pq-uc.a.run.app/ \
		--push-auth-service-account=preprocessdinvoker@cloudylabs.iam.gserviceaccount.com

