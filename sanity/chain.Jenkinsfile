chains = ['prysm']
pipeline {
  agent any
  parameters {
    string(name: 'chain', defaultValue: 'prysm', description: 'Blockchain to run')
  }
  stages {
    stage('Set up') {
      steps {
        println "Set up ${params.chain}"
        sh "rm -Rf ${params.chain};mkdir ${params.chain}"
        sh "~/bin/tester genesis testnet --blockchain ${params.chain} --numNodes 3 --volume `pwd`/${params.chain}:/var/output --ports 9000:9000 --ports 9001:9000 --ports 9002:9000 --file ./${params.chain}/testnetId"
        sleep 10
        sh "docker ps"
      }
    }

    stage('Test') {
      steps {
        println "Test ${params.chain}"
        sh "telnet localhost 9000 | grep -v refused"
        sh "telnet localhost 9001 | grep -v refused"
        sh "telnet localhost 9002 | grep -v refused"
      }
    }

    stage('Tear down') {
      steps {
        println "Tear down ${params.chain}"
        sh "~/bin/tester genesis destroy --testnetId `cat ./${params.chain}/testnetId`"
        //sh "~/bin/tester network --testnet `cat ./${params.chain}/testnetId`"
      }
    }

  }
}
