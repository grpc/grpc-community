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
- [`contributors/lieutenants.md`](contributors/lieutenants.md)
- [`contributors/maintainers.md`](contributors/maintainers.md)

Contributor Ladder status _may_ also be tracked in individual project repositories (e.g. Golang maintainers in the [grpc-go repo](https://github.com/grpc/grpc-go)), but these listings will not be authoritative.


## Organization Member

[**Current Organization Members**](contributors/organization_members.md)

Description: An Organization Member (or Org Member) is an established contributor who regularly participates in the project. 

#### Delegation to an Organization Member

Individual PR reviews and issues are delegated to Organization Members by
[Lieutenants](#lieutenant) or [Maintainers](#maintainer) explicitly, via
written communication. When a Lieutenant or Maintainer delegates a review or
issue to an Organization Member, it is the responsibility of the delegator to
do initial triage of the PR or issue and ensure that it generally makes sense
and is appropriate for this repository. In addition, the delegator should
provide oversight to the Organization Member, ensuring quality.

In other words, it is not expected that an Organization Member will be the
_sole_ reviewer of a PR, but will be able to help with initial passes.

#### Rights and Responsibilities

An Organization Member:

* Responsibilities include:
    * Continues to contribute regularly, as demonstrated by having at least 4 accepted pull requests or resolved issues per year
* Requirements:
    * Must have successful contributions to the project, including at least one of the following:
        * 4 accepted PRs,
        * Successfully reviewed 4 PRs (as assessed by a maintainer for the relevant repository)
        * Resolved and closed 4 Issues,
        * Become responsible for a key project management area,
        * Or some equivalent combination or contribution
    * Must be actively contributing to at least one project area
    * Must have two sponsors who are also Organization Members or higher

* Privileges:
    * May be delegated Issues and Reviews by a Lieutenant or Maintainer
    * May give commands to CI/CD automation as interpreted by the relevant repository's maintainers
    * Can recommend other contributors to become Org Members


The process for a Contributor to become an Organization Member is as follows:

The official list of Org Members is kept up to date [here](contributors/organization_members.md). Anyone may create a PR to propose the addition of a new Organization Member. With the approval of two existing Organization Members (or higher), the PR must be merged.

### Lieutenant

[**Current Lieutenants**](contributors/lieutenants.md)

Description: A Lieutenant has responsibility for specific code, documentation, test, or other project areas. They are collectively responsible, with other Lieutenants and Maintainers, for reviewing all changes to those areas and indicating whether those changes are ready to merge. They have a track record of contribution and review in the project.

The term "Lieutenant" comes from the Linux Kernel project, where Linus Torvalds
delegates responsibility of specific domains to his "Lieutenants," while
exercising overall control himself. Similarly, [Maintainers](#maintainers)
maintain overall control for a gRPC subproject (such as an implementation in a
particular language), but delegate authority for a specific subdomain to one or
more gRPC Project Lieutenants.

Lieutenants are responsible for a "specific area." This can be a specific repository, code directory, chapter of the docs, test job, event, or other clearly-defined project. The "specific area" below refers to this area of responsibility.

#### Delegation to a Lieutenant

Individual PR reviews and issues are delegated to Lieutenants by Maintainers
either implicitly (via prior verbal agreement on repository policy) or
explicitly, via written communication. When a Maintainer delegates a review or
issue to a Lieutenant, it is the responsibility of that Maintainer to do initial
triage of the PR or issue and ensure that it generally makes sense and is
appropriate for this repository, and that the Lieutenant to which it is
delegated is the appropriate person to look into the details.

Put informally, the delegating Maintainer must still provide "approval" but not
necessarily "LGTM," which is left to the Lieutenant.

#### Rights and Responsibilities

Lieutenants have all the rights and responsibilities of an Organization Member, plus:

* Responsibilities include:
    * Reviewing most Pull Requests against their specific areas of responsibility
    * Helping other contributors become Lieutenants
* Requirements:
    * Is an Organization Member
    * Has demonstrated an in-depth knowledge of the specific area
    * Commits to being responsible for that specific area
    * Is supportive of new and occasional contributors and helps get useful PRs in shape to commit
    * Must have two sponsors who are Lieutenants or higher within the relevant repo(s)
* Additional privileges:
    * Has permissions to approve pull requests relating to their area of expertise
    * Can recommend and review other contributors to become Lieutenants
    
The process of becoming a Reviewer is as follows:

The official list of Org Members is kept up to date [here](contributors/lieutenants.md). Anyone may create a PR to propose the addition of a new Lieutenant, specifically naming their area of expertise and which repositories it applies to. With the approval of two existing Lieutenants (or higher), the PR must be merged.

### Maintainer

[**Current Maintainers**](contributors/maintainers.md)

Description: Maintainers are very established contributors who are collectively
responsible for the entire project and personally responsible for large
subprojects. As such, they have broad voting powers, sole authority over their
area of expertise, and are expected to participate in making decisions about the
strategy and priorities of the project.

A Maintainer must meet the responsibilities and requirements of a Lieutenant, plus:

* Responsibilities include:
    * Mentoring new Lieutenants and Organization Members
    * Writing refactoring PRs
    * Participating in CNCF maintainer activities
    * Participating in, and leading, community meetings
* Requirements
    * Demonstrates a broad knowledge of the project across multiple areas
    * Is able to exercise judgment for the good of the project, independent of their employer, friends, or team
    * Mentors other contributors
    * Can commit to spending at least 10 hours per month working on the project
* Additional privileges:
    * Approve PRs in their area of ownership with no oversight
    * Represent the project in public as a Maintainer
    * Communicate with the CNCF on behalf of the project
    * Vote in Steering Committee Elections

#### Collaboration between Maintainers

Maintainers have broad authority over their area of expertise, as well as a
weighty amount of influence through their voting powers. However, many changes
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
1. Any current Maintainer may nominate a current Reviewer to become a new Maintainer, by opening a PR to [maintainers.md](contributors/maintainers.md).
2. The nominee will add a comment to the PR testifying that they agree to all requirements of becoming a Maintainer.
3. A majority of the current Maintainers (as listed in [maintainers.md](contributors/maintainers.md)) must then approve the PR.


## Inactivity
It is important for contributors to be and stay active to set an example and show commitment to the project. Inactivity is harmful to the project as it may lead to unexpected delays, contributor attrition, and a loss of trust in the project.

* Inactivity is measured by:
    * Periods of no contributions for longer than 6 months, including Github comments, PRs, and issues.
    * Periods of no communication for longer than 6 months, including emails, messages, and attendance at project meetings.
* Consequences of being inactive include:
    * Involuntary removal or demotion
    * Being asked to move to Emeritus status

## Involuntary Removal or Demotion

Involuntary removal/demotion of a contributor happens when responsibilities and requirements aren't being met. This may include repeated patterns of inactivity, extended period of inactivity, a period of failing to meet the requirements of your role, and/or a violation of the Code of Conduct. This process is important because it protects the community and its deliverables while also opens up opportunities for new contributors to step in.

Involuntary removal or demotion is handled through a vote by a majority of the current Maintainers.

## Stepping Down/Emeritus Process
If and when contributors' commitment levels change, contributors can consider stepping down (moving down the contributor ladder) vs moving to emeritus status (completely stepping away from the project).

Contact the Maintainers about changing to Emeritus status, or reducing your contributor level.

## Contact
For inquiries, please reach out to [grpc-steering@googlegroups.com](mailto:grpc-steering@googlegroups.com).
