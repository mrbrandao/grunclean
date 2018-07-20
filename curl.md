#####List all projects  
```curl -s -H "Content-Type:application/json" -H "X-RunDeck-Auth-Token:$API_TOKEN" http://$RUNDECK_HOST:4440/api/1/projects | xmlstarlet sel -t -v "//name/text()"```
  
#####List all job ids  
```curl -G -s -H "Content-Type:application/json" -H "X-RunDeck-Auth-Token:$API_TOKEN" http://$RUNDECK_HOST:4440/api/1/jobs -d project=PROJNAME | xmlstarlet sel -t -v "//job/@id"```
  
#####List all job names  
```curl -G -s -H "Content-Type:application/json" -H "X-RunDeck-Auth-Token:$API_TOKEN" http://$RUNDECK_HOST:4440/api/1/jobs -d project=PROJNAME | xmlstarlet sel -t -v "//job/name/text()"```
  
#####List Recent executions based on period of time: http://rundeck.org/docs/api/#execution-query  
```curl -s -H "Content-Type:application/json" -H "X-RunDeck-Auth-Token:$API_TOKEN" http://$RUNDECK_HOST:4440/api/1/project/[Project]/executions -d "recentFilter=1w"```
  
_This sample is the same as above but with max=0 setted will print all executions results, instead of default value which is only 20._  
```curl -s -H "Accept:application/json" -H "X-Rundeck-Auth-Token:GKrfka6yPg145IQuvvXZXbU2GxU5fKzJ" https://localhost:4440/api/21/project/Teste/executions -d "recentFilter=1d" -d "max=0"```
  
#####List Older executions based on period of time: http://rundeck.org/docs/api/#execution-query  
```curl -s -H "Content-Type:application/json" -H "X-RunDeck-Auth-Token:$API_TOKEN" http://$RUNDECK_HOST:4440/api/1/project/[Project]/executions -d "olderFilter=30d"```
  
#####Get execution ids of a job (comma separated)  
```curl -G -s -H "Content-Type:application/json" -H "X-RunDeck-Auth-Token:$API_TOKEN" http://$RUNDECK_HOST:4440/api/1/job/$job_id/executions -d max=10 -d offset=200 | xmlstarlet sel -t -v //execution/@id | tr "\\n" "," | sed 's/,$//'```
  
#####Delete executions  
```curl -X POST -s -H "Content-Length:0" -H "X-RunDeck-Auth-Token:$API_TOKEN" http://$RUNDECK_HOST:4440/api/12/executions/delete?ids=$exec_id```

####Delete all executions for a job
```curl -s -X DELETE -H "Accept:application/json" -H "X-RunDeck-Auth-Token:$API_TOKEN" $RUNDECK_HOST/api/21/job/265b8dc0-47a4-43d3-8f5f-d6fa0f4db41a/executions```  


####Rundeck heap problem
```ERROR GrailsExceptionResolver - OutOfMemoryError occurred when processing request: [POST] /api/21/project/Teste/executions
GC overhead limit exceeded. Stacktrace follows:
org.codehaus.groovy.grails.web.servlet.mvc.exceptions.ControllerExecutionException: Executing action [apiExecutionsQueryv14] of controller [rundeck.controllers.ExecutionController]  caused ex
ception: Runtime error executing action
...
Caused by: java.lang.reflect.InvocationTargetException
... 47 more
Caused by: java.lang.OutOfMemoryError: GC overhead limit exceeded```
---> References http://rundeck.org/docs/administration/tuning-rundeck.html

####Rundeck too many connections error :bug: Need to fix routines to avoid this
com.mysql.jdbc.exceptions.jdbc4.MySQLNonTransientConnectionException: Too many connections
