# kube-dtn-api

## Overview

This is the API for the kube-dtn project. It includes the following API:

- Kubernetes Custom Resource Definition (CRD)
    - NetworkNode CRD

- gRPC API
    - NetworkLink: 
        - rpc ApplyLinks (LinksBatchQuery) returns (BoolResponse);

It also includes the following components which is convenient for developing:

- Kubernetes Custom Resource Definition (CRD)
    - NetworkNode ClientSet (provides a client for interacting with the NetworkNode CRD)

- gRPC API
    - Makefile for generating gRPC code automatically

## Usage

You should use git submodule to include this repository in your project. 

### Initialize the submodule

Change directory to the root of your project and run the following command:

```bash
git submodule add git@github.com:dtn-dslab/kube-dtn-api.git external_api
git submodule init
git submodule update
```

### Update the submodule

Change directory to the submodule(`external_api`) and use git like you would normally.

### Clone the project with submodule

If you are cloning a project with a submodule, you need to run the following command to initialize the submodule:

```bash
git clone --recurse-submodules YOUR_GIT_URL
```

or if you have already cloned the project, you can run the following command to initialize the submodule:

```bash
git submodule init
git submodule update
```

### Generate gRPC code

You can include `api.mk` provided in this repository in your Makefile to generate gRPC code automatically. 

```makefile
include external_api/api.mk
```

Add following lines to your .gitignore file:

```bash
bin/*
```

Then you can run the following command to generate gRPC code:

```bash
make proto
```