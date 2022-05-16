#!/bin/bash

scp *.tgz vimgate-test001:/tmp
scp *.tgz vimgate-test002:/tmp
scp *.tgz vimgate-test003:/tmp
scp *.tgz vimgate-test004:/tmp
scp *.tgz vimgate-test005:/tmp
scp *.tgz vimgate-test006:/tmp
scp *.tgz vimgate-test007:/tmp
scp *.tgz vimgate-test008:/tmp
scp *.tgz vimgate-test009:/tmp

ssh vimgate-test001 'cd /tmp; gunzip *.tgz'
ssh vimgate-test002 'cd /tmp; gunzip *.tgz'
ssh vimgate-test003 'cd /tmp; gunzip *.tgz'
ssh vimgate-test004 'cd /tmp; gunzip *.tgz'
ssh vimgate-test005 'cd /tmp; gunzip *.tgz'
ssh vimgate-test006 'cd /tmp; gunzip *.tgz'
ssh vimgate-test007 'cd /tmp; gunzip *.tgz'
ssh vimgate-test008 'cd /tmp; gunzip *.tgz'
ssh vimgate-test009 'cd /tmp; gunzip *.tgz'
ssh vimgate-test001 'cd /tmp; ctr -n k8s.io images import *.tar --no-unpack'
ssh vimgate-test002 'cd /tmp; ctr -n k8s.io images import *.tar --no-unpack'
ssh vimgate-test003 'cd /tmp; ctr -n k8s.io images import *.tar --no-unpack'
ssh vimgate-test004 'cd /tmp; ctr -n k8s.io images import *.tar --no-unpack'
ssh vimgate-test005 'cd /tmp; ctr -n k8s.io images import *.tar --no-unpack'
ssh vimgate-test006 'cd /tmp; ctr -n k8s.io images import *.tar --no-unpack'
ssh vimgate-test007 'cd /tmp; ctr -n k8s.io images import *.tar --no-unpack'
ssh vimgate-test008 'cd /tmp; ctr -n k8s.io images import *.tar --no-unpack'
ssh vimgate-test009 'cd /tmp; ctr -n k8s.io images import *.tar --no-unpack'
ssh vimgate-test001 'rm -rf /tmp/*.tgz /tmp/*.tar'
ssh vimgate-test002 'rm -rf /tmp/*.tgz /tmp/*.tar'
ssh vimgate-test003 'rm -rf /tmp/*.tgz /tmp/*.tar'
ssh vimgate-test004 'rm -rf /tmp/*.tgz /tmp/*.tar'
ssh vimgate-test005 'rm -rf /tmp/*.tgz /tmp/*.tar'
ssh vimgate-test006 'rm -rf /tmp/*.tgz /tmp/*.tar'
ssh vimgate-test007 'rm -rf /tmp/*.tgz /tmp/*.tar'
ssh vimgate-test008 'rm -rf /tmp/*.tgz /tmp/*.tar'
ssh vimgate-test009 'rm -rf /tmp/*.tgz /tmp/*.tar'
