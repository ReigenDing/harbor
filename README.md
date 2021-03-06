# Harbor

Project Harbor is an enterprise-class registry server.It extends the open source Docker Registry server by adding more functionalities usually required by an enterprice.Harbor os designed to be deployed in a private environment of an organization. A private registry id important for organization who care much about security. In addition, a private registry improves priductivity by elliminating the need to download images from public network. This is very helpful to container users who do not have a good network to the Internet. In particular, Harbor accelerates the progress of Chinese developers, because they no logger need to pull images from the Internet.

### Features
* **Role Based Access Control**: Users and docker repositories are organized via "projects", a user can have differernt permission for images under a namespace.
* ** Graphical user portal **: User can easily browse, search docker repositories, manage projects/namespaces.
* ** AD/LDAP support **: Harbor integrates with existing AD/LDAP of enterprise for user authentication and management.
* ** Audting **: All the operatiorns to the repositories are tracked and can be used for auditing purpose.
* ** Internationalization **: Localized for English and Chinese languages. More languages can be added.
* ** RESTful API **: RESTful APIs are provided for most administrative operations of Harbor. The integration with other management sofrwate becomes easy.

### Try it
Harbor is self contained and can be easily deployed via docker-compose.
```sh
$ cd Deploy
#make update to the parameters in ./harbor.cfg
$ ./prepare
Generated configuration file: ./config/ui/env
Generated configuration file: ./config/ui/app.conf
Generated configuration file: ./config/registry/config.yml
Generated configuration file: ./config/db/env
$ docker-compose up
```

### License
Harbor is available under the [Apache 2 license](LICENSE).

