pipeline {
    agent any

    stages {
        stage('Get Code') {
            steps {
                echo 'Getting code from repository'
                checkout scm
            }
        }

        stage('Test Backend') {
            steps {
                echo 'Testing Go backend'
                dir('backend') {
                    sh 'go version'
                    sh 'go mod tidy'
                }
            }
        }

        stage('Test Frontend') {
            steps {
                echo 'Testing React frontend'
                dir('frontend') {
                    sh 'npm --version'
                    sh 'npm ci'
                }
            }
        }
    }
}