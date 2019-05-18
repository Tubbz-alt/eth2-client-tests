pipeline {
  agent any
  parameters {
    int(name: 'tcpPort', defaultValue: 12000, description: 'Port to run')
    string(name: 'chain', defaultValue: 'prysm', description: 'Blockchain to run')
  }
  stages {
    stage('Set up') {
      steps {
        println "Set up ${params.chain}"
        sh "rm -Rf ${params.chain};mkdir ${params.chain}"
        sh "~/bin/tester genesis testnet --blockchain ${params.chain} --numNodes 3 --volume `pwd`/${params.chain}:/var/output --file ./${params.chain}/testnetId"
        sleep 10
        sh "docker ps"
      }
    }

    stage('Test') {
      steps {
        println "Test ${params.chain}"
        sh "~/bin/tester network --testnet `cat ./${params.chain}/testnetId` --port ${params.tcpPort}"
      }
    }

    stage('Tear down') {
      steps {
        println "Tear down ${params.chain}"
        sh "~/bin/tester genesis destroy --testnetId `cat ./${params.chain}/testnetId`"
      }
    }

  }
}
