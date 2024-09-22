# demo-ingestion

## Trying some 'serverless' rather 'fast' data ingestion pipelines...




### OpenTofu Explore

This is my first exploration with OpenToFu, so adding/exploring some new tool(s)

ex: https://github.com/tofuutils/pre-commit-opentofu and https://github.com/gamunu/vscode-opentofu




## CREDS
Esp. on 'new' devcontainer container:

`gcloud init` .. and to through login flow
also `gcloud auth application-default login` -- considering this acceptable for development, esp. ahead of better CI.
These are both while doing local/testing, ahead of tofu deploys happening as part of CI.



# Collector


## RUNNING ( locally )
GCP_PROJECT=brucearctor-demo-ingestion TOPIC=demo-topic FUNCTION_TARGET=ReceiveAndPublish LOCAL_ONLY=true go run cmd/main.go






##### TODO

* configure buf registry ...
