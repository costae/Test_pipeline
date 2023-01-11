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
}
