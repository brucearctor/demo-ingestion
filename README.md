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

* Missing Loadbalancer

* How to handle security for the incoming data?

* revise proto fields

* Read data -- share how...

* Fix code for 'inair' flights

* Document architecture and add to README

* Get checkov action running ... plus what other security?

* Cleanup README to be followable

* emulators rather than 'real' deploy [ esp. thinking ahead for initial testing and quicker development ]

#### 'edge' cases
* What to do on start of flight/takeoff [ and if that is missing ]
* same with landed




### corefirestore [ function ]

* remove extra logging functions
* firestoreClient to init [ for reuse ]



## Devcontainer

Get Buf Auth working, so not run into the rate limiting when unauthenticated using remote plugins [ https://buf.build/docs/bsr/rate-limits ]






###### Generic Cleanup

[Best Practices for Cloud Functions](https://cloud.google.com/functions/docs/bestpractices/tips)





###### Remember

* if redeploying a function, with a push subscription, this might take a bit for permissions to propogate
*** Strangely it looks like it often takes a second TOFU Apply for the IAM to get applied.

* will need to understand expected frequency of updates, tolerance for delays, and more.  Several ways to handle.
