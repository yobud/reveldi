# Reveldi - A Revel Dependency Injection module

This library allows you to work with services in Revel framework. This is a simple implementation of a dependency injection container, hope you'll like it.
Extra simple stuff, just a map that you can fill with your own Structs.
In facts, this library isn't Revel exclusive, it would work on any go project

## Installation

> go get github.com/zebigduck/reveldi

## Usage

To use it, you need to create a new controller

```go

    // app/controllers/dic.go

    package controllers

    import "github.com/revel/revel"
    import "github.com/zebigduck/reveldi"
    import "app/app/modules/message"

    type Dic struct {
        container *reveldi.Container
    }

    func (c *Dic) Init() {
        c.container = new(reveldi.Container)

        // Example : Messenger service

            // Register a simple instanciated Struct as a service
            // c.Container.Register("messenger", new(message.Messenger)) // Hi Jérémy!

            // You can easily change the Struct to use
            // c.Container.Register("messenger", new(message.SuperMessenger)) // super Hi Jérémy!

            // Create instance and configure service
            messenger := new(message.PrefixMessenger)
            // messenger.SetPrefix("Hi")

            // Or use config.conf parameter
            messenger.SetPrefix(revel.Config.StringDefault("messenger.prefix", "Hey"))

            // Then register the Service to the container
            c.container.Register("messenger", messenger)

        // Now you can use it as follow anywhere in your controllers
        // myMessenger := Container.Get("messenger").(message.MessengerInterface)
    }

    // This is a shortcut to Get Container's func
    func (c *Dic) Get(name string) reveldi.Service {
        return c.container.Get(name)
    }

```

And then, you just have to create global variable and init Dic in init.go

```go

    package controllers

    import (
        "github.com/revel/revel"
    )

    // Create the Container global variable
    var Container *Dic = new(Dic)

    func init() {
        // Init the container with your services
        revel.OnAppStart(func () { Container.Init() })
    }

```

Now, you're able to abstract your managers, just use them with the following instruction in your controllers:

> myMessenger := Container.Get("messenger").(message.MessengerInterface)