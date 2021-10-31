# Agent Allocation
This project is used to allocate agent. I'm using task queue backed by `postgres`.

This project has 2 service:
* Webhook: webhook will handle incoming custom agent allocation from qiscus and insert customer queue to postgres.
* Worker: worker will worked as a task queue that receive queue from postgres and assign agent to customer room.

This project has 4 layer:
* Entity Layer (this layer will contain data structure)
* Repository Layer (this will handle any db or external request transaction)
* Service Layer (this will handle bussiness logic, you might want to check /internal/worker/agent/service.go for agent allocation logic)
* Delivery Layer (this will handle how you deliver the service)

Project structure:

```tree
root:
├───cmd                       //this will contain main program to run
│   ├───webhook
│   └───worker
├───domain                    //this will contain all data structure for each service
├───internal
│   ├───webhook
│   │   └───customer         //this will contain customer webhook service
│   └───worker
│       └───agent            //this will contain agent worker service
└───util                     //this will contain global utility
    └───db                   //this will contain db connection
```
