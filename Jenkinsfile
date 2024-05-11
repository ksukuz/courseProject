pipeline {
    agent {
        label 'MeRunner'
    }

    stages {
        stage('Build Docker Image') {
            steps {
                script {
                    sh 'docker build -t app-go-image .'
                }
            }
        }
        stage('Push Docker Image') {
            steps {
                script {
                    sh 'docker tag app-go-image kuzmitskaya/app-go-image:latest'
                    sh 'docker push kuzmitskaya/app-go-image:latest'
                }
            }
        }
        stage('Cleanup') {
            steps {
                script {
                    sh 'docker stop app-go-container || true'
                    sh 'docker rm app-go-container || true'
                }
            }
        }
        stage('Deploy Container') {
            steps {
                script {
                    sh 'docker run -d --name app-go-container --rm -p 8888:8888 app-go-image:latest'
                }
            }
        }
    }
}
