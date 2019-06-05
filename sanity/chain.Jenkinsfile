pipeline {
  agent any
  parameters {
    string(name: 'tcpPort', defaultValue: '12000', description: 'Port to run')
    string(name: 'chain', defaultValue: 'prysm', description: 'Blockchain to run')
    string(name: 'numNodes', defaultValue: '4', description: 'Number of nodes')
    string(name: 'setUpTime', defaultValue: '600', description: 'Seconds to wait before testing testnet')
    string(name: 'privateKey', defaultValue: '', description: 'Goerli private key')
  }
  stages {
    stage('Set up') {
      steps {
        sh "rm -Rf ${params.chain};mkdir ${params.chain};rm -Rf reports;mkdir -p reports"
        println "Set up deposit contract"
        sh "build/bin/tester contract --priv-key ${params.privateKey} --output-file ./${params.chain}/contract"
        println "Set up ${params.chain}"
        sh "~/bin/tester genesis testnet --blockchain ${params.chain} --numNodes ${params.numNodes} --logFolder `pwd`/${params.chain} --file ./${params.chain}/testnetId --contract `cat ./${params.chain}/contract`"
        println "Send transactions to deposit contract"
        sendTxs("${params.numNodes}" as Integer)
        sleep params.setUpTime
        sh "docker ps"
      }
    }

    stage('Test network connectivity') {
      steps {
        println "Test network connectivity for ${params.chain}"
        sh "~/bin/tester network --testnet `cat ./${params.chain}/testnetId` --port ${params.tcpPort}  --testoutput reports/${params.chain}-network.xml"
      }
    }

    stage('Test Prometheus reporting') {
      steps {
        println "Test prometheus reporting for ${params.chain}"
        sh "~/bin/tester prometheus up --testnet `cat ./${params.chain}/testnetId` --testoutput reports/${params.chain}-prometheus.xml"
      }
    }

    stage('Test Consensus') {
          steps {
            println "Test prometheus reporting for ${params.chain}"
            sh "~/bin/tester consensus --type finalized_block_root --folder `pwd`/${params.chain} --blockchain ${params.chain} --testoutput reports/${params.chain}-finalized_block_root.xml"
            sh "~/bin/tester consensus --type finalized_state_root --folder `pwd`/${params.chain} --blockchain ${params.chain} --testoutput reports/${params.chain}-finalized_state_root.xml"
            sh "~/bin/tester consensus --type justified_block_root --folder `pwd`/${params.chain} --blockchain ${params.chain} --testoutput reports/${params.chain}-justified_block_root.xml"
            sh "~/bin/tester consensus --type justified_state_root --folder `pwd`/${params.chain} --blockchain ${params.chain} --testoutput reports/${params.chain}-justified_state_root.xml"
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

//No NonCPS required
def sendTxs(numberOfNodes) {
  for (int i = 0; i < numberOfNodes; i++) {
    sh "~/bin/tester sendTx --priv-key ${params.privateKey} --password ./${params.chain}/password${i} --keystore ./${params.chain}/key{i} --contract `cat ./${params.chain}/contract` --amount 3200"
  }
}
