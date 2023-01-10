// Code generated by piper's step-generator. DO NOT EDIT.

package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/SAP/jenkins-library/pkg/config"
	"github.com/SAP/jenkins-library/pkg/log"
	"github.com/SAP/jenkins-library/pkg/splunk"
	"github.com/SAP/jenkins-library/pkg/telemetry"
	"github.com/SAP/jenkins-library/pkg/validation"
	"github.com/spf13/cobra"
)

type gctsRollbackOptions struct {
	Username                  string `json:"username,omitempty"`
	Password                  string `json:"password,omitempty"`
	Repository                string `json:"repository,omitempty"`
	Host                      string `json:"host,omitempty"`
	Client                    string `json:"client,omitempty"`
	Commit                    string `json:"commit,omitempty"`
	GithubPersonalAccessToken string `json:"githubPersonalAccessToken,omitempty"`
}

// GctsRollbackCommand Perfoms a rollback of one (default) or several commits
func GctsRollbackCommand() *cobra.Command {
	const STEP_NAME = "gctsRollback"

	metadata := gctsRollbackMetadata()
	var stepConfig gctsRollbackOptions
	var startTime time.Time
	var logCollector *log.CollectorHook
	var splunkClient *splunk.Splunk
	telemetryClient := &telemetry.Telemetry{}

	var createGctsRollbackCmd = &cobra.Command{
		Use:   STEP_NAME,
		Short: "Perfoms a rollback of one (default) or several commits",
		Long: `This step performs a rollback of commits in a local ABAP repository. If a ` + "`" + `commit` + "`" + ` parameter is specified, it will be used as the target commit for the rollback.
If no ` + "`" + `commit` + "`" + ` parameter is specified and the remote repository domain is 'github.com', the last commit with the status 'success' will be used for the rollback. Otherwise,
` + "`" + `gctsRollback` + "`" + ` will roll back to the previously active commit in the local repository.`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			startTime = time.Now()
			log.SetStepName(STEP_NAME)
			log.SetVerbose(GeneralConfig.Verbose)

			GeneralConfig.GitHubAccessTokens = ResolveAccessTokens(GeneralConfig.GitHubTokens)

			path, _ := os.Getwd()
			fatalHook := &log.FatalHook{CorrelationID: GeneralConfig.CorrelationID, Path: path}
			log.RegisterHook(fatalHook)

			err := PrepareConfig(cmd, &metadata, STEP_NAME, &stepConfig, config.OpenPiperFile)
			if err != nil {
				log.SetErrorCategory(log.ErrorConfiguration)
				return err
			}
			log.RegisterSecret(stepConfig.Username)
			log.RegisterSecret(stepConfig.Password)
			log.RegisterSecret(stepConfig.GithubPersonalAccessToken)

			if len(GeneralConfig.HookConfig.SentryConfig.Dsn) > 0 {
				sentryHook := log.NewSentryHook(GeneralConfig.HookConfig.SentryConfig.Dsn, GeneralConfig.CorrelationID)
				log.RegisterHook(&sentryHook)
			}

			if len(GeneralConfig.HookConfig.SplunkConfig.Dsn) > 0 {
				splunkClient = &splunk.Splunk{}
				logCollector = &log.CollectorHook{CorrelationID: GeneralConfig.CorrelationID}
				log.RegisterHook(logCollector)
			}

			if err = log.RegisterANSHookIfConfigured(GeneralConfig.CorrelationID); err != nil {
				log.Entry().WithError(err).Warn("failed to set up SAP Alert Notification Service log hook")
			}

			validation, err := validation.New(validation.WithJSONNamesForStructFields(), validation.WithPredefinedErrorMessages())
			if err != nil {
				return err
			}
			if err = validation.ValidateStruct(stepConfig); err != nil {
				log.SetErrorCategory(log.ErrorConfiguration)
				return err
			}

			return nil
		},
		Run: func(_ *cobra.Command, _ []string) {
			stepTelemetryData := telemetry.CustomData{}
			stepTelemetryData.ErrorCode = "1"
			handler := func() {
				config.RemoveVaultSecretFiles()
				stepTelemetryData.Duration = fmt.Sprintf("%v", time.Since(startTime).Milliseconds())
				stepTelemetryData.ErrorCategory = log.GetErrorCategory().String()
				stepTelemetryData.PiperCommitHash = GitCommit
				telemetryClient.SetData(&stepTelemetryData)
				telemetryClient.Send()
				if len(GeneralConfig.HookConfig.SplunkConfig.Dsn) > 0 {
					splunkClient.Send(telemetryClient.GetData(), logCollector)
				}
			}
			log.DeferExitHandler(handler)
			defer handler()
			telemetryClient.Initialize(GeneralConfig.NoTelemetry, STEP_NAME)
			if len(GeneralConfig.HookConfig.SplunkConfig.Dsn) > 0 {
				splunkClient.Initialize(GeneralConfig.CorrelationID,
					GeneralConfig.HookConfig.SplunkConfig.Dsn,
					GeneralConfig.HookConfig.SplunkConfig.Token,
					GeneralConfig.HookConfig.SplunkConfig.Index,
					GeneralConfig.HookConfig.SplunkConfig.SendLogs)
			}
			gctsRollback(stepConfig, &stepTelemetryData)
			stepTelemetryData.ErrorCode = "0"
			log.Entry().Info("SUCCESS")
		},
	}

	addGctsRollbackFlags(createGctsRollbackCmd, &stepConfig)
	return createGctsRollbackCmd
}

