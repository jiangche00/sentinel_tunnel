#!/bin/bash

scp *.tgz vimgate-test001:/tmp
scp *.tgz vimgate-test002:/tmp
scp *.tgz vimgate-test003:/tmp
scp *.tgz vimgate-test004:/tmp
scp *.tgz vimgate-test005:/tmp
scp *.tgz vimgate-test006:/tmp

ssh vimgate-test001 'cd /tmp; gunzip -c *.tgz | docker load; rm -rf *.tgz'
ssh vimgate-test002 'cd /tmp; gunzip -c *.tgz | docker load; rm -rf *.tgz'
ssh vimgate-test003 'cd /tmp; gunzip -c *.tgz | docker load; rm -rf *.tgz'
ssh vimgate-test004 'cd /tmp; gunzip -c *.tgz | docker load; rm -rf *.tgz'
ssh vimgate-test005 'cd /tmp; gunzip -c *.tgz | docker load; rm -rf *.tgz'
ssh vimgate-test006 'cd /tmp; gunzip -c *.tgz | docker load; rm -rf *.tgz'
