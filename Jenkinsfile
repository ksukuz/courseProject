pipeline {
    agent {
    	label 'devPC'	
    }

    stages {
        stage('Docker Build') {
            steps {
                script {
                  sh 'docker build -t app-go-image .'
                  sh 'docker run -d --rm -p 8888:8888 app-go-image:latest'
                }
            }
        }
    }
}
