pipeline {
    agent none
    stages {
        stage('Build') {
            agent {
                docker { image 'obraun/vss-micro-jenkins' }
            }
            steps {
                sh 'cd movie_service && go build movie.go'
                sh 'cd room_service && go build room.go'
                sh 'cd user_service && go build user.go'
                sh 'cd screening_service && go build screening.go'
                sh 'cd reservation_service && go build reservation.go'
                sh 'cd client && go build client.go'
            }
        }
        stage('Test') {
            agent {
                docker { image 'obraun/vss-micro-jenkins' }
            }
            steps {
                sh 'echo run tests...'
                sh 'go test ./...'
            }
        }
        stage('Build Docker Image') {
            agent any
            steps {
                sh "docker-build-and-push -b ${BRANCH_NAME} -s movie_service -f movie_service.dockerfile"
                sh "docker-build-and-push -b ${BRANCH_NAME} -s room_service -f room_service.dockerfile"
                sh "docker-build-and-push -b ${BRANCH_NAME} -s user_service -f user_service.dockerfile"
                sh "docker-build-and-push -b ${BRANCH_NAME} -s screening_service -f screening_service.dockerfile"
                sh "docker-build-and-push -b ${BRANCH_NAME} -s reservation_service -f reservation_service.dockerfile"
                sh "docker-build-and-push -b ${BRANCH_NAME} -s client -f client.dockerfile"
            }
        }
        stage('Run Example') {
            agent any
            steps {
                sh "docker-compose up --exit-code-from client"
            }
        }
    }
}
