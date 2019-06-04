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
        sh "rm -Rf ${params.chain};mkdir ${params.chain};rm -Rf reports;mkdir -p reports"
        sh "~/bin/tester genesis testnet --blockchain ${params.chain} --numNodes ${params.numNodes} --logFolder `pwd`/${params.chain}:/var/output --file ./${params.chain}/testnetId"
        sleep params.setUpTime
        sh "docker ps"
      }
    }

    stage('Test network connectivity') {
      steps {
        println "Test network connectivity for ${params.chain}"
        sh "~/bin/tester network --testnet `cat ./${params.chain}/testnetId` --port ${params.tcpPort}"
      }
    }

    stage('Test Prometheus reporting') {
      steps {
        println "Test prometheus reporting for ${params.chain}"
        sh "~/bin/tester prometheus up --testnet `cat ./${params.chain}/testnetId`"
      }
    }

    stage('Test Consensus') {
          steps {
            println "Test prometheus reporting for ${params.chain}"
            sh "~/bin/tester consensus --type finalized_block_root --folder `pwd`/${params.chain} --blockchain ${params.chain} --testoutput reports/${params.chain}-finalized_block_root.xml"
            sh "~/bin/tester consensus --type finalized_state_root --folder `pwd`/${params.chain} --blockchain ${params.chain} --testoutput reports/${params.chain}-finalized_state_root.xml"
            sh "~/bin/tester consensus --type justified_block_root --folder `pwd`/${params.chain} --blockchain ${params.chain} --testoutput reports/${params.chain}-justified_block_root.xml"
            sh "~/bin/tester consensus --type finalized_state_root --folder `pwd`/${params.chain} --blockchain ${params.chain} --testoutput reports/${params.chain}-finalized_state_root.xml"
          }
        }

    stage('Tear down') {
      steps {
        println "Tear down ${params.chain}"
        sh "~/bin/tester genesis destroy --testnet `cat ./${params.chain}/testnetId`"
      }
    }

  }
  post {
    always {
      junit 'reports/*.xml'
    }
  }
}
