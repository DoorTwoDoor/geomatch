#
# Copyright 2017-present, DoorTwoDoor, Inc.
# All rights reserved.
#
# This source code is licensed under the Apache-style license found in the
# LICENSE file in the root directory of this source tree.
#

# Go command macros.
GOCMD    = go
GOAPP    = goapp
GOBUILD  = $(GOCMD) build
GOCLEAN  = $(GOCMD) clean
GOGET    = $(GOCMD) get
GOTEST   = $(GOCMD) test
GOVET    = $(GOCMD) vet
EXECNAME = geomatch

# Google Cloud Platform commnad-line interface macros.
GCLOUDCMD    = gcloud
GCLOUDAPP    = $(GCLOUDCMD) app
GCLOUDCONFIG = $(GCLOUDCMD) config
PROJECTNAME  = doortwodoor-7e677
APPVERSION   = v1

.PHONY: clean

all: build

build:
	@echo "Compiling packages and dependencies..."
	$(GOBUILD) -v -o $(EXECNAME)

clean:
	@echo "Removing object files..."
	$(GOCLEAN)

deploy:
	@echo "Deploying application to App Engine..."
	$(GCLOUDCONFIG) set project $(PROJECTNAME)
	$(GCLOUDAPP) deploy --version $(APPVERSION)

deps:
	@echo "Downloading and installing packages and dependencies..."
	$(GOGET) -u -v -d ./...

logs:
	@echo "Reading log entries from App Engine..."
	$(GCLOUDAPP) logs read

run: build
	@echo "Compiling and running Go program..."
	chmod u+x $(EXECNAME)
	./$(EXECNAME)

serve:
	@echo "Starting local development server..."
	$(GOAPP) serve

test:
	@echo "Testing packages..."
	$(GOTEST) -v ./...

vet:
	@echo "Examining source code..."
	$(GOVET) -v ./...
