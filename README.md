### GrunClean

_GrunClean extends for `Go Rundeck Cleaner`, it's a tool to clean the rundeck  
executions by rundeck api._  

![](gif/grunclean.gif?raw=true)

### Use

This tool follows the K.I.S.S philosophy style, to do one thing and do well.  
So to use it's also very simple, you'll need only two things basically:  
* 1 - The rundeck url  
_Like: https://myrundeck.com_  
* 2 - A rundeck token  
_Read more, how to make a rundeck token [here](http://rundeck.org/docs/api/index.html#token-authentication)._  

Now you'll be able to use `grunclean` trough command line like this:  


```
grunclean -url https://myrundeck.com -token mytokenhash
```

The result of this command will be the list of the projects on your rundeck.  
  

1. Listing executions:  
  
```grunclean -url https://myrundeck.com -token mytokenhash -type exec```
  
_This command will list all the last 20 executions older than 1 day._ 

2. Querying all executions older than 30 days:  
  
```grunclean -url https://myrundeck.com -token mytokenhash -type exec -period 30d -max 0```
  
_This command use max=0 to list all executions and period=30d to list the executions older than 30 days._  

*WARNING!*
_Use `max=0` narrowed by few days like 1day or 1week, if you try to list Millions of thousans  
executions you probably will hurt your rundeck feelings. I tested some curl executions with  
150000 executions but i have to tune my rundeck session to `-Xmx4096m -Xms1024` to lead with  
this query. Take care with `OutOfMemoryError`, see more about tunning [here](http://rundeck.org/docs/administration/tuning-rundeck.html)._  
  

The flag `-period` use the same values of rundeck [execution query](http://rundeck.org/docs/api/#execution-query) which is:  
`h`: hour  
`d`: day  
`w`: week  
`m`: month  
`y`: year  

You can also change the flag `-query` to specify if it will be `older` or `recent` period.  
For e.g:  
  
```grunclean -url https://myrundeck.com -token mytokenhash -type exec -period 30d -max 0 -query recent```
  
_This query brings only the executions more recent than 30 days._
  
3. Listing executions from a specific job.  
  
```grunclean -url https://myrundeck.com -token mytokenhash -type exec -period 30d -max 0 -name MyJobName```
  
_This command will return all the executions older than 30 days only for the job `MyJobName`_

4. Listing jobs:
  
Sure you can list all your jobs  very easily.  
  
```grunclean -url https://myrundeck.com -token mytokenhash -type job```
  
_With the flag `-type` you can list jobs, executions, and projects, Just set the type like : job, exec or proj._  

5. Deleting a job execution:
  
You can delete an execution as the same way you can list them. Just add the flag `-action delete`  
and it will be deleted. For example:  
  
```grunclean -url https://myrundeck.com -token mytokenhash -type exec -max 1 -query recent -action delete```  
_This command will delete only the last one more recent execution._
  
  
```grunclean -url https://myrundeck.com -token mytokenhash -type exec -max 0 -period 60d -action delete -name MyJobName```  
_This command will delete all executions older than 60 days only for the job `MyJobName`._
  
  
```grunclean -url https://myrundeck.com -token mytokenhash -type exec -max 0 -period 7w -action delete```  
_This command will delete all executions older then 7 weeks for all projects._  
  
  
### Installation  

```
go get github.com/isca0/grunclean
```


### Author  
  
Igor Brandao [isca](isca.space)  
  
Hope you Enjoy this tool... :wink:  
  
