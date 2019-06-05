pipeline {
  agent any
  parameters {
    string(name: 'tcpPort', defaultValue: '12000', description: 'Port to run')
    string(name: 'chain', defaultValue: 'prysm', description: 'Blockchain to run')
    string(name: 'numNodes', defaultValue: '4', description: 'Number of nodes')
    string(name: 'setUpTime', defaultValue: '600', description: 'Seconds to wait before testing testnet')
    credentials(name: 'privateKey', defaultValue: '', description: 'Goerli private key', credentialType: "Secret text", required: true)
  }
  environment {
          PRIVATE_KEY = credentials("${params.privateKey}")
  }
  stages {
    stage('Set up') {
      steps {
        sh "sudo rm -Rf ${params.chain};mkdir ${params.chain};rm -Rf reports;mkdir -p reports"
        println "Set up deposit contract"
        sh "~/bin/tester contract --priv-key " + PRIVATE_KEY + " --output-file ./${params.chain}/contract"
        println "Set up ${params.chain}"
        sh "~/bin/tester genesis testnet --blockchain ${params.chain} --numNodes ${params.numNodes} --logFolder `pwd`/${params.chain} --file ./${params.chain}/testnetId --contract `cat ./${params.chain}/contract` --validatorsPassword password"
        println "Wait for build to finish"
        sh "~/bin/tester genesis build-status --testnet `cat ./${params.chain}/testnetId`"
        println "Send transactions to deposit contract"
        sh "sudo chmod -R 644 ./${params.chain}/key*"
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
      archiveArtifacts artifacts: './${params.chain}/*', fingerprint: true
    }
  }
}

def sendTxs(numberOfNodes) {
  for (int i = 0; i < numberOfNodes; i++) {
    for (int j = 1 ; j <= 8 ; j++) {
      sh "~/bin/tester sendTx --priv-key " + PRIVATE_KEY + " --password password --keystore ./${params.chain}/key${i}-${j} --contract `cat ./${params.chain}/contract` --amount 3200"
    }
  }
}
