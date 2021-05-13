package plugins

const (
	configurationAsCodePlugin           = "configuration-as-code:latest"
	gitPlugin                           = "git:latest"
	jobDslPlugin                        = "job-dsl:latest"
	kubernetesCredentialsProviderPlugin = "kubernetes-credentials-provider:latest"
	kubernetesPlugin                    = "kubernetes:latest"
	workflowAggregatorPlugin            = "workflow-aggregator:latest"
	workflowJobPlugin                   = "workflow-job:latest"
)

// basePluginsList contains plugins to install by operator.
var basePluginsList = []Plugin{
	Must(New(kubernetesPlugin)),
	Must(New(workflowJobPlugin)),
	Must(New(workflowAggregatorPlugin)),
	Must(New(gitPlugin)),
	Must(New(jobDslPlugin)),
	Must(New(configurationAsCodePlugin)),
	Must(New(kubernetesCredentialsProviderPlugin)),
}

// BasePlugins returns list of plugins to install by operator.
func BasePlugins() []Plugin {
	return basePluginsList
}
