package domain

// ConfigTemplate is a template of a configuration file.
const ConfigTemplate = `
---
# akoi - binary version control system
# https://github.com/suzuki-shunsuke/akoi
bin_path: /usr/local/bin/{{.Name}}-{{.Version}}
link_path: /usr/local/bin/{{.Name}}
# packages:
#   consul:
#     url: https://releases.hashicorp.com/consul/{{.Version}}/consul_{{.Version}}_darwin_amd64.zip
#     version: 1.2.1
#     files:
#     - name: consul
#       archive: consul
#   jq:
#     url: https://github.com/stedolan/jq/releases/download/jq-{{.Version}}/jq-osx-amd64
#     version: 1.5
#     archive_type: unarchived
#     files:
#     - name: jq
`