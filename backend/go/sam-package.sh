#!/bin/bash
# set -e
# set -x

PROGNAME=$(basename -- "${0}")
PROJROOT=$(d=$(dirname -- "${0}"); cd "${d}/.." && pwd)

DT=$(date +%Y)
SEMVER=0.0.16
S3PREFIX="packaged/$DT/0.0.1/acentera-${PLUGINNAME}"
BUCKETNAME=${S3_BUCKET:-"lambda-at-edge-dev-serverlessdeploymentbucket-1gmbbmp4ajnba"}

                                                                                
if [ -e .aws ]; then
  source .aws
fi

cp -f template.yml .template.yml
# First update the Path: .... no hard-coded value ideally                     
sed -ri "s~<%PLUGIN_NAME%>~${PLUGINNAME}~g" .template.yml                  
sed -ri "s~<%SEMVER%>~${SEMVER}~g" .template.yml
[[ -e packaged-template.yml ]] && rm -f .packaged-template.yml

sam package --debug --template-file .template.yml --output-template-file output.yml --s3-bucket ${BUCKETNAME} --s3-prefix ${S3PREFIX}
