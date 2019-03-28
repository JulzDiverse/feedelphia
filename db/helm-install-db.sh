#!/bin/bash

helm install stable/mysql \
  --name feed-mysql \
  --namespace feed-db \
  --set persistence.storageClass=hostpath,mysqlDatabase=feedelphia
