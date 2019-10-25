pipeline {
  agent any
  stages {
  stage('Build') {
      steps {
        script {
            def root = tool name: 'Go', type: 'go'

            withEnv(["GOPATH=${env.WORKSPACE}/go", "GOROOT=${root}", "GOBIN=${root}/bin", "PATH+GO=${root}/bin"]) {
            sh "mkdir -p ${env.WORKSPACE}/go/src"

            echo 'Stage 1 - building'
            sh 'go version'
            sh 'docker-compose build'
            sh 'docker-compose up'

            }
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