func addGctsRollbackFlags(cmd *cobra.Command, stepConfig *gctsRollbackOptions) {
	cmd.Flags().StringVar(&stepConfig.Username, "username", os.Getenv("PIPER_username"), "User that authenticates to the ABAP system. **Note** - Don't provide this parameter directly. Either set it in the environment, or in the Jenkins credentials store, and provide the ID as value of the `abapCredentialsId` parameter.")
	cmd.Flags().StringVar(&stepConfig.Password, "password", os.Getenv("PIPER_password"), "Password of the ABAP user that authenticates to the ABAP system. **Note** - Don´t provide this parameter directly. Either set it in the environment, or in the Jenkins credentials store, and provide the ID as value of the `abapCredentialsId` parameter.")
	cmd.Flags().StringVar(&stepConfig.Repository, "repository", os.Getenv("PIPER_repository"), "Specifies the name (ID) of the local repsitory on the ABAP system")
	cmd.Flags().StringVar(&stepConfig.Host, "host", os.Getenv("PIPER_host"), "Protocol and host of the ABAP system, including the port. Please provide in the format `<protocol>://<host>:<port>`. Supported protocols are `http` and `https`.")
	cmd.Flags().StringVar(&stepConfig.Client, "client", os.Getenv("PIPER_client"), "Specifies the client of the ABAP system to be addressed")
	cmd.Flags().StringVar(&stepConfig.Commit, "commit", os.Getenv("PIPER_commit"), "Specifies the target commit for the rollback")
	cmd.Flags().StringVar(&stepConfig.GithubPersonalAccessToken, "githubPersonalAccessToken", os.Getenv("PIPER_githubPersonalAccessToken"), "GitHub personal access token with at least read permissions for the remote repository")

	cmd.MarkFlagRequired("username")
	cmd.MarkFlagRequired("password")
	cmd.MarkFlagRequired("repository")
	cmd.MarkFlagRequired("host")
	cmd.MarkFlagRequired("client")
}

// retrieve step metadata
func gctsRollbackMetadata() config.StepData {
	var theMetaData = config.StepData{
		Metadata: config.StepMetadata{
			Name:        "gctsRollback",
			Aliases:     []config.Alias{},
			Description: "Perfoms a rollback of one (default) or several commits",
		},
		Spec: config.StepSpec{
			Inputs: config.StepInputs{
				Secrets: []config.StepSecrets{
					{Name: "abapCredentialsId", Description: "ID taken from the Jenkins credentials store containing user name and password of the user that authenticates to the ABAP system on which you want to execute the rollback.", Type: "jenkins"},
					{Name: "githubPersonalAccessTokenId", Description: "GitHub personal access token with at least read permissions for the remote repository", Type: "jenkins"},
				},
				Parameters: []config.StepParameters{
					{
						Name: "username",
						ResourceRef: []config.ResourceReference{
							{
								Name:  "abapCredentialsId",
								Param: "username",
								Type:  "secret",
							},
						},
						Scope:     []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:      "string",
						Mandatory: true,
						Aliases:   []config.Alias{},
						Default:   os.Getenv("PIPER_username"),
					},
					{
						Name: "password",
						ResourceRef: []config.ResourceReference{
							{
								Name:  "abapCredentialsId",
								Param: "password",
								Type:  "secret",
							},
						},
						Scope:     []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:      "string",
						Mandatory: true,
						Aliases:   []config.Alias{},
						Default:   os.Getenv("PIPER_password"),
					},
					{
						Name:        "repository",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{},
						Default:     os.Getenv("PIPER_repository"),
					},
					{
						Name:        "host",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{},
						Default:     os.Getenv("PIPER_host"),
					},
					{
						Name:        "client",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{},
						Default:     os.Getenv("PIPER_client"),
					},
					{
						Name:        "commit",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
						Default:     os.Getenv("PIPER_commit"),
					},
					{
						Name: "githubPersonalAccessToken",
						ResourceRef: []config.ResourceReference{
							{
								Name: "githubPersonalAccessTokenId",
								Type: "secret",
							},
						},
						Scope:     []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:      "string",
						Mandatory: false,
						Aliases:   []config.Alias{},
						Default:   os.Getenv("PIPER_githubPersonalAccessToken"),
					},
				},
			},
		},
	}
	return theMetaData
}