pipeline {
    agent any 

    environment {
        IMAGE_REPO_NAME="test_image"
        IMAGE_TAG="latest"

        REPOSITORY_URI = "public.ecr.aws/q5n4x1h2/main-repository"
        AWS_DEFAULT_REGION="us-east-1"
    }

    stages {
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

        stage('Test') { 
            steps {
                script {
                    echo 'Test Started!'
                }
            }
        }
        stage('Deploy') { 
            steps {
                script {
                    echo 'Deploy Started!'
                }
            }
        }
    }
}