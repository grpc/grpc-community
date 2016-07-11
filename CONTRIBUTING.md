# [gRPC-ecosystem](https://github.com/grpc-ecosystem) is the organization under which we are collecting community contributions around gRPC.

# How to contribute

This is a place for various components in the gRPC ecosystem that aren't part of the gRPC core. We welcome contributions in this repo which either build extensions around gRPC or showcase how to use gRPC in various use cases and/or with other technologies.
Here is some guideline and information about how to do so


## Getting started

### Legal requirements

In order to protect both you and ourselves, you will need to sign the
[Contributor License Agreement](https://cla.developers.google.com/clas).Also, no third party libraries which are under AGPL etc should not be used.

### Guidelines to contribute

gRPC team will control which repos are curated into grpc-ecosystem org and their decision will be final. Contributors will be required to sign [Google CLA](https://cla.developers.google.com/clas) and all license, legal and patent rights will be determined by CLA and [Apache license](http://www.apache.org/licenses/LICENSE-2.0). Work on these will be done in a public manner and each contributing team will have full admin control of their repos. 

Each contribution needs to have a a top level readme explaining what the contribution does, how to use it with gRPC, how to build and test it and what are its external technical dependencies. So each repository should have

- Have at least a top level readme.md describing overview, how to use, dependencies, and how to build and test.
- Third party libraries: Note that no third party libraries with AGPL license etc should not be used in the codebases.
- Automated tests - will have a badge called “Verified” for tested contributions. Contributors should have automated tests present in every contribution and they should run on commit. We (gRPC team) will set up travis CI to facilitate this. Tests must return green before we merge them.


### How contributions will be accepted?

Anyone who wants to contribute a new repo in grpc-ecosystem should fill up this [form](https://docs.google.com/a/google.com/forms/d/119zb79XRovQYafE9XKjz9sstwynCWcMpoJwHgZJvK74/edit). Once gRPC team approves, we will either create a new repo (with license and patents file) for new projects or enable you to move your repo into grpc-ecosystem org. (we do this by making you admin to to grpc-ecosystem-admins org and enabling transfer of repo . Once repo has moved to gRPC ecosystem admins, gRPC team can move it to gRPC-ecosystem and give contributor full admin rights for subsequent control. Code reviews will be done on a best effort basis. It is however expected that the community will address the comments from core team members.

