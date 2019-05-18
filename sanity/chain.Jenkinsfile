pipeline {
  agent any
  parameters {
    string(name: 'tcpPort', defaultValue: '12000', description: 'Port to run')
    string(name: 'chain', defaultValue: 'prysm', description: 'Blockchain to run')
    string(name: 'numNodes', defaultValue: '4', description: 'Number of nodes')
    string(name: 'setUpTime', defaultValue: '20', description: 'Seconds to wait before testing testnet')
  }
  stages {
    stage('Set up') {
      steps {
        println "Set up ${params.chain}"
        sh "rm -Rf ${params.chain};mkdir ${params.chain}"
        sh "~/bin/tester genesis testnet --blockchain ${params.chain} --numNodes ${params.numNodes} --volume `pwd`/${params.chain}:/var/output --file ./${params.chain}/testnetId"
        sleep params.setUpTime
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
