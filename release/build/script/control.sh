#!/bin/bash

#获取shell进程
workplace="/opt/soft/mars/agent"

case "$1" in
  start)
        echo "Starting Service"
        nohup /opt/soft/mars/agent/mars-agent 2>&1>/tmp/mars-agent.log &
        sleep 5s
        curl -I -s --connect-timeout 3 -m 30 --retry 3 --retry-delay 3 --retry-max-time 15 "http://0.0.0.0:8000/hc"
        echo "RUNNING" > ${workplace}/MARS_AGENT_STATUS
        echo "Start Service "
        ;;
  stop)
        echo "Stopping Service"
        echo "HANDLE" > ${workplace}/MARS_AGENT_STATUS
        kill -9 `ps -ef | grep "${workplace}"/mars-agent | grep -v grep | awk '{print $2}'`
        echo "Stopped Service"
        ;;
  status)
        cat ${workplace}/MARS_AGENT_STATUS
        ;;
  restart|reload)
        echo "Stoping Service"
        kill -9 `ps -ef | grep "${workplace}/mars-agent" | grep -v grep | awk '{print $2}'`
        echo "Start Service"
        nohup /opt/soft/mars/agent/mars-agent 2>&1>/tmp/mars-agent.log &
        sleep 5s
        curl -I -s --connect-timeout 3 -m 30 --retry 3 --retry-delay 3 --retry-max-time 15 "http://0.0.0.0:8000/hc"
        echo "RUNNING" > ${workplace}/MARS_AGENT_STATUS
        ;;
  *)
        echo "Usage: ${workpalce}/control.sh {start|stop|status|restart}"
        exit 1
esac
