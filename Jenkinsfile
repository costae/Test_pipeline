library 'piper-lib-os'


node() {
  stage('init') {
    deleteDir()
	checkout scm
	def folder = "P2006255217-flow4";
    def filePath = folder + ".zip";
    zip dir: folder, glob: '', zipFile: filePath;
  }
  stage('deployIntegrationArtifact and Get MPL Status') {
  	 setupCommonPipelineEnvironment script: this
	   integrationArtifactUpload script: this
     integrationArtifactDeploy script: this
	   integrationArtifactGetMplStatus script: this
	   print "MPL Status:"
	   print  commonPipelineEnvironment.getValue("integrationFlowMplStatus")
  }
	stage('undeployIntegrationArtifact and Download Artifact'){
		integrationArtifactUnDeploy script: this
		integrationArtifactGetMplStatus script: this
	   	print "MPL Status:"
	   	print  commonPipelineEnvironment.getValue("integrationFlowMplStatus")
		sh 'rm -rf P2006255217-flow4'
		integrationArtifactDownload script: this
		
	}
	stage('Unzip') {
            
                sh 'unzip P2006255217-flow4/P2006255217-flow4.zip -d P2006255217-flow4/P2006255217-flow4'
		sh 'rm -rf P2006255217-flow4/P2006255217-flow4.zip'
            
        }
        stage('Git push') {
            
		def branch = "test-CICD"
		def tkn = " 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"
                sh 'git add P2006255217-flow4/*'
		sh 'git config --global user.email "you@example.com"'
		sh 'git config --global user.name "carlos"'
                sh 'git commit -m "Adding the downloaded files"'
		sh 'git push "https://${tkn}@github.com/costae/Test_pipeline.git" ${branch}:${branch}'
           
        }
	
}
