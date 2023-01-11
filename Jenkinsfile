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
		integrationArtifactDownload script: this
		
	}
	stage('Unzip') {
            steps {
                sh 'unzip P2006255217-flow4.zip'
            }
        }
        stage('Git push') {
            steps {
                sh 'git add P2006255217-flow4'
                sh 'git commit -m "Adding the downloaded files"'
		sh 'git branch -M "P2006255217"'
                sh 'git push  --set-upstream https://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@github.com/costae/Test_pipeline.git P2006255217'
            }
        }
	
}
