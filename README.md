# Hypershadow Contract

This repository provides all external interfaces and contracts that can be used for developing plugins within the
Hypershadow ecosystem.

## 🔌 Purpose

It serves as the foundation for cross-plugin and plugin-to-core communication, defining shared interfaces, data
structures, and extension points. This allows plugin developers to build and integrate safely without direct dependency
on internal application details.

## 📦 Dependency Injection (DI)

All dependencies that a plugin can **provide** or **consume** from the core application or another plugin-are managed
through dependency injection (DI).

You can explore the DI contract and utilities in the [di package](./di).
This enables loose coupling between plugins and the platform, and allows injecting services on-demand via shared
lifecycle rules.

## 🧩 Dependency Injection: Available Contracts

Only selected interfaces are exposed through the DI container for plugin consumption.

The following **contract types** can be safely requested from DI:

- `Client` - used to interact with the host system or other plugins
- `Builder`(s) - used as an intermediate layer (builder pattern) to access or configure a specific component

All other types are either:

- internal implementation details of plugins, or
- not injectable via the DI container by design

## 🗂️ Available Modules

The repository consists of multiple independent modules, each contained in its own folder. Every folder is:

- a standalone Go package
- importable via its own path
- versioned separately

## ✅ List of Available Packages

- [agent](./agent) - defines the global entity type identifier for Agent
    - [agent/ctx](./agent/ctx) - defines an interface for storing/retrieving Agent ID from context
    - [agent/find](./agent/find) - collection of interfaces for searching Agent models
        - [agent/find/byid](./agent/find/byid) - search for an Agent by its ID
    - [agent/hook](./agent/hook) - Agent hook client interface
    - [agent/httprouter](./agent/httprouter) - defines internal Agent HTTP router
    - [agent/model](./agent/model) - defines global Agent model
- [agenttoken](./agenttoken) - defines the global entity type identifier for AgentToken
    - [agenttoken/create](./agenttoken/create) - collection of interfaces for creates AgentToken models
        - [agenttoken/create/fromparams](./agenttoken/create/fromparams) - creates AgentToken from raw input parameters
    - [agenttoken/find](./agenttoken/find) - collection of interfaces for searching AgentToken models
        - [agenttoken/find/byid](./agenttoken/find/byid) - search for an AgentToken by its ID
        - [agenttoken/find/bylookupkey](./agenttoken/find/bylookupkey) - search for an AgentToken by its LookupKey
    - [agenttoken/hook](./agenttoken/hook) - AgentToken hook client interface
    - [agenttoken/model](./agenttoken/model) - defines global AgentToken model
    - [agenttoken/transport](./agenttoken/transport) - network-safe representation and utilities for AgentToken
      transport
- [api/httprouter](./api/httprouter) - defines public API HTTP router
- [apitoken](./apitoken) - defines the global entity type identifier for API token
    - [apitoken/find](./apitoken/find) - collection of interfaces for searching API token models
        - [apitoken/find/byid](./apitoken/find/byid) - search for an API token by its ID
    - [apitoken/hook](./apitoken/hook) - API token hook client interface
    - [apitoken/model](./apitoken/model) - defines global API token model
    - [apitoken/transport](./apitoken/transport) - network-safe representation and utilities for API token transport
- [archive](./archive) - defines base archiving interface for entity-based storage
- [auth](./auth) - token-based authentication and scoped access control
    - [auth/scope](./auth/scope) - defines access scopes for authorization
    - [auth/token](./auth/token) - defines the token interfaces
        - [auth/token/codec](./auth/token/codec) - defines interface for encoding/decoding Auth tokens
        - [auth/token/ctx](./auth/token/ctx) - defines interface for storing/retrieving Auth tokens in context
- [cache](./cache) - defines base cache interface
    - [cache/local](./cache/local) - in-memory cache implementation
