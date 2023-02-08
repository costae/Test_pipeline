library 'piper-lib-os'


node() {
  stage('init') {
    deleteDir()
	checkout scm
	def folder = "P2006255217-flow4";
    def filePath = folder + ".zip";
    zip dir: folder, glob: '', zipFile: filePath;
  }
//   stage('deployIntegrationArtifact and Get MPL Status') {
// //   	 setupCommonPipelineEnvironment script: this
// 	   integrationArtifactUpload script: this
//      integrationArtifactDeploy script: this
// 	   integrationArtifactGetMplStatus script: this
// 	   print "MPL Status:"
// 	   print  commonPipelineEnvironment.getValue("integrationFlowMplStatus")
//   }
	stage(' deploy all'){
		integrationArtifactDeploy script: this
// 		integrationArtifactGetMplStatus script: this
// 	   	print "MPL Status:"
// 	   	print  commonPipelineEnvironment.getValue("integrationFlowMplStatus")
		sh 'rm -rf test_mapping'
		sh 'rm -rf P2006255217-flow4'
		
// 		valueMappingArtifactDownload script: this
		valueMappingDeploy script: this
		messageMappingDeploy script: this
		scriptCollectionDeploy script: this
// 		integrationPackageDownload script: this
// 		sh 'mv CICD/test.zip CICD/test3.zip'
// 		valueMappingArtifactUpload script: this
		
		
	}
	stage('Download all') {
            
//                 sh 'unzip P2006255217-flow4/P2006255217-flow4.zip -d P2006255217-flow4/P2006255217-flow4'
// 		sh 'rm -rf P2006255217-flow4/P2006255217-flow4.zip'
		integrationArtifactDownload script: this
		valueMappingArtifactDownload script: this
		messageMappingDownload script: this
		scriptCollectionDownload script: this
		sh 'mv P2006255217-flow4/P2006255217-flow4.zip P2006255217-flow4/P2006255217-flow4_1.zip'
		sh 'mv test_mapping/test_mapping.zip test_mapping/test_mapping_1.zip'
		sh 'mv test_mess/test_mess.zip test_mess/test_mess_1.zip'
		sh 'mv test_script/test_script.zip test_script/test_script_1.zip'
            
        }
	stage('Upload all') {
            
		integrationArtifactUpload script: this
		valueMappingArtifactUpload script: this
		messageMappingUpload script: this
		scriptCollectionUpload script: this
            
        }
	stage('Download & Upload an Integration Package') {
            
		integrationPackageDownload script: this
		sh 'mv CICD/test.zip CICD/test3.zip'
		integrationPackageUpload script: this
            
        }
//         stage('Git push') {
            
// 		def branch = "test-CICD"
// 		def tkn = "ghp_UYl37YF3hh4Wux"
// 		def tkn2 = "GtuVDSwJ8I6nnsdi2Mr00E"
// 		def orgn = "https://"+tkn+tkn2+"@github.com/costae/Test_pipeline.git"
//                 sh 'git add P2006255217-flow4/*'
// 		sh 'git config --global user.email "you@example.com"'
// 		sh 'git config --global user.name "carlos"'
//                 sh 'git commit -m "Adding the downloaded files"'
// 		sh "git branch -M ${branch}"
// 		sh "git push ${orgn} ${branch}"
           
//         }
	
}
