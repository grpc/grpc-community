If you are interested in contributing to the gRPC project, below are some general guidelines that apply to all the repositories in gRPC Github organization. Please read the additional instructions in CONTRIBUTING.md file of the concerned repository.

## Before You Contribute

We welcome and encourage contributions from the community. Please read the gRPC project [governance](https://github.com/grpc/grpc-community/blob/main/governance.md) document. Contributions should be made via pull requests. Pull requests will be reviewed by one or more maintainers and merged when acceptable. Please read the following guidelines carefully to maximize the chances of your PR being merged.

For contributions not related to gRPC design and implementation but useful to gRPC community, please see gRPC [ecosystem guidelines](https://github.com/grpc/grpc-community/blob/main/grpc_ecosystem.md).

## Legal Requirements
In order to protect both you and the gRPC project, you will need to sign the CNCF [Contributor License Agreement](https://identity.linuxfoundation.org/projects/cncf) before your PR can be merged.

## Communication

Trivial changes and small bug fixes do not need prior communication. You can just submit a PR with minimum details. For larger PRs, we ask that before contributing, please make the effort to coordinate with the maintainers of the project via a Github issue or via [grpcio](https://groups.google.com/forum/#!forum/grpc-io) mailing list. This will prevent you from doing extra or redundant work that may or may not be merged.

## Have Questions?
It is best to ask questions on forums other than Github repos. Github repos are for filing issues and submitting PRs. You have a higher chance of getting help from the community if you ask your questions on [grpcio](https://groups.google.com/forum/#!forum/grpc-io) mailing list or [Stackoverflow](https://stackoverflow.com/).

## Guidelines For Pull Requests

How to get your contributions merged smoothly and quickly.
* Create smaller PRs that are narrowly focused on addressing a single concern. We often times receive PRs that are trying to fix several things at a time, making the review process difficult. Create more PRs to address different concerns for faster resolution.
* Make sure to add new tests for bugs in order to catch regressions and to test any newly added functionality.
* For speculative changes, consider opening an issue and discussing it first. If you are suggesting a behavioral or API change, consider starting with a [gRFC proposal](https://github.com/grpc/proposal).
* Provide a good PR description as a record of what change is being made and why it was made. Link to a GitHub issue if it exists.
* Don't fix code style and formatting unless you are already changing that line to address an issue. PRs with irrelevant changes won't be merged. If you do want to fix formatting or style, do that in a separate PR.
* Unless your PR is trivial, you should expect there will be reviewer comments that you'll need to address before merging. We expect you to be reasonably responsive to those comments; otherwise, the PR will be closed after 2-3 weeks of inactivity.
* If you have non-trivial contributions, please consider adding an entry to the AUTHORS.md file in the concerned repository listing the copyright holder for the contribution (yourself, if you are signing the individual CLA, or your company, for corporate CLAs) in the same PR as your contribution. This needs to be done only once for each company or individual. Please keep this file in alphabetical order.
* Maintain clean commit history and use meaningful commit messages. PRs with messy commit history are difficult to review and won't be merged. Use rebase -i upstream/master to curate your commit history and/or to bring in latest changes from master but avoid rebasing in the middle of a code review.
* Keep your PR up to date with upstream/master. If there are merge conflicts, we can't really merge your change.

Above are general guidelines that apply to all the repositories in gRPC Github organization. Please read the additional instructions in CONTRIBUTING.md file of the concerned repository.

