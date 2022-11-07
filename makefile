REPO_NAME=asia.gcr.io/vds-client-amc
IMG_NAME=notifier
IMG_TAG=latest
BUCKET_NAME=vds-client-amc-notifiers-config
PROJECT_ID=vds-client-amc

all: build push apply deploy

build:
	@echo "Building docker image..."
	@docker build . -t ${REPO_NAME}/${IMG_NAME}:${IMG_TAG} --platform linux/amd64 -f ./slack/Dockerfile

push:
	@echo "Pushing docker image..."
	@docker push ${REPO_NAME}/${IMG_NAME}:${IMG_TAG}

apply:
	@echo "Applying Configfile to Cloud Storage..."
	@gsutil cp ./slack/slack.yaml gs://${BUCKET_NAME}/slack.yaml

deploy:
	@echo "Deploying to Cloud Run..."
	@gcloud run deploy slack-notifier \
		--image=${REPO_NAME}/${IMG_NAME}:${IMG_TAG} \
		--no-allow-unauthenticated \
		--max-instances=1 \
		--update-env-vars=CONFIG_PATH=gs://${BUCKET_NAME}/slack.yaml,PROJECT_ID=${PROJECT_ID}