# Agent Allocation
This project use [DDD](https://en.wikipedia.org/wiki/Domain-driven_design) Concept for project structur. I'm using task queue backed by `postgres`.

This project has 2 service:
* Webhook: webhook will handle incoming custom agent allocation from qiscus and insert customer queue to redis.
* Worker: worker will worked as a task queue that receive queue from redis and assign agent to customer room.

This project has 4 domain layer:
* Domain Layer (entity)
* Repository Layer (this will handle any db or external request transaction)
* Service Layer (this will handle bussiness logic, you might want to check /internal/worker/agent/service.go for agent allocation logic)
* Delivery Layer (this will handle how you deliver the service)

Project structure:

```tree
root:
├───cmd                       //this will contain main program to run
│   ├───webhook
│   └───worker
├───domain                    //this will contain all data structure for each domain
├───internal
│   ├───webhook
│   │   └───customer         //this will contain customer webhook service
│   └───worker
│       └───agent            //this will contain agent worker service
└───util                     //this will contain global utility
```
