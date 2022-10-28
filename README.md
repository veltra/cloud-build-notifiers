# Cloud Build Notifiers customized for vds-amc-client

#### **_[Default Readme](./README-2022-10.md)_**

This repository is forked from [GoogleCloudPlatform/cloud-build-notifiers](https://github.com/GoogleCloudPlatform/cloud-build-notifiers).

This is used for [vds-amc-client](https://github.com/veltra/vds-amc-client) as sending slack notification after deployment by Cloud Build.

## Background of why needed to fork and customize
For following reasons, needed to make own image for Cloud Run.

- The newest [GoogleCloudPlatform/cloud-build-notifiers](https://github.com/GoogleCloudPlatform/cloud-build-notifiers) is not enough to use for slack notification.
- The newest slack image `us-east1-docker.pkg.dev/gcb-release/cloud-build-notifiers/slack:latest` is not recommended at this moment (2022.10.28). 
- [Google Document](https://cloud.google.com/build/docs/configuring-notifications/configure-slack#:~:text=Notifier%20%E3%82%92%20Cloud%20Run%20%E3%81%AB%E3%83%87%E3%83%97%E3%83%AD%E3%82%A4%E3%81%97%E3%81%BE%E3%81%99%E3%80%82) and [GoogleCloudPlatform/cloud-build-notifiers](https://github.com/GoogleCloudPlatform/cloud-build-notifiers/tree/master/slack#:~:text=For%20release%201.15%20and%20above%3A) recommend to use `us-east1-docker.pkg.dev/gcb-release/cloud-build-notifiers/slack:slack-1.14.0`.
- Although they say 'use `slack:slack-1.14.0`', it is going to fail on Cloud Run deployment.
- Needed to customize for adding notification contents (ex. branch name).

> Work log : https://github.com/veltra/vds-amc-client/issues/1642#issuecomment-1292845303


## Overview of what is custamized form the newest version of [GoogleCloudPlatform/cloud-build-notifiers](https://github.com/GoogleCloudPlatform/cloud-build-notifiers)
### Branch
There are two the principal axis branches.
| Branch | Usage |
----|---- 
| `master` | The newest source code [GoogleCloudPlatform/cloud-build-notifiers](https://github.com/GoogleCloudPlatform/cloud-build-notifiers).<br> Usually we do not update.|
| `amc-master` | Start revision from [slack-1.14.0](https://github.com/GoogleCloudPlatform/cloud-build-notifiers/commit/ac48f4d42d36ffcb81844c521da7a112a5bdc4ed).<br> Make PR to this branch.|

### Customized slack notification contents
- Default contents sended to slack (Before)
    - Project ID
    - Build ID
    - Build Status
- Fixed contents sended to slack (After)
    - Environment (where to deploy)
    - Branch (deployed branch)
    - Deployed Commit (deployed commit)
    - Cluster (deployed cluster)
    - Trriger (used cloud build trigger)


## How to build, push slack image and deploy to Cloud Run

### Preparation
Set `gcloud config`

### Check vds-client-amc configuration is set on your local.
- Execute following command to check it
  ```
  % gcloud config configurations list
  NAME              IS_ACTIVE  ACCOUNT                      PROJECT          COMPUTE_DEFAULT_ZONE  COMPUTE_DEFAULT_REGION
  ana-op            True       norihide.yoshida@veltra.com  vds-client-amc   asia-northeast1-a     asia-northeast1
  .
  .
  ```

- If you already have the config, switch config to vds-client-amc
  ```
  % gcloud config configurations activate ana-op
  ```

- If not set vds-client-amc configuration on local, execute following command.
  ```
  % gcloud config configurations create ana-op       // create configuration
  % gcloud config set compute/region asia-northeast1 // set regison
  % gcloud config set compute/zone asia-northeast1-a // set zone
  % gcloud config set core/project vds-client-amc    // set project id
  ```

### How to build, push slack image and deploy to Cloud Build
- A makefile exists 
  ```
  % cd [this project root]
  % make all
  ```

- When get message on your terminal, it is successfully deployed to Cloud run.
  ```
  Service [slack-notifier] revision [slack-notifier-00070-lug] has been deployed and is serving 100 percent of traffic.
  Service URL: https://slack-notifier-qtswhpgk2a-an.a.run.app
  ```



## Used GCP services
- [Cloud Build > anaop-build-and-deploy](https://console.cloud.google.com/cloud-build/triggers;region=global/edit/e66f2634-c28a-4d10-9e75-162eb4bfc5d1?project=vds-client-amc)
- [GCR > asia.gcr.io/vds-client-amc/notifier](https://console.cloud.google.com/gcr/images/vds-client-amc/asia/notifier?project=vds-client-amc)
- [Cloud Run > slack-notifier](https://console.cloud.google.com/run/detail/asia-northeast1/slack-notifier/metrics?project=vds-client-amc)
- [Pub/Sub > cloud-builds](https://console.cloud.google.com/cloudpubsub/topic/detail/cloud-builds?project=vds-client-amc)
- [Cloud Storage > vds-client-amc-notifiers-config](https://console.cloud.google.com/storage/browser/vds-client-amc-notifiers-config;tab=objects?forceOnBucketsSortingFiltering=false&project=vds-client-amc&prefix=&forceOnObjectsSortingFiltering=false)
- [Secret > deploy-slack-notification](https://console.cloud.google.com/security/secret-manager/secret/deploy-slack-notification/versions?project=vds-client-amc)



