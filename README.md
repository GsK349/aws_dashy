# aws_dashy

This application interact with the Rest API's of the AWS to generate a cost report spent in AWS

## Local setup
1. [setup GoLang on your local machine.](https://go.dev/doc/install)
2. [install AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)
3. run `aws configure` in your command line.
4. populate  `AWS Access Key ID`, `AWS Secret Access Key`, `Default region name` to complete the AWS local setup

## Fork setup
1. Fork the repo in your GitHub Account
2. Setup Reposetory secrets `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, `AWS_region`
3. run `run.yml` to run the application
      - Click on Actions button
      - Click on run code
      - click on `Run workflow` button
      - select branch and run
