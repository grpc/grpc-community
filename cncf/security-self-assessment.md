# Self-assessment

The Self-assessment is the initial document for projects to begin thinking about the
security of the project, determining gaps in their security, and preparing any security
documentation for their users. This document is ideal for projects currently in the
CNCF **sandbox** as well as projects that are looking to receive a joint assessment and
currently in CNCF **incubation**.

For a detailed guide with step-by-step discussion and examples, check out the free
Express Learning course provided by Linux Foundation Training & Certification:
[Security Assessments for Open Source Projects](https://training.linuxfoundation.org/express-learning/security-self-assessments-for-open-source-projects-lfel1005/).

## Self-assessment outline

### Table of contents

* [Metadata](#metadata)
  * [Security links](#security-links)
* [Overview](#overview)
  * [Actors](#actors)
  * [Actions](#actions)
  * [Background](#background)
  * [Goals](#goals)
  * [Non-goals](#non-goals)
* [Self-assessment use](#self-assessment-use)
* [Security functions and features](#security-functions-and-features)
* [Project compliance](#project-compliance)
* [Secure development practices](#secure-development-practices)
* [Security issue resolution](#security-issue-resolution)
* [Appendix](#appendix)

### Metadata

A table at the top for quick reference information, later used for indexing.

|||
| -- | -- |
| Assessment Stage | Complete |
| Software | https://github.com/grpc |
| Security Provider | No |
| Languages | C++, Java, Go, Python, Ruby, C#, Objective-C, PHP, Node.js, Dart |
| SBOM | gRPC follows the SBOM generation best practices for each of its supported languages. For example, for Go, a user can generate an SBOM using `syft` by running `syft . -o cyclonedx-json > sbom.json` in their project's root. |

#### Security links

| Doc | url |
| -- | -- |
| gRPC CVE Process | https://github.com/grpc/proposal/blob/master/P4-grpc-cve-process.md |

### Overview

gRPC is a modern, open-source, high-performance Remote Procedure Call (RPC) framework that can run in any environment. It efficiently connects services in and across data centers with pluggable support for load balancing, tracing, health checking, and authentication.

#### Background

gRPC is built on top of HTTP/2, which provides features like bidirectional streaming, flow control, header compression, and multiplexing requests over a single TCP connection. It uses Protocol Buffers as its default Interface Definition Language (IDL) for serializing structured data. This combination makes gRPC highly efficient and performant, making it a popular choice for microservices communication, real-time streaming, and connecting mobile and IoT devices to backend services.

#### Actors

The primary actors in a gRPC system are:

*   **gRPC Client:** Initiates the RPC call. It uses a generated stub to call the remote method as if it were a local function.
*   **gRPC Server:** Listens for incoming RPC calls and implements the service interface. It processes the client's request and sends back a response.
*   **Proxy/Load Balancer:** (Optional) Sits between clients and servers to provide features like load balancing, routing, and protocol translation (e.g., gRPC-Web proxy).
*   **xDS Control Plane:** (Optional) In a service mesh environment, the control plane provides configuration to gRPC clients and servers for features like service discovery, load balancing, routing, and security.

Isolation between these actors is typically achieved through network boundaries (e.g., different machines or containers) and secured using transport-level security (TLS) and authentication mechanisms.

#### Actions

A typical gRPC unary RPC call involves the following actions:

1.  The client creates a "channel" to the server, specifying the server address and credentials.
2.  The client calls a method on the generated stub, passing a request message.
3.  The gRPC client library serializes the request message using Protocol Buffers and sends it to the server over an HTTP/2 stream.
4.  The gRPC server receives the request, deserializes the message, and calls the corresponding service implementation method.
5.  The service implementation processes the request and returns a response message.
6.  The gRPC server library serializes the response message and sends it back to the client over the same HTTP/2 stream.
7.  The client receives the response, deserializes it, and returns it to the application code.

For streaming RPCs, the client and server can send multiple messages over the same stream.

#### Goals

The primary goals of gRPC are:

*   To provide a high-performance, low-latency RPC framework.
*   To support a wide range of programming languages and platforms.
*   To enable efficient communication between services in a distributed system.
*   To provide built-in support for authentication, load balancing, and other essential features.
*   To offer a secure-by-default communication channel through TLS.

#### Non-goals

*   gRPC is not a service mesh, but it can be a core component of one.
*   gRPC does not provide a built-in solution for service discovery, but it can integrate with existing solutions like DNS, Consul, or an xDS control plane.
*   gRPC does not aim to be a general-purpose HTTP client/server library. Its focus is on RPC.

### Self-assessment use

This self-assessment is created by the gRPC team to perform an internal analysis of the project's security. It is not intended to provide a security audit of gRPC, or function as an independent assessment or attestation of gRPC's security health.

This document serves to provide gRPC users with an initial understanding of gRPC's security, where to find existing security documentation, gRPC plans for security, and general overview of gRPC security practices, both for development of gRPC as well as security of gRPC.

This document provides the CNCF TAG-Security with an initial understanding of gRPC to assist in a joint-assessment, necessary for projects under incubation. Taken together, this document and the joint-assessment serve as a cornerstone for if and when gRPC seeks graduation and is preparing for a security audit.

### Security functions and features

*   **Critical:**
    *   **Transport Security (TLS):** gRPC provides built-in support for TLS to encrypt communication between clients and servers. It is the recommended and default way to use gRPC securely.
    *   **Authentication:** gRPC supports various authentication mechanisms, including SSL/TLS with client certificates and token-based authentication (e.g., JWT, OAuth 2.0).
    *   **HTTP/2 Transport:** gRPC's use of HTTP/2 provides a secure and efficient transport layer, with features like stream multiplexing and header compression.

*   **Security Relevant:**
    *   **Authorization:** gRPC provides interceptors that can be used to implement authorization logic, such as Role-Based Access Control (RBAC).
    *   **Pluggable Credentials:** gRPC's credentials mechanism is extensible, allowing users to integrate custom authentication systems.
    *   **xDS Integration:** gRPC's integration with xDS control planes enables centralized management of security policies, including mTLS configuration.

### Project compliance

While gRPC itself is not certified against any specific compliance standards, it provides the necessary security features (e.g., encryption in transit, authentication, audit logging) that enable users to build applications that meet the requirements of standards like PCI-DSS, HIPAA, and GDPR.

### Secure development practices

*   **Development Pipeline:**
    *   All code changes are submitted via Pull Requests and require review and approval from project maintainers.
    *   We employ continuous fuzzing tests and use sanitizers (ASan, TSan, etc.) in our CI to detect memory-related bugs and other vulnerabilities.
    *   Container images are not a primary distribution method for the core gRPC libraries.
*   **Communication Channels:**
    *   **Internal:** Team members communicate via [the grpc-io mailing list](https://groups.google.com/g/grpc-io), GitHub issues/PRs, and video calls.
    *   **Inbound:** Users can communicate with the team through GitHub issues and [the grpc-io mailing list](https://groups.google.com/g/grpc-io).
    *   **Outbound:** We communicate with users through [the grpc-io mailing list](https://groups.google.com/g/grpc-io), [the gRPC blog](https://grpc.io/blog/), release notes on GitHub, and the [grpc-io-announce](https://grpc.io/blog/grpc-io-announce/) mailing list.
*   **Ecosystem:**
    gRPC is a foundational component of the cloud-native ecosystem. It is used by numerous CNCF projects, including Kubernetes (for CRI and CSI), etcd, and Vitess. It is also the default RPC framework for many service meshes like Istio and Linkerd.

### Security issue resolution

*   **Responsible Disclosures Process:**
    *   We follow a responsible disclosure process for security vulnerabilities.
    *   Vulnerabilities should be reported privately to `grpc-security@googlegroups.com`.
    *   The full process is documented in the [gRPC CVE Process](https://github.com/grpc/proposal/blob/master/P4-grpc-cve-process.md).
*   **Incident Response:**
    *   The gRPC security team is responsible for responding to vulnerability reports.
    *   The security team acknowledges and analyzes reports within three working days.
    *   A fix is developed and backported to all supported release branches.
    *   A CVE is requested from MITRE.
    *   A public disclosure date is negotiated with the reporter, and the vulnerability is announced on the grpc-io mailing list and grpc-io-announce mailing list.

### Appendix

*   **Security Audit Summary:** In 2020, the CNCF sponsored a third-party security audit of gRPC's C-core implementation, performed by Cure53. The audit identified 13 issues in total: 1 critical, 2 high, 3 medium, 5 low, and 2 informational. The critical issue related to a potential memory leak in the HTTP/2 transport. All identified vulnerabilities were promptly fixed by the gRPC team and the fixes were verified by the auditors. The full report is available [here](https://github.com/grpc/grpc/blob/master/doc/grpc_security_audit.pdf).
[*   **Known Issues Over Time:** A list of past security vulnerabilities can be found in the security advisories on GitHub for each language-specific repository (e.g., [gRPC-Core (C++/Python/etc)](https://github.com/grpc/grpc/security/advisories), [gRPC-Go](https://github.com/grpc/grpc-go/security/advisories), [gRPC-Java](https://github.com/grpc/grpc-java/security/advisories)) or with [NIST CVE search](ihttps://nvd.nist.gov/vuln/search#/nvd/home?cpeFilterMode=cpe&cpeName=cpe:2.3:a:grpc:grpc:1.0.0:*:*:*:*:java:*:*&resultType=records). An analysis of these advisories shows that the most common vulnerability class is Denial of Service (DoS), often related to resource exhaustion in the HTTP/2 transport layer. The project has a strong track record of fixing and disclosing these issues in a timely manner.
*   **Case Studies:**
    *   **Netflix:** Netflix uses gRPC for its internal microservices communication, handling millions of RPCs per second.
    *   **Square:** Square uses gRPC for communication between its backend services, enabling them to build a scalable and resilient platform.
    *   **OpenAI:** OpenAI uses gRPC for its API, allowing developers to integrate powerful AI models into their applications.
*   **Related Projects / Vendors:**
    *   **REST:** REST is a popular architectural style for building APIs, but it can be less performant than gRPC for internal microservices communication.
    *   **Thrift:** Thrift is another RPC framework developed by Facebook. It is similar to gRPC but has a different set of features and trade-offs.
    *   **Service Meshes (Istio, Linkerd):** Service meshes provide a layer of infrastructure for managing and securing microservices communication. They often use gRPC as the underlying RPC framework.
