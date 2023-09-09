pipeline {
    agent { docker { image 'golang:1.21.1-alpine3.18' } }
    stages {
        stage('build') {
            steps {
                sh 'go version'
            }
        }
    }
}