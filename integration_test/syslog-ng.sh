#!/bin/bash

syslog-ng --no-caps --foreground -f syslog-ng.conf -c syslog-ng.control -R syslog-ng.persist -p `pwd`/syslog-ng.pid
