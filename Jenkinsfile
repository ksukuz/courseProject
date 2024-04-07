pipeline {
    agent {
    	label 'dev'	
    }

    stages {
        stage('Checkout') {
            steps {
                git url: 'https://github.com:ksukuz/courseProject.git'
            }
        }
        stage('Docker Build') {
       	    when {
       	    	anyOf {
       	    		branch 'master'
       	    		branch 'develop'
       	    		branch 'feature'
       	    	}
       	    }
            steps {
                script {
                  sh 'docker build -t app-go-image .'
                  sh 'docker run -d --rm -p 8888:8888 app-go-image:latest'
                }
            }
        }
    }
}
