#!/bin/sh

#start nsq services
vagrant nsqlookupd &
vagrant nsqd --lookupd-tcp-address=127.0.0.1:4160 &
vagrant nsqadmin --lookupd-http-address=127.0.0.1:4161 &