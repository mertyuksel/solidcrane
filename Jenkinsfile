pipeline {
    agent any 

    environment {
        IMAGE_REPO_NAME="test_image"
        IMAGE_TAG="latest"
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

        stage('Test') { 
            steps {
                // 
            }
        }
        stage('Deploy') { 
            steps {
                // 
            }
        }
    }
}