![commit stage](https://github.com/thefuga/go-poc/actions/workflows/commit-stage.yml/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/thefuga/go-poc)](https://goreportcard.com/report/github.com/thefuga/go-poc)
------

# Go POCs

This repository holds proof of concepts related to Go web applications. All concepts applied will be documented here. The applied concepts will be documented on the application itself, using go docs.

## Architecture overview

The architecture presented here is aimed toward the [hexagonal architecture](https://alistair.cockburn.us/hexagonal-architecture/). To achieve it, interfaces are used as much as possible to isolate every external resource from the domain. This leads to a lot of indirection, which is resolved by a dependency injection container. [FX](https://github.com/uber-go/fx/) was chosen to fulfil this purpose.

### Dependency injection

As mentioned earlier, [FX](https://github.com/uber-go/fx/) is used to resolve and inject all dependencies in this application. FX is built with two basic concepts: construction and invocation.

This application delegates the responsibility of defining all constructors and invokables from a package itself. Each package that requires DI has a module, declaring the constructors and invokables. Outer packages will import inner packages and compose their own module with the imported modules. This is propagated up to the main application modules. The main application module is then used to run the application.

#### Constructors

Constructors are used to building a given dependency. FX uses registered constructors to resolve the dependencies from other constructors.

#### Invokables

Invokables are functions called after all dependencies are resolved. FX calls these with the resulting dependencies. They are used to kickstart the application and perform actions that are typically executed once in the application lifetime. Their use includes, but is not limited to opening database connections, registering HTTP routes.

#### Interface resolution

One of the most common patterns in go is accepting interfaces and return structs. This, along with the client-side interface definition, has the power to make testing extremely simple. The downside is that it creates a lot of indirection.

FX allows us to resolve this indirection by mapping interfaces to implementations. Because this application is composed of modules that compose other modules, resolving this mapping is pretty straightforward and should be done ate the module importing both client and provider.

With this approach, adapters are completely isolated and are consumed - through the ports - by the use cases. 

See the whole user module for complete examples.

## Configs

Configs are managed by [Viper](https://github.com/spf13/viper). They are loaded by FX at the application's startup. The Viper struct is also registered in the container and can be used as a service if needed.

The application loads and registers configs as any other dependency: in the FX lifecycle. As with other modules, the configs module can be nested and combined to form the main configs module. See configs package for details.
