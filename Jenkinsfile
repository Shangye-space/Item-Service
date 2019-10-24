### setup environment: creates directories, sshagent, checking
### deployment date, set PATH environment

stage "setup environment"
 node('master'){
 sshagent (credentials: ['2e1b7262-0912-409f-83fb-179860abee79']){
 sh 'ssh -o StrictHostKeyChecking=no -l ec2-user 111.222.123.121'
 sh 'DATE=`ssh ec2-user@111.222.123.121 "date +%Y%m%d%H%M%S"`'
 }
 dir('/var/lib/jenkins/jobs/[your_project]/workspace/bin'){}
 dir('/var/lib/jenkins/jobs/[your_project]/workspace/pkg'){}
 dir('/var/lib/jenkins/jobs/[your_project]/workspace/'){
 sh 'rm -rf *.tar.gz'
 }
 dir('/var/lib/jenkins/jobs/[your_project]/workspace/src/bitbucket.org/[project]/[project_go]'){
 sh 'rm -rf *'
 }
 withEnv(['PATH=$PATH:/opt/go/bin:/usr/local/bin','GOROOT=/opt/go','GOPATH=/var/lib/jenkins/jobs/[your_project]/workspace/']){
 dir('/var/lib/jenkins/jobs/[your_project]/workspace/src/bitbucket.org/[project]/[project_go]') {
 git branch: 'develop', credentialsId: 'c1db80fe-1eab-43f7-bd08-f4f80b706e61', url: 'git@bitbucket.org:rezast/golang.git'
 sh 'go get github.com/spf13/viper'
 sh 'go get gopkg.in/gin-gonic/gin.v1'
 sh 'go get github.com/op/go-logging'
 sh 'go get github.com/dgrijalva/jwt-go'
 sh 'go get github.com/stretchr/testify'
 }
 }
 }

### Run test script

stage "GO Test"
 node('master'){
 withEnv(['PATH=$PATH:/opt/go/bin','GOROOT=/opt/go','GOPATH=/var/lib/jenkins/jobs/[your_project]/workspace/']){
 dir('/var/lib/jenkins/jobs/[your_project]/workspace/src/bitbucket.org/[project]/[project_go]') {
 sh 'go test bill_test.go'
 }
 }
 }
 stage "Build Go!"
 node('master'){
 withEnv(['PATH=$PATH:/opt/go/bin:','GOROOT=/opt/go','GOPATH=/var/lib/jenkins/jobs/[your_project]/workspace/']){
 dir('/var/lib/jenkins/jobs/[your_project]/workspace/src/bitbucket.org/[project]/[project_go]'){
 sh 'go install'
 }
 }
 }

### Send binary to server

stage "Deploy Go"
 node('master'){
 dir('/var/lib/jenkins/jobs/[your_project]/workspace'){
 sh 'tar -czf deploy.tar.gz bin'
 sshagent (credentials: ['2e1b7262-0912-409f-83fb-179860abee79']){
 sh 'scp deploy.tar.gz ec2-user@111.222.123.121:~/'
 sh 'ssh ec2-user@111.222.123.121 "cd ~;tar xzf deploy.tar.gz"'
 echo 'Finished!'
 }
 }
 }