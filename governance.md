This document describes the governance rules of the gRPC project(organization). It is meant to be followed by all the repositories in the project and the gRPC community.

## Principles

The gRPC community adheres to the following principles:

* Open: gRPC is open source. See repository guidelines and CLA below.
* Welcoming and respectful: See Code of Conduct below.
* Transparent and accessible: Work and collaboration are done in public.
* Merit: Ideas and contributions are accepted according to their technical merit and alignment with project objectives, scope, and design principles.

## Code Of Conduct

The gRPC community abides by the CNCF [code of conduct](https://github.com/cncf/foundation/blob/master/code-of-conduct.md). Here is an excerpt:

*As contributors and maintainers of this project, and in the interest of fostering an open and welcoming community, we pledge to respect all people who contribute through reporting issues, posting feature requests, updating documentation, submitting pull requests or patches, and other activities.*

The gRPC developers and community are expected to follow the values defined in the [CNCF charter](https://www.cncf.io/about/charter/). As a member of the gRPC project, you represent the project and your fellow contributors. We value our community tremendously and we'd like to keep cultivating a friendly and collaborative environment for our contributors and users. We want everyone in the community to have [positive experiences](https://www.cncf.io/blog/2016/12/14/diversity-scholarship-series-one-software-engineers-unexpected-cloudnativecon-kubecon-experience).


## Decision Making And Voting

gRPC is an open-source [project](https://github.com/grpc/)(organization) with an
open collaboration philosophy. This means that the Github project is the source
of truth for every aspect of the project, including its philosophy, design, road
map, and APIs. *If it's part of the project, it's in the repos. If it's in the
repos, it's part of the project.*

As a result, all decisions can be expressed as changes to the repository. An implementation change is a change to the source code. An API change is a change to the API specification. A philosophy change is a change to the philosophy manifesto, and so on.

All decisions affecting gRPC, big and small, follow the same 3 steps:

* Step 1: Open a pull request. Anyone can do this.
* Step 2: Discuss the pull request. Anyone can do this.
* Step 3: [Maintainers](contributor_ladder.md#maintainer) merge or refuse the pull request.

Exceptions to this are made for:

- the election of the steering committee, the details of which are in this document
- changes to the contributor ladder status of an individual, the requirements for which are outlined in [the contributor ladder](contributor_ladder.md)

In general, we prefer that technical issues are amicably worked out between the persons involved. If a dispute cannot be decided independently, the Maintainers can be called in to resolve the issue by voting. The same PR can be used or a separate PR can be opened in the concerned repository for voting. The title of a PR related to voting should be prefixed with “[vote]”. Such PRs should remain open for a minimum of two weeks unless a decision has been reached sooner. A formal vote on the PR is not required if majority of the maintainers have already agreed in other forums or meetings. In such cases, a detailed comment must be added by a Maintainer before approving or rejecting the PR. In such a case, only an existing Maintainer can ask for a formal vote to challenge the decision. Each Maintainer can cast a maximum of one vote regardless of the number of repositories the maintainer is listed in. A simple majority is required to approve the PR. Only the [Maintainers](contributor_ladder.md#maintainer) associated with the concerned repository (as documented in [maintainers.md](contributors/maintainers.md)) may vote on the PR. For cross-repository issues, the PR can be opened in [grpc-community](https://github.com/grpc/grpc-community) or [proposal](https://github.com/grpc/proposal) repositories. In these cross-repository cases, all [Maintainers](contributor_ladder.md#maintainer) are eligible to vote. The official roster of all Maintainers will be kept up to date in this repo in the [contributors/maintainers.md](contributors/maintainers.md) file. [The Steering Committee](#steering-committee) will have write access to these two repositories, but must use the official voting process outlined above for any changes to them.

## The gRFC Process

We use [gRFCs](https://github.com/grpc/proposal/blob/master/README.md) for any substantial changes to gRPC. This process involves an upfront design that will provide increased visibility to the community. If you're considering a PR that will bring in a new feature across several languages, affect how gRPC is implemented, or may be a breaking change; then you must start with a gRFC. The process is documented in the [proposal](https://github.com/grpc/proposal) repository and a template is provided [here](https://github.com/grpc/proposal/blob/master/GRFC-TEMPLATE.md).

## How To Contribute

See the general guidelines
[here](https://github.com/grpc/grpc-community/blob/master/CONTRIBUTING.md) on
how to contribute to the gRPC project.

## Adding New Repositories

New repositories can be added to gRPC GitHub organization as long as they adhere to the CNCF [charter](https://www.cncf.io/about/charter/) and gRPC governance guidelines. After a project proposal has been announced on a public forum (GitHub issue or mailing list), the existing Maintainers have two weeks to discuss the new project, raise objections and cast their vote. Projects must be approved by a majority vote following the process described above. If a project is approved, a Maintainer will add the project to the gRPC GitHub organization, and make an announcement on a public forum.

## Handling Security Vulnerabilities

The process for handling any security vulnerabilities or concerns found in gRPC is described [here](https://github.com/grpc/proposal/blob/master/P4-grpc-cve-process.md). This process is known as the gRPC CVE (Common Vulnerabilities and Exposure) process.

## Releases

Details on gRPC release process and schedule can be found [here](https://github.com/grpc/grpc/blob/master/doc/grpc_release_schedule.md). Each repository has its own release process. We are in the process of making the release instructions publicly available. Only maintainers can do the releases.

## Changes In Governance

Any change in this document must go through the voting process described above.

## Steering Committee

The following responsibilities and powers belong to the Steering Committee:

* Define, evolve, and defend the vision, mission and the values of the project.
  group.
* Request funds and other support from the CNCF (e.g. marketing, press, etc.)
* Coordinate with the CNCF regarding usage of the gRPC brand and deciding
  project scope, core requirements, and conformance, as well as how that brand
  can be used in relation to other efforts or vendors.

### Committee Meetings

The Steering Committee will meet once per quarter, or as needed.
Meetings are held online, and are public by default.

Given the private nature of some of these discussions (e.g. privacy, private
emails to the committee, code of conduct violations, escalations, disputes
between members, security reports, etc.) some meetings may be held in private.

Meeting notes will be made available to members of the
[grpc-io mailing list](https://groups.google.com/g/grpc-io).
Public meetings will be recorded and the recordings made available publicly.

### Committee members

Seats on the Steering Committee are held by an individual, not by their
employer.

The current membership of the committee is (listed alphabetically by
first name):

| &nbsp;                                                         | Member           | Organization |
| -------------------------------------------------------------- | ---------------- | ------------ |
** TO BE FILLED IN AFTER THE FIRST ELECTION **


### Decision process

The Steering Committee desires to always reach consensus.

Decisions requiring a vote include: issuing written policy, amending existing
written policy, all spending, hiring, and contracting, official responses to
the CNCF, or any other decisions that at least two of the members
present decide require a vote.

Decisions are made in meetings when a quorum of more than half of the members
are present (and all members have been informed of the meeting), and may pass
with more than half the members of the committee supporting it.

### Getting in touch

There are two ways to raise issues to the steering committee for decision:

1. Emailing the Steering Committee at
   [grpc-steering@googlegroups.com](mailto:grpc-steering@googlegroups.com).
   This is a private discussion list to which all members of the committee have
   access.
2. Open an issue on
[grpc/grpc-community](https://github.com/grpc/grpc-community) and indicate that
you would like attention from the steering committee.

### Composition

The steering committee has 7 seats. These seats are
open to anyone.

### Election Procedure

#### Timeline

Steering Committee elections are held annually and terms are a year in length.
Six weeks or more before the election, the Steering Committee will appoint
Election Officer(s) (see below).  Two weeks or more before the election, the
Election Officer(s) will issue a call for nominations. Three days before the
election, the call for nominations will be closed. The election will be open
for voting not less than two weeks and not more than four. The results of the
election will be announced within one week of closing the election. New
Steering Committee members will take office on December 16th each year.

#### Election Officer(s)

Six weeks or more before the election, the Steering Committee will appoint
between one and three Election Officer(s) to administer the election. Elections
Officers will be community members in good standing who are eligible to
vote, are not running for Steering in that election, who are not currently part
of the Steering Committee and can make a public promise of impartiality. They
will be responsible for:

- Making all announcements associated with the election
- Preparing and distributing electronic ballots
- Assisting candidates in preparing and sharing statements
- Tallying voting results according to the rules in this charter

#### Eligibility to Vote

gRPC project [Maintainers](contributor_ladder.md#maintainer), as defined by
[the contributor ladder](contributor_ladder.md), are eligible to vote. The
official list of Maintainers will be kept up to date in this repository in
[contributors/maintainers.md](contributors/maintainers.md).

Anyone holding status as a [Maintainer](contributor_ladder.md#maintainer) at
any any point while the voting is open will be able to cast a vote. The
Election Officer(s) will make available a voter's guide within this repository.

#### Voting Procedure

Elections will be held using a time-limited
[Condorcet](https://en.wikipedia.org/wiki/Condorcet_method) ranking using the
IRV method. Voters will be able to rank the candidates from most to least
preferred, possibly omitting candidates. The top vote-getters will be elected to
the open seats.

### Vacancies

In the event of a resignation or other loss of an elected steering committee
member, the candidate with the next most votes from the previous election will
be offered the seat. This process will continue until the seat is filled.

In case this fails to fill the seat, a special election for that position will
be held as soon as possible, unless the regular steering committee election is
less than 7 weeks away. Eligible voters from the most recent election will vote
in the special election. Eligibility will not be redetermined at the time of the
special election. Any replacement steering committee member will serve out the
remainder of the term for the person they are replacing, regardless of the
length of that remainder.
