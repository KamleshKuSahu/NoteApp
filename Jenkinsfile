pipeline {
    agent any

    stages {
        stage('Get Code') {
            steps {
                echo 'Getting code from repository'
                checkout scm
            }
        }

        stage('Check Files') {
            steps {
                echo 'Checking project structure'
                sh 'ls -la'
                sh 'ls -la backend/'
                sh 'ls -la frontend/'
            }
        }

        stage('Build Docker Images') {
            steps {
                echo 'Building Docker containers'
                sh 'docker compose build'
            }
        }

        stage('Deploy') {
            steps {
                echo 'Starting application'
                sh 'docker compose down || true'
                sh 'docker compose up -d'
            }
        }
    }
}