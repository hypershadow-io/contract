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

- [auth](./auth) - token-based authentication and scoped access control
- [codec](./codec) - generic serialization interface
- [crypt](./crypt) - cryptographic interface
- [di](./di) - dependency injection contracts
- [eb](./eb) - centralized error builder
- [ebimpl](./ebimpl) - default implementation of the error builder interface
- [entity](./entity) - base entity types and identifiers
- [fmt](./fmt) - customizable fmt interface
- [hook](./hook) - event hook system for subscribing to and modifying operations in other plugins
- [hookimpl](./hookimpl) - default implementation of hook registry and provider
- [httpserver](./httpserver) - HTTP server contracts, handlers, routing, middleware
- [httpserverstatic](./httpserverstatic) - Static handler interfaces
- [httpserverws](./httpserverws) - WebSocket connection interfaces
- [id](./id) - unique identifier generation and conversion utilities
- [json](./json) - JSON codec wrapper for serialization and streaming
- [meta](./meta) - key-value metadata container
- [plugin](./plugin) - core plugin interfaces
- [runner](./runner) - lifecycle-managed command execution framework
- [utilslice](./utilslice) - generic slice transformation helpers

## 📌 Versioning

Each package inside this repository is independently versioned. This allows for maximum flexibility in plugin
development and backward compatibility control.