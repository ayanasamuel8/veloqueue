# Veloqueue

> A Postgres-backed job queue in Go, built to learn distributed
> systems by hitting failure modes directly.

## What is veloqueue

Veloqueue is a Postgres-backed job queue written in Go, built 
single-machine as a learning project. It accepts jobs through 
a Go API, persists them in Postgres, and workers pull and 
execute them. This gives at-least-once delivery — a job may 
run more than once if a worker crashes, so job handlers must 
be idempotent.

## V1 scope

Single device. Single database instance. No sharding, no 
scaling. Just a working API that accepts jobs and workers 
that execute them. Starting point.

## What v1 does NOT do

- No distributed system
- No database sharding
- No scaling
- No exactly-once delivery
- No job dependencies / pipelines
- No priorities

## Architecture

Client → API server → Postgres jobs table → Worker pool → execute

## Job lifecycle

>A job starts as pending. A worker moves it to running. On success → completed. On error → failed, then retried back to running with backoff. After max retries, dead.

## Open questions

Things I don't know yet but will figure out:
1. How to handle high job throughput
2. How to mitigate the gap between scheduled and actual run time
3. How to design idempotent job handlers
4. How to handle long-running jobs that crash mid-execution
5. How to handle job dependencies (probably v2+)