# Contributor Ladder

Hello! We are excited that you want to learn more about our project contributor
ladder! This contributor ladder outlines the different contributor roles within
the project, along with the responsibilities and privileges that come with them.
Community members generally start at the first levels of the ladder and advance
up it as their involvement in the project grows.  Our project members are happy
to help you advance along the contributor ladder.

Each of the contributor roles below is organized into lists of three types of
things. "Responsibilities" are things that a contributor is expected to do.
"Requirements" are qualifications a person needs to meet to be in that role, and
"Privileges" are things contributors on that level are entitled to.

While privileges are outlined here, the exact mapping of these privileges to
physical permissions within a repository is left to the maintainers for that
repository, as each repository has different code review and CI/CD workflows
with different considerations.

The authoritative source for contributor status will be the following files housed in this repository:

- [`contributors/organization_members.md`](contributors/organization_members.md)
- [`contributors/core_contributors.md`](contributors/core_contributors.md)
- [`contributors/maintainers.md`](contributors/maintainers.md)

Contributor Ladder status _may_ also be tracked in individual project
repositories (e.g. Golang maintainers in the [grpc-go
repo](https://github.com/grpc/grpc-go)), but these listings will not be
authoritative.


## Organization Member

[**Current Organization Members**](contributors/organization_members.md)

Description: An Organization Member (or Org Member) is an established
contributor who regularly participates in the project. 

#### Delegation to an Organization Member

Individual PR reviews and issues are delegated to Organization Members by
[Core Contributors](#core-contributor) or [Maintainers](#maintainer) explicitly, via
written communication. When a Core Contributor or Maintainer delegates a review or
issue to an Organization Member, it is the responsibility of the delegator to
do initial triage of the PR or issue and ensure that it generally makes sense
and is appropriate for this repository. In addition, the delegator should
provide oversight to the Organization Member, ensuring quality.

In other words, it is not expected that an Organization Member will be the
_sole_ reviewer of a PR, but will be able to help with initial passes.

#### Rights and Responsibilities

An Organization Member:

* Responsibilities:
    * Continues to contribute regularly, as demonstrated by having at least 4 accepted pull requests or resolved issues per year
* Requirements to become an Organization Member:
    * Must have successful contributions to the project, including at least one of the following:
        * 4 accepted PRs,
        * Successfully reviewed 4 PRs (as assessed by a maintainer for the relevant repository)
        * Resolved and closed 4 Issues
        * Become responsible for a key project management area,
        * Or some equivalent combination or contribution
    * Must be actively contributing to at least one project area
    * Must have two sponsors who are also Organization Members or higher, who indicate their support via the PR process outlined below.

* Privileges:
    * May be delegated Issues and Reviews by a Core Contributor or Maintainer
    * May give commands to CI/CD automation as interpreted by the relevant repository's maintainers
    * Can recommend other contributors to become Org Members


The process for a contributor to become an Organization Member is as follows:

The official list of Org Members is kept up to date
[here](contributors/organization_members.md). Anyone may create a PR to propose
the addition of a new Organization Member. The PR must have the approval of two
existing Organization Members (or higher). If the author of the PR is not the
nominee, then the nominee must comment on the PR indicating their support for
the new role and its responsibilities. If all of the above conditions are met,
the PR must be merged.

### Core Contributor

[**Current Core Contributors**](contributors/core_contributors.md)

Description: A Core Contributor has responsibility for specific code, documentation,
test, or other project areas. They are collectively responsible, with other
Core Contributors and Maintainers, for reviewing all changes to those areas and
indicating whether those changes are ready to merge. They have a track record
of contribution and review in the project.

[Maintainers](#maintainers)
maintain overall control for a gRPC subproject (such as an implementation in a
particular language), but delegate authority for a specific subdomain to one or
more gRPC Project Core Contributors.

Core Contributors are responsible for a "specific area." This can be a specific
repository, code directory, chapter of the docs, test job, event, or other
clearly-defined project. The "specific area" below refers to this area of
responsibility.

#### Delegation to a Core Contributor

Individual PR reviews and issues are delegated to Core Contributors by Maintainers
either implicitly (via prior verbal agreement on repository policy) or
explicitly, via written communication. When a Maintainer delegates a review or
issue to a Core Contributor, it is the responsibility of that Maintainer to do initial
triage of the PR or issue and ensure that it generally makes sense and is
appropriate for this repository, and that the Core Contributor to which it is
delegated is the appropriate person to look into the details.

Put informally, the delegating Maintainer must still provide "approval" but not
necessarily "LGTM," which is left to the Core Contributor.

#### Rights and Responsibilities

Core Contributors have all the rights and responsibilities of an Organization Member, plus:

* Responsibilities:
    * Reviewing most Pull Requests against their specific areas of responsibility
    * Helping other contributors become Core Contributors
* Requirements to become a Core Contributor:
    * Is already an Organization Member
    * Has demonstrated an in-depth knowledge of the specific area
    * Commits to being responsible for that specific area
    * Is supportive of new and occasional contributors and helps get useful PRs in shape to commit
    * Must have two sponsors who are Core Contributors or higher within the relevant repo(s), who indicate their support via the PR process outlined below.
* Additional privileges:
    * Has permissions to approve pull requests relating to their area of expertise
    * Can recommend and review other contributors to become Core Contributors
    
The process of becoming a Core Contributor is as follows:

The official list of Core Contributors is kept up to date
[here](contributors/core_contributors.md). Anyone may create a PR to propose the
addition of a new Core Contributor, specifically naming their area of expertise and
which repositories it applies to. The approval of two existing Core Contributors
(or higher) is required for merge. These two approvers must have standing
within the relevant repository listed in the PR. If the author of the PR is not
the nominee, then the nominee must comment on the PR indicating their support
for the new role and its responsibilities. If all of the above conditions are
met, then the PR must be approved.

### Maintainer

[**Current Maintainers**](contributors/maintainers.md)

Description: Maintainers are very established contributors who are collectively
responsible for the entire project and personally responsible for large
subprojects. As such, they have broad voting powers, sole authority over their
area of expertise, and are expected to participate in decision-making about the
strategy and priorities of the project.

A Maintainer must meet the responsibilities and requirements of a Core Contributor, plus:

* Responsibilities include:
    * Mentoring new Core Contributors and Organization Members
    * Writing refactoring PRs
    * Participating in CNCF maintainer activities
    * Participating in, and leading, community meetings
* Requirements to become a Maintainer:
    * Is already a Core Contributor.
    * Demonstrates a broad knowledge of the project across multiple areas
    * Is able to exercise judgment for the good of the project, independent of their employer, friends, or team
    * Mentors other contributors
    * Can commit to spending at least 10 hours per month working on the project
    * Majority vote of existing Maintainers.
* Additional privileges:
    * Approve PRs in their area of ownership with no oversight
    * Represent the project in public as a Maintainer
    * Communicate with the CNCF on behalf of the project
    * Vote in Steering Committee Elections

#### Collaboration between Maintainers

Maintainers have broad authority over their area of expertise, as well as
weighty influence in their voting powers. However, many changes
will span multiple areas of expertise. While the [primary workflow for driving
changes in the gRPC project requires the approval of only a single
maintainer](governance.md#decision-making-and-voting), some changes will span
_multiple_ areas of expertise. This is especially true for
[gRFCs](governance.md#the-grfc-process). In this case, _one_ maintainer acts as
the triager and is responsible for adding a maintainer to the review to
represent all relevant areas of expertise. For example, a new load balancing
policy proposed in a gRFC must be implementable in all supported languages, so a
Maintainer for C++, Go, Java, etc. should be added to the review by the triaging
Maintainer.

Process of becoming a maintainer:
1. Any current Maintainer may nominate a current Core Contributor to become a new Maintainer, by opening a PR to [maintainers.md](contributors/maintainers.md).
2. The nominee will add a comment to the PR testifying that they agree to all requirements of becoming a Maintainer.
3. A majority of the current Maintainers (as listed in [maintainers.md](contributors/maintainers.md)) must then approve the PR or give verbal approval in a meeting.


## Inactivity
It is important for contributors to remain active to set an example and show commitment to the project. Inactivity is harmful to the project as it may lead to unexpected delays, contributor attrition, and a loss of trust in the project.

* Inactivity is measured by:
    * Periods of no contributions for longer than 6 months, including Github comments, PRs, and issues.
    * Periods of no communication for longer than 6 months, including emails, messages, and attendance at project meetings.
* Consequences of being inactive include:
    * Involuntary removal or demotion
    * Being asked to move to Emeritus status

## Involuntary Removal or Demotion

Involuntary removal/demotion of a contributor happens when responsibilities and requirements aren't being met. This may include repeated patterns of inactivity, extended period of inactivity, a period of failing to meet the requirements of your role, and/or a violation of the Code of Conduct. This process is important because it protects the community and its deliverables while also opening up opportunities for new contributors to step in.

Involuntary removal or demotion is handled through a vote by a majority of the current Maintainers.

## Stepping Down/Emeritus Process
If and when contributors' commitment levels change, contributors can consider stepping down (moving down the contributor ladder) vs moving to emeritus status (completely stepping away from the project).

Contact the Maintainers about changing to Emeritus status, or reducing your contributor level.

Emeritus status is reflected by moving a contributor to the Emeritus table at the bottom of the official list for each ladder rung.

## Contact
For inquiries, please reach out to [grpc-steering@googlegroups.com](mailto:grpc-steering@googlegroups.com).
