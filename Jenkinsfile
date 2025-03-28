pipeline {
    agent any

    environment {
        DOCKER_IMAGE = "tapia254/cardon-tour-be:latest"
        AWS_REGION = "sa-east-1"
        ECS_CLUSTER = "CardonTour"
        ECS_SERVICE = "CardonTourTask"
    }

    stages {
        stage('Download Code') {
            steps {
                git 'https://github.com/tapiaw38/cardon-tour-be'
            }
        }

        stage('Autenticate with Docker Hub') {
            steps {
                withCredentials([usernamePassword(credentialsId: 'docker-hub-cred', usernameVariable: 'DOCKER_USER', passwordVariable: 'DOCKER_PASS')]) {
                    sh 'echo $DOCKER_PASS | docker login -u $DOCKER_USER --password-stdin'
                }
            }
        }

        stage('Deploy to AWS') {
            steps {
                withAWS(credentials: 'aws-credentials', region: "$AWS_REGION") {
                    sh 'aws ecs update-service --cluster $ECS_CLUSTER --service $ECS_SERVICE --force-new-deployment'
                }
            }
        }
    }
}
