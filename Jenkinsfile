pipeline {
    agent any
    
    stages {
        stage('1. Build Docker Image') {
            steps {
                echo '=== Memasak Image Anteraja ==='
                // Jenkins bakal otomatis ngejalanin docker build
                sh 'docker build -t anteraja-app:latest .'
            }
        }
        
        stage('2. Deploy to Docker') {
            steps {
                echo '=== Menjalankan Web ==='
                // Hapus container lama kalau ada, lalu jalankan yang baru
                sh '''
                docker rm -f anteraja-web || true
                docker run -d -p 8080:8080 --name anteraja-web anteraja-app:latest
                '''
            }
        }
    }
}