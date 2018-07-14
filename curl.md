#####List all projects 
curl -s -H "Content-Type:application/json" -H "X-RunDeck-Auth-Token:$API_TOKEN" http://$RUNDECK_HOST:4440/api/1/projects | xmlstarlet sel -t -v "//name/text()"

#####List all job ids 
curl -G -s -H "Content-Type:application/json" -H "X-RunDeck-Auth-Token:$API_TOKEN" http://$RUNDECK_HOST:4440/api/1/jobs -d project=PROJNAME | xmlstarlet sel -t -v "//job/@id"

#####List all job names 
curl -G -s -H "Content-Type:application/json" -H "X-RunDeck-Auth-Token:$API_TOKEN" http://$RUNDECK_HOST:4440/api/1/jobs -d project=PROJNAME | xmlstarlet sel -t -v "//job/name/text()"

#####List Recent executions based on period of time: http://rundeck.org/docs/api/#execution-query
curl -s -H "Content-Type:application/json" -H "X-RunDeck-Auth-Token:$API_TOKEN" http://$RUNDECK_HOST:4440/api/1/project/[Project]/executions -d "recentFilter=1w"
curl -s -H "Accept:application/json" -H "X-Rundeck-Auth-Token:GKrfka6yPg145IQuvvXZXbU2GxU5fKzJ" https://isca3.pgd.to/api/21/project/Teste/executions -d "recentFilter=1d"

#####List Older executions based on period of time: http://rundeck.org/docs/api/#execution-query
curl -s -H "Content-Type:application/json" -H "X-RunDeck-Auth-Token:$API_TOKEN" http://$RUNDECK_HOST:4440/api/1/project/[Project]/executions -d "olderFilter=30d"

#####Get execution ids of a job (comma separated) 
curl -G -s -H "Content-Type:application/json" -H "X-RunDeck-Auth-Token:$API_TOKEN" http://$RUNDECK_HOST:4440/api/1/job/$job_id/executions -d max=10 -d offset=200 | xmlstarlet sel -t -v //execution/@id | tr "\\n" "," | sed 's/,$//'

#####Delete executions 
curl -X POST -s -H "Content-Length:0" -H "X-RunDeck-Auth-Token:$API_TOKEN" http://$RUNDECK_HOST:4440/api/12/executions/delete?ids=$exec_id



