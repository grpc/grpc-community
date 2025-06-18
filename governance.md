This document describes the governance rules of the gRPC project(organization). It is meant to be followed by all the repositories in the project and the gRPC community.

## Principles

The gRPC community adheres to the following principles:

* Open: gRPC is open source. See repository guidelines and CLA below.
* Welcoming and respectful: See Code of Conduct below.
* Transparent and accessible: Work and collaboration are done in public.
* Merit: Ideas and contributions are accepted according to their technical merit and alignment with project objectives, scope, and design principles.

## Code Of Conduct

The gRPC community abides by the CNCF [code of conduct](https://github.com/cncf/foundation/blob/main/code-of-conduct.md). Here is an excerpt:

*As contributors and maintainers of this project, and in the interest of fostering an open and welcoming community, we pledge to respect all people who contribute through reporting issues, posting feature requests, updating documentation, submitting pull requests or patches, and other activities.*

The gRPC developers and community are expected to follow the values defined in the [CNCF charter](https://www.cncf.io/about/charter/). As a member of the gRPC project, you represent the project and your fellow contributors. We value our community tremendously and we'd like to keep cultivating a friendly and collaborative environment for our contributors and users. We want everyone in the community to have [positive experiences](https://www.cncf.io/blog/2016/12/14/diversity-scholarship-series-one-software-engineers-unexpected-cloudnativecon-kubecon-experience).

## Decision Making And Voting

gRPC is an open-source [project](https://github.com/grpc/)(organization) with an open collaboration philosophy. This means that the Github project is the source of truth for every aspect of the project, including its philosophy, design, road map, and APIs. *If it's part of the project, it's in the repos. If it's in the repos, it's part of the project.*

As a result, all decisions can be expressed as changes to the repository. An implementation change is a change to the source code. An API change is a change to the API specification. A philosophy change is a change to the philosophy manifesto, and so on.

All decisions affecting gRPC, big and small, follow the same 3 steps:

* Step 1: Open a pull request. Anyone can do this.
* Step 2: Discuss the pull request. Anyone can do this.
* Step 3: Maintainers merge or refuse the pull request.

In general, we prefer that technical issues and maintainer membership are amicably worked out between the persons involved. If a dispute cannot be decided independently, the maintainers can be called in to resolve the issue by voting. The same PR can be used or a separate PR can be opened in the concerned repository for voting. The title of a PR related to voting should be prefixed with “[vote]”. Such PRs should remain open for a minimum of two weeks unless a decision has been reached sooner. A formal voting on the PR is not required if majority of the maintainers have already agreed in other forums or meetings. In such cases, a detailed comment must be added by a maintainer before approving or rejecting the PR. In such a case, only an existing maintainer can ask for a formal vote to challenge the decision. Each maintainer can cast a maximum of one vote regardless of the number of repositories the maintainer is listed in. A simple majority is required to approve the PR. Only the maintainers listed in MAINTAINERS.md file of the concerned repository may vote on the PR. For cross-repository issues, the PR can be opened in [grpc-community](https://github.com/grpc/grpc-community) or [proposal](https://github.com/grpc/proposal) repositories. The list of maintainers in these repositories is a superset of the list of maintainers in all other repositories in the gRPC Github organization. For ease of maintenance only a few senior maintainers will have write access to grpc-community and proposal repositories.

## The gRFC Process

We use [gRFCs](https://github.com/grpc/proposal/blob/master/README.md) for any substantial changes to gRPC. This process involves an upfront design that will provide increased visibility to the community. If you're considering a PR that will bring in a new feature across several languages, affect how gRPC is implemented, or may be a breaking change; then you should start with a gRFC. We've got the process documented in [proposal](https://github.com/grpc/proposal) repository and have a [template](https://github.com/grpc/proposal/blob/master/GRFC-TEMPLATE.md) for you to get started.

## How To Contribute

See the [general guidelines](https://github.com/grpc/grpc-community/blob/main/CONTRIBUTING.md) on how to contribute to gRPC project. If you want to become a maintainer see the section below.

## How To Become A Maintainer

Maintainers (also known as Committers) are first and foremost contributors that have shown they are committed to the long term success of a project. Becoming a maintainer is not required for almost all contributions. In addition to submitting PRs, a maintainer can:
* Triage issues and PRs filed in the repo and assign appropriate labels and reviewers.
* Trigger tests and merge PRs once other checks and criteria pass.
* Cast a vote on issues requiring voting.

Contributors wanting to become maintainers are expected to:
* Be involved in contributing code, pull request review, triage of issues and addressing user questions in one or more forums such as [Github](https://github.com/grpc), [grpcio](https://groups.google.com/forum/#!forum/grpc-io) mailing list, [Stackoverflow](https://stackoverflow.com/search?q=grpc) and [Gitter](https://gitter.im/grpc/grpc).
* Maintain sustained contribution to the gRPC project and spend a reasonable amount of time on it.
* Show deep understanding of the areas contributed to, and good consideration of various reliability, usability, backward compatibility and performance requirements.

These are a few ways to become a maintainer:
* Create a PR adding yourself to MAINTAINERS.md in the appropriate repository. Before doing so it is a good practice to socialize with some existing maintainers to get a good feel for whether you meet the above criteria. A simple majority vote is required to approve the PR. See above for the voting process.
* Current maintainers may nominate a contributor and confer maintainer status. A formal voting on the PR is not required if majority maintainers have already agreed in other forums or meetings.

Please note that in order to be part of the organization, your Github account needs to have [two factor security](https://help.github.com/articles/securing-your-account-with-two-factor-authentication-2fa/) enabled.

## Losing Maintainer Status

If a maintainer is no longer interested or cannot perform the maintainer duties listed above, they should volunteer to be moved to emeritus status. If possible, try to complete your work or help find someone to pick up your work before stepping down. If a maintainer has stopped contributing for a reasonable amount of time, other maintainers may propose to move such maintainers to emeritus list without prior notice. The PR for a such as change would serve as the notice. Such a PR should have @mention of the maintainer in question and should remain open for at least a period of two weeks. Any disagreements will be resolved by a vote of the maintainers per the voting process above.

## Adding New Repositories

Similar to adding maintainers, new repositories can be added to gRPC GitHub organization as long as they adhere to the CNCF [charter](https://www.cncf.io/about/charter/) and gRPC governance guidelines. After a project proposal has been announced on a public forum (GitHub issue or mailing list), the existing maintainers have two weeks to discuss the new project, raise objections and cast their vote. Projects must be approved by a majority vote following the process described above. If a project is approved, a maintainer will add the project to the gRPC GitHub organization, and make an announcement on a public forum.

## Handling Security Vulnerabilities

There is a process for handling any security vulnerabilities or concerns found in gRPC. This process is known as the [gRPC CVE (Common Vulnerabilities and Exposure) process](https://github.com/grpc/proposal/blob/master/P4-grpc-cve-process.md).

## Releases

See the [gRPC release process and schedule](https://github.com/grpc/grpc/blob/master/doc/grpc_release_schedule.md). Each repository has its own release process. We are in the process of making the release instructions publicly available. Only maintainers can do the releases.

## Changes In Governance

Any change in this document must go through the voting process described above.


