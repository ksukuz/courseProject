pipeline {
    agent {
    	label 'dev'	
    }

    stages {
        stage('Checkout') {
            steps {
                git url: 'https://github.com/ksukuz/courseProject.git'
            }
        }
        stage('Docker Build') {
            steps {
                script {
                  sh 'docker build -t app-go-image ./workspace/courseProject'
                  sh 'docker run -d --rm -p 8888:8888 app-go-image:latest'
                }
            }
        }
    }
}
