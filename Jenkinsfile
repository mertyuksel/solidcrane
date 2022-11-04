pipeline {
    agent any 

    environment {
        IMAGE_REPO_NAME="test_image"
        IMAGE_TAG="latest"

        REPOSITORY_URI = "public.ecr.aws/q5n4x1h2/main-repository"
        AWS_DEFAULT_REGION="us-east-1"

        SCANNER_HOME="C:\\Programs\\sonar-scanner-4.7.0.2747-windows"
    }

    stages {

        stage('Run Tests') {
            steps {
                echo 'Test Started!'
                bat """go test -v"""  
            }
        }

        stage('Static Code Analysis') {
            steps {
                echo 'Static Code Analysis Started!'
                
                withSonarQubeEnv('sq1') { 
                     bat "${SCANNER_HOME}\\bin\\sonar-scanner"
                }                
            }
        }

        stage('Build Docker Image') { 
            steps {
                script {
                    echo 'Build Started!'
                    dockerImage = docker.build "${IMAGE_REPO_NAME}:${IMAGE_TAG}"
                }
            }
        }

        stage('Logging into AWS ECR') {
            steps {
                script {
                    echo 'Logging into AWS ECR Started!'
                    // aws ecr-public get-login-password --region us-east-1 | docker login --username AWS --password-stdin public.ecr.aws
                    bat """aws ecr-public get-login-password --region ${AWS_DEFAULT_REGION} | docker login --username AWS --password-stdin public.ecr.aws"""
                }
            }
        }

        stage('Push the image to ECR') {
            steps {
                script {
                    echo 'Push the image to ECR Started!'

                    bat """docker tag ${IMAGE_REPO_NAME}:${IMAGE_TAG} ${REPOSITORY_URI}:$IMAGE_TAG"""
                    bat """docker push ${REPOSITORY_URI}"""
                }
            }
        }

        stage('Update ECS') {
            steps {
                script {
                    echo 'Update ECS Started!'

                    // pulls the most recent image version
                    // duplicates the current task definition
                    // cause the service to redeploy all running tasks
                    bat """ecs deploy main-cluster main-service"""
                }
            }
        }

    }
}