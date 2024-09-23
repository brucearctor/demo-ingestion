# demo-ingestion

## Trying some 'serverless' rather 'fast' data ingestion pipelines...



### Generating Protos
`buf generate --exclude-path pubsub_proto_stripped`

we need to ignore the proto in there

THere are some other ways to handle with a bit more engineering.

Also/otherwise, maybe there is an option for exclude path in relevnt buf yaml file?


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

## Transcoding
I forgot about grpc-gateway :-/




##### TODO

* configure buf registry ...

* get projectid/topicname from env vars, not hardcoded...

* KMS


### corefirestore [ function ]

* remove extra logging functions
* firestoreClient to init [ for reuse ]



## Devcontainer

Get Buf Auth working, so not run into the rate limiting when unauthenticated using remote plugins [ https://buf.build/docs/bsr/rate-limits ]
