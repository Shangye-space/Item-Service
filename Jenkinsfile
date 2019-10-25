pipeline {
  agent any
  stages {
  stage('Build') {
      steps {
        script {
          echo 'Stage 1 - building'
          sh 'go version'
          sh 'docker-compose build'
          sh 'docker-compose up'
        }
      }
    }
  stage('Tests') {
      steps {
        script {
          echo 'Stage 2 - testing'
        }
      }
    }
    stage('Approval') {
      steps {
        script {
          echo 'Stage 3 - approval'
        }
      }
    }
    stage('Deploy') {
      steps {
        script {
          echo 'Stage 4 - deploy'
        }
      }
    }
  }
  post { 
        always { 
            echo 'SUCCESS!'
        }
    }
}