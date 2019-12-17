#!/bin/bash

#获取shell进程
workplace="/opt/soft/mars/agent"

pcount=`ps -ef | grep "${workplace}/mars-agent" | grep -v grep | awk '{print $2}' | uniq -u | wc -l`

status="`cat ${workplace}/MARS_AGENT_STATUS`"

ps -ef | grep "${workdir}mars-agent" | grep -v grep

echo "pcount=${pcount}"

if [ "${status}" == "HANDLE" ]
then
    exit 0
elif [ ! -f ${workplace}/mars-agent -a ${pcount} -ge 1 ]
then
    kill -9 `ps -ef | grep "${workplace}/mars-agent" | grep -v grep | awk '{print $2}'`
    echo "stop process"
elif [ ${pcount} -ge 1 ]
then
    echo "process running"
else
    nohup /opt/soft/mars/agent/mars-agent 2>&1>/tmp/mars-agent.log &
    sleep 5s
    curl -I -s --connect-timeout 3 -m 30 --retry 3 --retry-delay 3 --retry-max-time 15 "http://0.0.0.0:8000/hc"
fi