- [choice](./choice) - defines base Choice abstractions
- [codec](./codec) - defines base serialization interface
- [crypt](./crypt) - cryptographic interface
- [db](./db) - core DB interface
- [dbhook](./dbhook) - event hook system for database query builders
- [di](./di) - dependency injection contracts
- [dispatcher/rest/schema](./dispatcher/rest/schema) - defines extended schema interface for REST dispatching
- [eb](./eb) - centralized error builder
    - [eb/impl](./eb/impl) - default implementation of the error builder interface
- [entity](./entity) - base entity types and identifiers
- [field](./field) - defines field abstractions
  [fielderror](./fielderror) - defines a field-level error interface and model
- [fmt](./fmt) - defines customizable fmt interface
- [hook](./hook) - event hook system for subscribing to and modifying operations in other plugins
    - [hook/impl](./hook/impl) - default implementation of hook registry and provider
- [httpauth](./httpauth) - dynamic scope builders for HTTP-based entity access control
- [httpserver](./httpserver) - HTTP server contracts, handlers, routing, middleware
    - [httpserver/cors](./httpserver/cors) - CORS handler builder for HTTP server
    - [httpserver/static](./httpserver/static) - Static handler interfaces
    - [httpserver/ws](./httpserver/ws) - WebSocket connection interfaces
- [id](./id) - unique identifier generation and conversion utilities
- [identity](./identity) - defines identity abstractions
- [integration](./integration) - defines the global entity type identifier for Agent
    - [integration/find](./integration/find) - collection of interfaces for searching Integration models
        - [integration/find/byid](./integration/find/byid) - search for an Integration by its ID
    - [integration/hook](./integration/hook) - Integration hook client interface
    - [integration/model](./integration/model) - defines global Integration model
    - [integration/schema](./integration/schema) - defines Schema interface for Integration structure
        - [integration/schema/impl](./integration/schema/impl) - defines base implementation of Integration Schema
          interface
- [json](./json) - JSON codec wrapper for serialization and streaming
    - [jsonint](./json/int) – defines int64 wrapper for marshaling/unmarshaling as JSON string.
- [meta](./meta) - key-value metadata container
    - [meta/json](./meta/json) - network-safe wrapper for meta.Meta, used for JSON transport and decoding
    - [meta/slog](./meta/slog) - converts meta.Meta into structured slog attributes for logging
- [metainfo](./metainfo) - defines base MetaInfo abstractions
- [operation](./operation) - defines the global entity type identifier for Operation
    - [operation/find](./operation/find) - collection of interfaces for searching Operation models
        - [operation/find/byid](./operation/find/byid) - search for an Operation by its ID
        - [operation/find/byids](./operation/find/byids) - search for an Operations by its IDs
    - [operation/hook](./operation/hook) - Operation hook client interface
    - [operation/model](./operation/model) - defines global Operation model
    - [operation/schema](./operation/schema) - defines Schema interface for Operation structure
    - [operation/schema/impl](./operation/schema/impl) - defines base implementation of Operation Schema interface
    - [operation/schema/metainfo](./operation/schema/metainfo) - defines metadata fields for describing Operation Schema
      interface
    - [operation/setauth](./operation/setauth) - defines an interface for injecting auth values into Operation
      parameters
    - [operation/validate](./operation/validate) - defines an interface for validating parameters against Operation
      Schemas
- [organization](./organization) - defines the global entity type identifier for Organization
    - [organization/ctx](./organization/ctx) - defines interface for storing/retrieving Organization ID in context
    - [organization/db](./organization/db) - defines interface for working with Organization DB
    - [organization/httprouter](./organization/httprouter) - defines internal Organization HTTP router
- [pager](./pager) - defines Pager abstractions
- [plugin](./plugin) - core Plugin interfaces
- [qb](./qb) – query builder interfaces
- [runner](./runner) - lifecycle-managed command execution framework
- [utiliter](./utiliter) - generic iterator transformation helpers
- [utilslice](./utilslice) - generic slice transformation helpers

## 📌 Versioning

Each package inside this repository is independently versioned. This allows for maximum flexibility in plugin
development and backward compatibility control.