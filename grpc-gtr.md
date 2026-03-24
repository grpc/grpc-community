# General Technical Review - gRPC

  - **Project:** gRPC
  - **Project Version:** v1.76.0 (as of Oct 17th, 2025)
  - **Website:** [https://grpc.io/](https://grpc.io/)
  - **Date Updated:** 2025-12-26
  - **Template Version:** v1.0
  - **Description:** gRPC is a modern, open-source, high-performance Remote Procedure Call (RPC) framework that can run
    in any environment. It efficiently connects services in and across data centers with pluggable support for load
    balancing, tracing, health checking, and authentication. It is also applicable in the last mile of distributed
    computing to connect devices, mobile applications and browsers to backend services.

## Day 0 - Planning Phase

### Scope

  * **Describe the roadmap process, how scope is determined for mid to long term features, as well as how the roadmap
    maps back to current contributions and maintainer ladder?**

    The gRPC project's roadmap and technical evolution are managed through a formal, open, and community-driven
    governance model. The primary mechanism for this is the [gRPC Request for Comments (gRFC)
    process](https://github.com/grpc/proposal/blob/master/README.md). All significant technical changes, new features,
    and architectural proposals must be documented as gRFCs. These are developed in the open, submitted as pull requests
    to the [github.com/grpc/proposal](https://github.com/grpc/proposal/) repository and are subject to public review,
    discussion, and consensus-building among the community and project contributors.

    The authorship, review, and implementation of gRFCs is one of the core signals of attained leadership within the
    gRPC project and specifically within the [gRPC Project Contributor
    Ladder](https://github.com/grpc/grpc-community/blob/3711725571c0b94a9f6209e7304c36b4dc9f53da/contributor_ladder.md).
    Advancement from Organization Member to Core Contributor to Maintainer (the highest rung in the ladder) relies on
    many signals for the expertise and responsibility exhibited by an individual, with participation in the gRFC process
    being chief among these.

  * **Describe the target persona or user(s) for the project?**


    The gRPC project serves two primary personas:

    1.  **Application Developer / Software Engineer:** This is our primary user. This persona builds applications and
        services, often distributed systems. They value high performance, strong type-safety, and developer velocity. Their
        goal is to focus on business logic, not the complex plumbing of network protocols, data serialization, and error
        handling. They interact with gRPC by defining service contracts in Protocol Buffer definitions and using our code
        generators to create strongly-typed clients and servers. Developers often need to work in polyglot environments  and
        gRPC's ability to generate idiomatic code for over 10 languages (Go, Java, C\#, Python, C++, [and
        more](https://grpc.io/docs/languages/))  is a critical feature for them.

    2.  **Platform Engineer / Site Reliability Engineer (SRE):** This secondary persona is responsible for the
        infrastructure that *runs* gRPC services in production. They interact with gRPC's non-functional, operational
        features: its observability exports ([OpenTelemetry metrics](https://grpc.io/docs/guides/opentelemetry-metrics/)),
        its [standardized health-checking protocol](https://grpc.io/docs/guides/health-checking/), and [its rich
        configuration for resilience patterns like retries and deadlines](https://grpc.io/docs/guides/service-config/).

  * **Explain the primary use case for the project. What additional use cases are supported by the project?**

    The primary use case for gRPC is high-performance, low-latency microservice-to-microservice communication. Its
    efficiency, derived from using HTTP/2 as a transport and Protocol Buffers as the default mechanism for binary
    serialization, makes it the ideal choice for high-volume traffic inside a public or private cloud. This efficiency
    is paired with tooling enabling the design, deployment, and evolution of well-defined APIs.

    Additional supported use cases include:

      * **Real-Time Streaming:** gRPC has first-class, native support for client-side, server-side, and bi-directional
        streaming. This makes it an ideal fit for real-time applications like financial data feeds, live chat services,
        IoT command-and-control systems, or any reactive system.
      * **Polyglot System Integration:** Its contract-first design and broad language support make it a perfect
        mechanism for integrating systems implemented in different languages, as is often required within a large
        organization.
      * **"Last Mile" Connectivity:** gRPC is used to connect mobile devices, laptops, IoT endpoints, and network
        elements to backend services. In these scenarios, its low-bandwidth footprint (due to Protobuf) provides a
        significant advantage over text-based JSON.
      * **Cloud-Native Infrastructure:** gRPC is the protocol of choice for many foundational CNCF and infrastructure
        projects. A prime example is Kubernetes itself, which uses gRPC for its [Container Runtime Interface
        (CRI)](https://kubernetes.io/docs/concepts/containers/cri/) and [Container Storage Interface
        (CSI)](https://github.com/container-storage-interface/spec/blob/e981e2a057ca10f4a7f81289c97a4e829fd69152/spec.md).

  * **Explain which use cases have been identified as unsupported by the project.**

    The primary unsupported use case is direct, native communication from a web browser.

    This is not an oversight but a fundamental technical constraint. gRPC's protocol relies on direct control over
    HTTP/2 features (like full-duplex streaming and [trailing
    headers](https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Headers/Trailer)) that are not exposed to
    application-level JavaScript through standard browser APIs (e.g., `fetch` or `XMLHttpRequest`).

    To address this critical community need, the gRPC project supports and maintains
    [gRPC-Web](https://github.com/grpc/grpc-web). gRPC-Web is a related protocol and translation layer to native gRPC.
    It allows a browser-based client to speak the gRPC-Web protocol (which *can* be implemented in JavaScript) to a
    proxy (such as Envoy or a dedicated gRPC-Web proxy). This proxy then translates the calls into native gRPC for the
    backend services. This provides a documented and supported solution for browser-based applications.

  * **Describe the intended types of organizations who would benefit from adopting this project. (i.e. financial
    services, any software manufacturer, organizations providing platform engineering services)?**

    gRPC is a "universal RPC framework" intended for any organization building or operating modern distributed systems. Our adopters represent a considerable cross-section of the industry, including:

      * **Hyperscale & Tech:** Companies like
        [Netflix](https://netflixtechblog.com/practical-api-design-at-netflix-part-1-using-protobuf-fieldmask-35cfdc606518),
        [Reddit](https://old.reddit.com/r/RedditEng/comments/q5vmf2/reddits_move_to_grpc/),
        [LinkedIn](https://www.infoq.com/news/2023/12/linkedin-grpc-protobuf-rest-json/),
        [Square](https://medium.com/square-corner-blog/grpc-reaches-1-0-85728518393b#.x9w9iefe5), Google, and
        [Cloudflare](https://blog.cloudflare.com/road-to-grpc/) that manage millions of RPCs per second at minimal
        latency.
      * **Artificial Intelligence** Companies like [OpenAI](https://www.cncf.io/case-studies/openai/) and
        [Nvidia](https://github.com/ai-dynamo/dynamo/blob/09b26bf6b39df6fe9e2e1c635932af19fa8a6718/docs/frontends/kserve.md)
        find gRPC's efficiency, streaming, and polyglot support ideal for language models.
      * **Enterprise & SaaS:** Companies like [Salesforce](https://www.cncf.io/case-studies/salesforce/),
        [Cisco](https://www.cisco.com/c/en/us/td/docs/iosxr/ncs5000/programmability/75x/b-programmability-cg-ncs5000-75x/m-grpc-session.html),
        and [Juniper
        Networks](https://www.juniper.net/documentation/us/en/software/junos/grpc-network-services/topics/topic-map/grpc-services-configuring.html)
        that build large polyglot service ecosystems.
      * **FinTech:** Companies like
        [Coinbase](https://www.coinbase.com/blog/How-We-are-Improving-Product-Quality-at-Coinbase-with-AI-agents) that
        require high-performance, secure, and strongly-typed APIs for financial transactions.
      * **Platform Engineering & Observability:** Organizations building service meshes, databases, and monitoring
        platforms (e.g., [Datadog](https://www.datadoghq.com/blog/grpc-at-datadog/),
        [Authzed](https://authzed.com/docs/spicedb/getting-started/client-libraries)) often build on gRPC as their
        foundational communication layer.
      * **Startups & Growth-Stage:** Companies like
        [DoorDash](https://careersatdoordash.com/blog/building-a-grpc-client-standard-with-open-source/) and
        [GIPHY](https://web.archive.org/web/20240723105610/https://engineering.giphy.com/how-giphys-public-api-integrates-with-grpc-services/)
        that choose gRPC to build scalable, efficient, and maintainable backends from day one.

  * **Please describe any completed end user research and link to any reports.**

    As a foundational open-source framework, our "end user research" is conducted through deep, continuous, and public
    collaboration with our adopters, who are some of the largest-scale software producers in the world.

    This qualitative research is published in the form of public case studies, developer stories, and conference
    presentations, which are curated on [our project website](https://grpc.io/showcase/) and [by the
    CNCF](https://www.cncf.io/projects/grpc/](https://www.cncf.io/projects/grpc/). These reports detail adoption
    journeys, technical challenges, and real-world performance benefits.

    **Links to Adopter Reports (Case Studies):**

      * **gRPC Showcase (Canonical List):** [https://grpc.io/showcase/](https://grpc.io/showcase/)
      * **CNCF Project Page (Case Studies):** [https://www.cncf.io/projects/grpc/](https://www.cncf.io/projects/grpc/)
      * **Specific Adopter Stories:**
          * [Datadog: Adapting xDS for proxyless gRPC](https://www.youtube.com/watch?v=o5WiYFFMuD4&list=PLj6h78yzYM2P98usjCbJSkIS-rvIqEaVs&index=15)
          * [Reddit: Reddit's Move to gRPC](https://old.reddit.com/r/RedditEng/comments/q5vmf2/reddits_move_to_grpc/)
          * [Authzed (SpiceDB): Building a gRPC-first database](https://www.youtube.com/watch?v=1PiknT36218)
          * [Coinbase: Enhancing gRPC microservices](https://www.youtube.com/watch?v=GDJrw36wwWY)
          * [Salesforce](https://www.youtube.com/watch?v=MLS7TFHrn_c)

### Usability

  * **How should the target personas interact with your project?**

    As gRPC is a library and framework, not a standalone application, interaction is primarily through code and configuration.

      * **Developer Persona:** The primary interaction is a "contract-first" development loop:

        * **Define:** The developer defines service contracts, methods, and messages in a `.proto` (Protocol Buffer) file.
        * **Generate:** The developer uses the `protoc` compiler with a language-specific gRPC plugin (for example,
          `protoc-gen-go-grpc` in Go) to generate strongly-typed server interfaces and client stubs.
        * **Implement:** The developer implements the generated server interface with their business logic.
        * **Consume:** The developer imports the generated client stub, creates a "channel" (a logical connection) to
          the server, and invokes remote methods as if they were local functions.

      * **Platform Engineer Persona:** This persona interacts with gRPC at the *configuration* and *operations* level, typically via application config files or environment variables:

        * **Configure:** Platform Engineers configure gRPC channels with security credentials (e.g., TLS certificates),
          traffic management mechanisms (retries, deadlines), and load balancing policies (often via an
          [xDS](https://github.com/cncf/xds) control plane).
        * **Integrate:** Platform Engineers integrate gRPC's [standard health checks with orchestrators like
          Kubernetes](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/#define-a-grpc-liveness-probe).
        * **Observe:** Platform Engineers configure and consume the observability signals (metrics, traces, logs) that
          gRPC services export, primarily via OpenTelemetry.
        * **Govern:** Many Platform Engineers also use gRPC and Protocol Buffers to enforce API governance.
          Cross-cutting organizational policies surrounding what sorts of APIs are acceptable may be enforced by
          continous integration checks on the API definitions (for example, no backward-incompatible changes) and
          _runtime_ policies on acceptable usages of those APIs (for example, [data
          locality](https://youtu.be/9dBMPrwtGAQ?si=QzOZVuXlL_JbPewf&t=1686)) may be encoded by options on the API and
          enforced at runtime by middleware (for example, Envoy proxies).


  * **Describe the user experience (UX) and user interface (UI) of the project.**

      * **UI:** As a library, gRPC itself does not have a graphical user interface, however several popular UIs for
        invoking RPCs via gRPC have arisen, including [`grpcui`](https://github.com/fullstorydev/grpcui) and
        [Postman](https://blog.postman.com/postman-now-supports-grpc/).
      * **User Experience (UX):** The UX of gRPC is fundamentally its _Developer Experience_ (DX), which is [a core
        design principle of the project ("Coverage & Simplicity")](https://grpc.io/blog/principles/). The gRPC DX is
        centered on intentional API design flows, type-safety, and performance.
          * **Abstraction:** The core UX is the "local call" abstraction. Developers locally call a function and gRPC
            transparently handles network communication, serialization, and error propagation. This dramatically reduces
            the cognitive burden of building distributed systems.
          * **Type-Safety:** The contract-first approach  means that the compiler prevents entire classes of errors
            common in text-based protocols (for example, field name typos and data type mismatches).
      * The community has also built a rich ecosystem of CLI tools that enhance this experience, such as
        [`grpcurl`](https://github.com/fullstorydev/grpcurl) and
        [`grpcdebug`](https://github.com/grpc-ecosystem/grpcdebug).

  * **Describe how this project integrates with other projects in a production environment.**

    gRPC is the connective tissue of modern production environments.

      * **Orchestrators (Kubernetes):** gRPC is deeply integrated with Kubernetes.
          * It is the protocol used for the [Container Runtime Interface
            (CRI)](https://kubernetes.io/docs/concepts/containers/cri/) and [Container Storage Interface
            (CSI)](https://github.com/container-storage-interface/spec/blob/e981e2a057ca10f4a7f81289c97a4e829fd69152/spec.md).
          * It is the protocol used by [the etcdv3 API](https://etcd.io/docs/v3.4/learning/api/) and therefore involved
            in nearly every Kubernetes API server request.
          * Kubernetes (v1.24+) [provides native support for gRPC health
            probes](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/#define-a-grpc-liveness-probe),
            allowing kubelets to check service health without `exec` workarounds.
          * The Kubernetes Gateway API [provides official support for routing gRPC
            traffic](https://gateway-api.sigs.k8s.io/api-types/grpcroute/) with the `GRPCRoute` resource.
      * **Service Mesh (e.g., Istio, Linkerd, Cilium):**
          * Service meshes have first-class, L7-aware support for gRPC, providing transparent mTLS, advanced load
            balancing, and metrics.
          * gRPC's pluggable architecture also supports [proxyless service mesh
            deployments](https://www.youtube.com/watch?v=il2nekp_fIw). The gRPC client can integrate directly with an
            xDS control plane to receive routing, policy, and load balancing configuration, eliminating the need for a
            proxy of any kind.
      * **Observability (OpenTelemetry, Prometheus):**
          * gRPC provides [first-party integration with
            OpenTelemetry](https://grpc.io/docs/guides/opentelemetry-metrics/). This allows gRPC signals (such as
            metrics and logs) to be exported in the standard, vendor-neutral
            [OTLP](https://opentelemetry.io/docs/specs/otel/protocol/) format to any compatible backend (e.g.,
            Prometheus, Datadog, Jaeger).
      * **API Gateways (like Envoy, Kong, and `grpc-gateway`):**
          * To bridge the gap with other protocols (such as REST and gRPC-Web), gRPC integrates with API gateways for
            protocol transcoding. Ecosystem projects like
            [`grpc-gateway`](https://github.com/grpc-ecosystem/grpc-gateway) can automatically generate a REST/JSON
            reverse-proxy from a `.proto` definition.

### Design

  * **Explain the design principles and best practices the project is following.**

    The gRPC project's design is formally documented in the [Motivation and Design Principles blog
    post](https://grpc.io/blog/principles/). The core principles that guide our architecture are:

    1.  **Services not Objects, Messages not References:** Promote the microservices design philosophy of coarse-grained
        message exchange between systems while avoiding [the pitfalls of distributed
        objects](https://martinfowler.com/articles/distributed-objects-microservices.html) and [the fallacies of ignoring
        the network](https://en.wikipedia.org/wiki/Fallacies_of_distributed_computing).
    2.  **Coverage and Simplicity:** The stack should be available on every popular development platform and easy for
        someone to build for their platform of choice. It should be viable on CPU and memory-limited devices.
    3.  **Free and Open:** Make the fundamental features free for all to use. Release all artifacts as open-source
        efforts with licensing that should facilitate and not impede adoption.
    4.  **Interoperability & Reach:** The wire protocol must be capable of surviving traversal over common internet infrastructure.
    5.  **General Purpose & Performant:** The stack should be applicable to a broad class of use-cases while sacrificing
        little in performance when compared to a use-case specific stack.
    6.  **Layered:** Key facets of the stack must be able to evolve independently. A revision to the wire-format should
        not disrupt application layer bindings.
    7.  **Payload Agnostic:** Different services need to use different message types and encodings such as Protocol
        Buffers, JSON, XML, and Thrift; the protocol and implementations must allow for this. Similarly the need for payload
        compression varies by use-case and payload type: the protocol should allow for pluggable compression mechanisms.
    8.  **Streaming:** Storage systems rely on streaming and flow-control to express large data-sets. Other services,
        like voice-to-text or stock-tickers, rely on streaming to represent temporally related message sequences.
    9.  **Blocking & Non-Blocking:** Support both asynchronous and synchronous processing of the sequence of messages
        exchanged by a client and server. This is critical for scaling and handling streams on certain platforms.
    10. **Cancellation & Timeout:** Operations can be expensive and long-lived - cancellation allows servers to reclaim
        resources when clients are well-behaved. When a causal-chain of work is tracked, cancellation can cascade. A client
        may indicate a timeout for a call, which allows services to tune their behavior to the needs of the client.
    11. **Pluggable:** The core gRPC framework is minimal and extensible. Key functionalities like Authentication, Load
        Balancing, Tracing, and Health Checking are designed as pluggable interfaces. This allows adopters to integrate
        their own custom systems or use ecosystem standards like OpenTelemetry.
    12. **Lameducking:** Servers must be allowed to gracefully shut-down by rejecting new requests while continuing to
        process in-flight ones.
    13. **Flow Control:**  Computing power and network capacity are often unbalanced between client and server. Flow
        control allows for better buffer management as well as providing protection from DOS by an overly active peer.
    14. **Pluggable:** A wire protocol is only part of a functioning API infrastructure. Large distributed systems need
        security, health-checking, load-balancing and failover, monitoring, tracing, logging, and so on. Implementations
        should provide extension points to allow for plugging in these features and, where useful, default implementations.
    15. **Extensions as APIs:** Extensions that require collaboration among services should favor using APIs rather than
        protocol extensions where possible. Extensions of this type could include health-checking, service introspection,
        load monitoring, and load-balancing assignment.
    16. **Metadata Exchange:** Common cross-cutting concerns like authentication or tracing rely on the exchange of data
        that is not part of the declared interface of a service. Deployments rely on their ability to evolve these features
        at a different rate to the individual APIs exposed by services.
    17. **Standardized Status Codes:** Clients typically respond to errors returned by API calls in a limited number of
        ways. The status code namespace should be constrained to make these error handling decisions clearer. If richer
        domain-specific status is needed the metadata exchange mechanism can be used to provide that.

  * **Outline or link to the project’s architecture requirements? Describe how they differ for Proof of Concept,
    Development, Test and Production environments, as applicable.**

    gRPC's architecture requirements are captured by [its Motivation and Design
    principles](https://grpc.io/blog/principles/).

    While the gRPC library itself does not change between environments, a gRPC user's configuration of the runtime
    generally _does_ differ between environments:

      * **Concept / Development:** gRPC users begin with the simplest possible configuration for ease of use, generally
        on their workstation / laptop. This typically involves the use of `InsecureChannelCredentials` (that is,
        plaintext communication) and running both the client and server on a single machine, communicating over
        `localhost`. Service meshes, API Gateways, and middleware are generally not used at this stage.
      * **Test:** gRPC users begin rolling their service out to test / staging environments, which generally _do_
        include orchestrators like Kubernetes, service meshes, and API Gateways. While TLS encryption is generally used
        here, it does not involve production certificates. At this stage, all configuration is very similar to that used
        for production, but traffic is not exposed to end users of the service. Service mesh often aids users in the
        separation of their staging and production environments, even when running both in the same cluster. At this
        stage, developers often write end-to-end tests against their staging environment, taking advantage of gRPC's
        polyglot support to write the tests in languages besides those used in production (for example Python or Go when
        the production service is written in C++, Rust, or Java).
      * **Production:** In a production environment, all of the hardening mechanisms provided by gRPC and its ecosystem
        are enabled, providing the application with the strongest possible security and reliability postures:
          * **Security:** Using `SslCredentials` for TLS or mutual TLS (mTLS), or integrating with service mesh, which
            provides mTLS transparently.
          * **Resilience:** Configuring sensible default deadlines  and retry policies.
          * **Scalability:** Configuring client-side load balancing (e.g., `round_robin` or more advanced policies via xDS).
          * **Observability:** Enabling OpenTelemetry support  and registering the `grpc.health.v1` service.
          * **Security (Hardening):** Disabling server reflection, which is useful for debugging but can be a security
            risk in production as it exposes the entire API schema.

  * **Define any specific service dependencies the project relies on in the cluster.**

      As a library, gRPC itself has no runtime service dependencies. However, when an adopter deploys a gRPC application
      in a cluster, that application will typically depend on:

      * **DNS:** For basic service discovery. The gRPC client channel will resolve a hostname (for example,
        `my-service.my-namespace.svc.cluster.local`) via in-cluster DNS to find server IP addresses.
      * **(Optional) xDS Control Plane:** In proxyless service mesh deployments, the gRPC client can be configured to
        connect to an xDS-compatible control plane (such as Istio) to dynamically receive load balancing, routing, and
        security policies.

  * **Describe how the project implements Identity and Access Management.**

    As a library, the gRPC runtime provide developers the mechanisms to implement their desired IAM policies and the
    developer leverages them to _implement_ those policies. These mechanisms cover all aspects of AAA (Authentication,
    Access Management, and Accounting):

      * **Identity (Authentication / AuthN):**

        1. **Transport-Layer (TLS):** [The default secure mechanism is
          SSL/TLS](https://grpc.io/docs/guides/auth/#supported-auth-mechanisms), which authenticates the *server* to the
          client (and optionally the client to the server via mTLS) using X.509 certificates.
        2. **Per-Call (Tokens):** gRPC provides a `CallCredentials` abstraction for attaching per-RPC identity tokens
           (such as OAuth2 and JWT). These are typically sent as `Authorization` metadata headers.
        3. **Credentials Plugins:** If a user finds that the two mechanisms above are not sufficient on their own, gRPC
           provides [a credentials plugin
           API](https://grpc.io/docs/guides/auth/#extending-grpc-to-support-other-authentication-mechanisms), which allows
           developers to provide authentication however they'd like.

      * **Access Management (Authorization / AuthZ):**

        1.  **Interceptor-Based Authz:** Several gRPC languages provide batteries-included Authz-enforcing interceptors
           which derive their configuration from either a file on the filesystem or from a programmatically provided
           string, as defined in [gRFC
           A43](https://github.com/grpc/proposal/blob/fcabdfdbd50b3c088f5a5c2bf925755781ec076e/A43-grpc-authorization-api.md).
        2.  **Service Mesh-Enabled RBAC-based Authz:** gRPC supports [the same xDS RBAC filters as
          Envoy](https://github.com/envoyproxy/envoy/blob/23b03a82626b22fca5bb2c1b3147dcd4569e18e5/api/envoy/extensions/filters/http/rbac/v3/rbac.proto),
          as defined in [gRFC
          A41](https://github.com/grpc/proposal/blob/fcabdfdbd50b3c088f5a5c2bf925755781ec076e/A41-xds-rbac.md). Proxyless
          service mesh users may configure their xDS control plane with Authz policies. The control plane will then
          deliver this RBAC config to gRPC clients and servers, which will begin transparently enforcing those policies.
        3.  **User-Defined Interceptors:** If either of the above authz mechanisms does not satisfy a particular gRPC
          user, they are free to implement their own authz functionality using the pluggable server-side [interceptor
          mechanism](https://grpc.io/docs/guides/interceptors/). For example, [this community-built Kerberos
          implementation](https://github.com/jcmturner/grpckrb).

      * **Accounting (Audit):**

        1.  **Audit Logging:** gRPC's comprehensive audit logging mechanism is outlined in [gRFC
           A59](https://github.com/grpc/proposal/blob/fcabdfdbd50b3c088f5a5c2bf925755781ec076e/A59-audit-logging.md). It
           may be configured via the same mechanisms as the A41 RBAC mechanism mentioned above or via an xDS control plane
           (that is, a service mesh).
        2.  **Binary Logging:** Binary logging (detailed in [gRFC
          A16](https://github.com/grpc/proposal/blob/fcabdfdbd50b3c088f5a5c2bf925755781ec076e/A16-binary-logging.md)) is
          an earlier mechanism used to log _all_ aspects of an RPC including message payloads, making it useful for
          troubleshooting and replaying messages, but also for auditing.
        3.  **User-Defined Interceptors:** As with the other two aspects of AAA, if either of the above audit logging
          mechanisms does not satisfy the user, they are free to implement their own via interceptors, like [this
          community
          project](https://github.com/grpc-ecosystem/go-grpc-middleware/tree/390bcef25adebe4b0c7dbb365230c0a856737afe/interceptors/logging).

  * **Describe how the project has addressed sovereignty.**

      As a general-purpose communication framework, gRPC does not directly address data sovereignty. However, some of
      gRPC's features can be used as tools by gRPC users to achieve data sovereignty:

      * **Routing:** gRPC provides native _routing_ functionality via xDS. When paired with a sovereignty-aware xDS
        control plane, clients will only be aware of legally permissible backends and only connect to those backends.
      * **Sovereignty Tagging:** Services, methods, messages, and fields can be marked with sovereignty requirements,
        which will be enforced by proxies ([example](https://youtu.be/9dBMPrwtGAQ?si=2c2NUmLF98VGWGWT&t=1681)). In this
        scenario, the user would be responsible for supplying the code generation and proxy functionality.

  * **Describe any compliance requirements addressed by the project.**

      While gRPC itself does not address specific compliance regimes (e.g., PCI, HIPAA, FedRAMP), it does provide some
      _technical prerequisites_ that enable users to build compliant systems.

      * **Encryption in Transit:** Our built-in, easy-to-use TLS implementation allows adopters to meet the
        data-in-transit encryption requirements common to all major compliance standards.

      * **Audit Trail:** [Audit
        logging](https://github.com/grpc/proposal/blob/fcabdfdbd50b3c088f5a5c2bf925755781ec076e/A59-audit-logging.md)
        provides a formal, pluggable mechanism for creating the immutable audit trails required by standards like PCI,
        SOC2, and HIPAA.

  * **Describe the project’s High Availability requirements.**

      gRPC provides core features that enable users to build Highly Available applications. A naiive gRPC setup (one
      client, one server) is not Highly Available. A cloud-native, HA gRPC architecture is achieved by combining the
      following gRPC features / configurations with an orchestrator:

      *  Server Side
          *  Redundancy: The user runs multiple instances of their gRPC servers (for example, in a [Kubernetes
             Deployment](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/)).
          *  Health Checking: Each server instance exposes [the standard gRPC Health check
             service](https://grpc.io/docs/guides/health-checking/).
          *  Graceful Shutdown: Servers are configured to finish up serving outstanding requests without accepting new
            ones when preparing to shut down via the [graceful shutdown
            mechanism](https://grpc.io/docs/guides/server-graceful-stop/), ensuring that routine server instance
            restarts do not result in the failure to service any requests.
      *  Client-Side
          * Service Discovery (for example, xDS, DNS, or a [Kubernetes-native name resolver
            plugin](https://github.com/sercand/kuberesolver)) to find all server instances.
          * Health Checking in the control plane to actively probe backends and remove unhealthy ones from the pool
            actively receiving traffic from clients.
          * [Client-Side Load Balancing](https://grpc.io/blog/grpc-load-balancing/) (for example,
            [`round_robin`](https://youtu.be/G6PRjmXuBG8?si=kVtZrsRbSJ6Wf2yj&t=444)) to distribute requests across the
            *healthy* instances.
          * [Retries](https://grpc.io/docs/guides/retry/) to automatically retry requests that fail with `UNAVAILABLE`
            (e.g., during a brief network partition or pod restart). This is done as an exponential backoff with jitter
            to avoid the [thundering herd problem](https://en.wikipedia.org/wiki/Thundering_herd_problem).

  * **Describe the project’s resource requirements, including CPU, Network and Memory.**

    As a library, gRPC's resource requirements are a function of usage (e.g., RPCs per second, message size). Specific
    numbers will vary between the many languages, operating systems, platforms supported by gRPC. Up-to-date benchmarks
    with concrete numbers are available [here](https://grafana-dot-grpc-testing.appspot.com/).

      * **CPU:** gRPC is highly CPU-efficient. Its primary CPU cost is in Protobuf serialization/deserialization and
        HTTP/2 frame processing. This [is consistently benchmarked](https://grpc.io/docs/guides/benchmarking/) as being
        significantly lower than the CPU cost of text-based JSON parsing and serialization.
      * **Network:** gRPC is extremely network-efficient. The binary serialization format and HTTP/2 header compression
        result in a [2-3x speedup compared to HTTP+JSON](https://grpc.io/blog/mobile-benchmarks/).
      * **Memory:** While the gRPC library is designed and verified to have a low memory footprint, gRPC _users_ must
        also take some care to ensure that received messages do not consume an unacceptable amount of memory. gRPC
        provides several mechanisms to achieve this by limiting concurrency. For streaming RPCs, [gRPC provides support
        for flow control](https://grpc.io/docs/guides/flow-control/), which enables the server to apply pushback on
        individual clients if they send messages more quickly than the server can process them. [Other intelligent
        mechanisms for concurrency limitation](https://github.com/Netflix/concurrency-limits) have been developed by the
        community leveraging gRPC's pluggable interceptor mechanism. 

  * **Describe the project’s storage requirements, including its use of ephemeral and/or persistent storage.**

      * **N/A.** As a network communication framework, gRPC does not use ephemeral or persistent storage.

  * **Please outline the project’s API Design:**

      * **Describe the project’s API topology and conventions**

        gRPC provides APIs at three levels:

        1. Networked APIs authored by the user, _transported over_ gRPC, leveraging [gRPC HTTP/2
          framing](https://github.com/grpc/grpc/blob/688a2868b8dc6cabba89670d71e2df4fa3d1f9f7/doc/PROTOCOL-HTTP2.md) and
          typically using Protocol Buffers as the payload.
        2. Programmatic APIs provided in the user's language of choice enabling the user to instantiate and use servers
          and stubs at the level of a byte stream ([detailed per-language under "API
          Reference"](https://grpc.io/docs/languages/))
        3. Generated code APIs giving users per-language access to their IDL-defined (e.g. protobuf-defined) network
          APIs from their langauge of choice on both the client-side and server-side. These are implemented on top of the
          byte-oriented APIs mentioned above.

        gRPC users _start_ by authoring their networked API schema in their IDL of choice (generally Protocol Buffers or
        Flatbuffers). This defines the network protocol that clients and servers will use to communicate with one
        another. They then use a language-specific code generator provided by the gRPC project ([C++
        example](https://grpc.io/docs/languages/cpp/basics/#generating-client-and-server-code), [Go
        example](https://grpc.io/docs/languages/cpp/basics/#generating-client-and-server-code)) to generate functions
        corresponding to each RPC in their RPC schema and server interfaces on which they can build their application
        server logic. These generated APIs marshal and unmarshal between the IDL's runtime format (typically protobuf
        messages) and bytes. These bytes are then supplied to the IDL-agnostic byte-oriented API.

        [gRPC supports several different varieties of
        _streaming_](https://grpc.io/docs/what-is-grpc/core-concepts/#service-definition). The most common RPC methods
        are _unary_ -- a single client-originated request followed by a single server-originated response. _Streaming_
        RPC methods allow clients and servers to exchange multiple messages in support of that method. There are three
        kinds of streaming RPCs:

        * In server-streaming RPC methods, the client sends a single request message. The server then responds with 0 or
          more response messages back to the client. gRPC preserves the server's ordering of these messages.
        * In client-streaming RPC methods, the client sends 0 or more request messages, each separated by any amount of
          time. After the client has finished sending requests, the server then responds with a single response message.
        * In bidirectional (or bidi) RPC methods, the client initiates the connection, but the client and server may
          exchange messages in _any_ order. The stream of client-originated messages and the stream of server-originated
          messages are treated independently and the order of each stream is preserved.

        These 4 kinds of RPC (unary, server-streaming, client-streaming, and bidi) are referred to as the 4 _arities_.
        The arity of an RPC method is explicitly marked in the IDL with unary being the default.


      * **Describe the project defaults**

          * IDL: We default to Protocol Buffers (most documentation is geared toward them). Using other IDLs such as
            FlatBuffers requires additional configuration on the part of the user.
          * Load Balancing: The default client-side load-balancing policy is `pick_first`. This policy resolves all
            available IP addresses for a target, connects to the first one that succeeds, and uses that single
            connection until it fails.
          * Max receive message size: By default, clients and servers are configured to fail RPCs when receiving a
            message larger than the default limit of 4MB. This is part of the protection against memory overruns from
            untrusted or misbehaving peers. In systems where care is being taken elsewhere in the stack to protect
            against memory overruns, this configuration may be safely raised.
          * Max metadata size: For similar reasons of protection against memory overruns, max metadata size for a single
            RPC is limited to 8KB. This may also be overridden by the user.
          * Compression: gRPC does not enable compression by default. Users may choose to override this and enable an
            algorithm such as `Deflate` or `Gzip`.


      * **Outline any additional configurations from default to make reasonable use of the project**

        Users of gRPC in production should configure the following:

        * Client Side
            * [RPC timeout](https://grpc.io/docs/guides/deadlines/)
            * [Retry policy](https://grpc.io/docs/guides/retry/)
            * [An appropriate LB
              policy](https://github.com/grpc/grpc/blob/560e95a3f4656998063a4b0ca1dd88ed54a34dc9/doc/load-balancing.md)
              for their service topology (e.g. [Round
              Robin](https://github.com/grpc/grpc/blob/560e95a3f4656998063a4b0ca1dd88ed54a34dc9/doc/load-balancing.md))
        * Server Side
            * A [reflection service](https://grpc.io/docs/guides/reflection/)
            * A [health check service](https://grpc.io/docs/guides/health-checking/)
        * Both
            * TLS -- gRPC does _not_ provide any "insecure-by-default" APIs, but users should explicitly configure TLS
              or mTLS in production applications.
            * A [channelz service](https://grpc.io/blog/a-short-introduction-to-channelz/)

        While not a gRPC configuration per se, users deploying their gRPC applications on Kubernetes should be aware
        that Kubernetes ClusterIP Services load balance at the level of _TCP connections_ while gRPC utilizes long lived
        TCP connections and instead does _client side_ load balancing at the level of HTTP/2 streams. To achieve
        satisfactorily uniform load distribution, these users should ensure that the _gRPC library_ makes load balancing
        decisions rather than `kube-proxy`. To achieve this, they should do one of the following:

        * Enable a Service Mesh (such as proxyless gRPC or Istio)
        * Configure their target gRPC Service as a [headless
          Service](https://kubernetes.io/docs/concepts/services-networking/service/#headless-services) and enable the
          Round Robin LB policy within the gRPC client library.
        * Use a Kubernetes-native gRPC name resolver plugin such as [Kuberesolver](https://github.com/sercand/kuberesolver).

        A full case study in productionizing gRPC by Netflix is available [here](https://www.youtube.com/watch?v=NTf_2bzD7xM).

      * **Describe any new or changed API types and calls - including to cloud providers - that will result from this
        project being enabled and used**

          As an application library, gRPC itself does not install any new Kubernetes Resource definitions into user
          clusters. However, the [GRPCRoute resource](https://gateway-api.sigs.k8s.io/api-types/grpcroute/) is available
          within the upstream Kubernetes [Gateway API](https://gateway-api.sigs.k8s.io/) to aid with routing gRPC
          traffic within a Kubernetes cluster.

          gRPC will not make any API calls unless explicitly instructed to. gRPC will reach out to DNS servers and xDS
          control planes if _explicitly_ instructed to.

      * **Describe compatibility of any new or changed APIs with API servers, including the Kubernetes API server** 

          * **N/A.** gRPC does not interact with or change the Kubernetes API server.

      * **Describe versioning of any new or changed APIs, including how breaking changes are handled**
        This has two distinct parts: gRPC's *framework* versioning, and how we enable *adopters* to version their services.

        1.  **gRPC Framework Versioning:**

              * We follow a `v1.x.y` policy where `x` is the minor version and `y` is the patch version. Breaking
                changes to non-experimental APIs are only allowed in _major_ version bumps.
              * Minor releases occur approximately every 6 weeks.
              * While the vast majority of per-language gRPC implementations are still on `v1.x`, major version bumps
                (e.g., `v1.x` to `v2.x`) *must not* break _wire compatibility_. That is, a `v2.x` client *must* be able
                to interoperate with a `v1.x` server, and vice versa. This guarantee is fundamental to enabling safe
                service evolution for our adopters.

        2.  **Adopter's Service API Versioning:**

              While backward compatibility of the application API is ultimately the responsibility of the gRPC user,
              Protocol Buffers (the default IDL for gRPC) give users the tools to do this easily. 

              Protocol Buffers are intentionally designed to _minimize the need for_ breaking changes. If the user
              follows Protocol Buffer best practices ([1](https://protobuf.dev/best-practices/dos-donts/),
              [2](https://google.aip.dev/)), the odds that they will need to make a breaking change _at all_ in the
              future will be greatly reduced.

              When a new version of a protobuf-schema'd API _is_ required, there is a well-established pattern to handle
              this. By convention, protobuf packages include a version segment (e.g. `grpc.reflection.v1`), similar to
              how API versions are often a segment in HTTP APIs' URLs -- and in fact this package _does_ get included in
              the underlying HTTP/2 URL. When a new major version of the API is required, a new proto-based API should
              be introduced with the new version number in its package.

              Even when breaking changes are made, it is common for many fields to remain the same. To identify
              equivalent fields across major version bumps, projects including Envoy have developed a set of annotations
              to express this cross-version equivalence
              ([annotations](https://github.com/cncf/udpa/blob/c52dc94e7fbe6449d8465faaeda22c76ca62d4ff/udpa/annotations/migrate.proto),
              [Envoy API best
              practices](https://github.com/envoyproxy/data-plane-api/blob/2fd335d9c9ad57da08474e37eba091b40a67f561/STYLE.md)).

  * **Describe the project’s release processes, including major, minor and patch releases.**

      While the different subprojects within gRPC (mostly different language implementations) are free to release on
      their own schedule according to their own release policy, the C++, Java, Go, and Python _do_ attempt to
      synchronize their releases as much as possible, according to the following scheme:

      * **Minor Releases:** These are our "checkpoint" releases. They happen on a fixed, public schedule, approximately
        every 6 weeks. A new release branch is cut from the `main` / `master` branch, which is kept stable at all times.
      * **Patch Releases:** These are cut from an existing release branch to fix critical bugs or security
        vulnerabilities.
      * **Major Releases:** These are extremely rare and are only permitted for significant, necessary API-breaking
        changes (e.g., fixing a security flaw, adapting to a language-ecosystem break). They require a formal gRFC and,
        as stated, *must not* break wire protocol compatibility. To this point, C# is the only implementation that has
        performed a major version bump beyond `v1.x`.

### Installation

  * **Describe how the project is installed and initialized, e.g. a minimal install with a few lines of code or does it
    require more complex integration and configuration?**

    As a library, gRPC is installed as a dependency using standard, language-native package managers.

      * Installation Examples:

          * Go: `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`
          * Python: `pip install grpcio grpcio-tools`
          * Java (Gradle): Adding `grpc-netty-shaded`, `grpc-protobuf`, and `grpc-stub` to `dependencies` in `build.gradle`.
          * C++: Requires building from source using `cmake`.

      * Basic Initialization: A minimal server or client can be initialized in just a few lines of code. This looks
        roughly the same across all languages:
        1.  Import the generated stubs and the core gRPC library.
        2.  On the server side, create a server instance, register your service implementation, and listen on a port.
        3.  On the client side, create a channel instance pointing to the server address, create a new client stub using
            that channel, and send RPCs on that stub.

  * **How does an adopter test and validate the installation?**

    If the user deploys both a gRPC client _and_ a gRPC server, validation of the gRPC deployment is the same as
    validation of the client and server applications, i.e. if the client and server work then gRPC is working.

    If only a _server_ is deployed, then tools like [`grpcurl`](https://github.com/fullstorydev/grpcurl) can be used to
    manually invoke RPCs and verify the proper functioning of the gRPC deployment.

### Security

  * **Please provide a link to the project’s cloud native [security self assessment](https://tag-security.cncf.io/community/assessments/).**

    The gRPC project partnered with [Cure53](https://cure53.de/) to conduct an external security audit as part of the
    CNCF graduation process. The results are published
    [here](https://github.com/grpc/grpc/blob/8d22d6239230b9f3c446f0bac719e46460cbdb30/doc/grpc_security_audit.pdf).

  * **Please review the [Cloud Native Security
    Tenets](https://github.com/cncf/tag-security/blob/main/community/resources/security-whitepaper/secure-defaults-cloud-native-8.md)
    from TAG Security.**

      * **How are you satisfying the tenets of cloud native security projects?**

        * **Make security a design requirement**

            Every major feature in the gRPC project must go through [the gRFC
            process](https://github.com/grpc/proposal), during which domain experts conduct a thorough review of the
            requirements and the design itself. The contributor ladder [explicitly calls out "Security" as an area of
            expertise](https://github.com/grpc/grpc-community/blob/3711725571c0b94a9f6209e7304c36b4dc9f53da/contributors/core_contributors.md)
            and these contributors are expected to weigh in on all relevant gRFCs.

        * **Applying secure configuration has the best user experience**

            gRPC proxyless service mesh / xDS support provides turn-key support for mTLS, traditionally a major burden
            to configure, as outlined in gRFCs
            [A87](https://github.com/grpc/proposal/blob/2405799e80581b3a3add27663116db6eada59154/A87-mtls-spiffe-support.md)
            and
            [A65](https://github.com/grpc/proposal/blob/2405799e80581b3a3add27663116db6eada59154/A65-xds-mtls-creds-in-bootstrap.md).

           Traditionally, configuring mTLS required developers to manually manage file paths for certificates and keys
           in code. With gRPC's xDS support, a developer simply uses xDS credentials options. The gRPC client then
           dynamically fetches security configuration (certificate bundles, validation contexts, and mTLS settings) from
           the control plane. This makes "secure by default" a zero-code-change operation for application developers,
           abstracting the complexity of certificate rotation and trust domain management entirely out of the
           application binary. 

        * **Selecting insecure configuration is a conscious decision**

            gRPC has embraced the "secure by default" philosophy from the outset. All cleartext APIs explicitly include
            either the word "insecure" or "plaintext," signalling their intent to take on the risk associated with these
            APIs.

        * **Transition from insecure to secure state is possible**

           The gRPC library gives users all the tools they need to transition an insecure/plaintext gRPC service to TLS.
           The process is as follows:

            * Adding a second port on the server which will serve requests over TLS.
            * Publishing the new Service via (e.g.) a Kubernetes Service resource
            * Gradually migrating clients from plaintext to TLS
            * Monitoring the server's traffic (for example, via gRPC's OTel integrations) to verify that plaintext
              traffic has fallen to zero
            * Removing the original plaintext port


        * **Secure defaults are inherited**

            gRPC is designed to inherit the security guarantees of the underlying transport protocol: TLS. We do not
            reinvent the wheel on encryption algorithms or handshake protocols. Instead, we expose the mature,
            battle-tested implementations of the underlying platforms (BoringSSL for wrapped languages like C++ and
            Python, `crypto/tls` for Go, Netty/JDK for Java).

            When a user enables TLS in gRPC, they automatically inherit the secure defaults of those libraries,
            including modern cipher suite preferences and protocol versions, without having to configure them manually.
            This shared responsibility model ensures that as the underlying SSL/TLS libraries patch vulnerabilities or
            deprecate weak ciphers, gRPC applications inherit these hardened postures instantly upon rebuild.

        * **Exception lists have first class support**

            For plaintext communication, gRPC users are free to bypass the library's warnings and use the various APIs
            with "insecure" or "plaintext" in the name and incur the attendant risks.

        * **Secure defaults protect against pervasive vulnerability exploits**

            gRPC makes extensive use of [fuzzers](https://github.com/grpc/grpc/tree/master/test/core/end2end/fuzzers) to
            identify vulnerabilities in the codebase as they are introduced. This has prevented many vulnerabilities
            from being introduced.

            Design review also helps protect users from vulnerabilities. [The compression API, for example, was
            expanded](https://grpc.io/docs/guides/compression/#specific-disabling-of-compression) to allow effective
            defense against [BEAST](https://en.wikipedia.org/wiki/Transport_Layer_Security#BEAST_attack) and
            [CRIME](https://en.wikipedia.org/wiki/CRIME) attacks.

        * **Security limitations of a system are explainable**

            We are transparent that gRPC is a transport and RPC framework, not a full-featured Identity Provider or
            Policy Engine.

            We document clearly that while we provide the hooks for authentication (SSL/TLS, Token auth) and
            authorization (Interceptors), gRPC does not ship with an embedded "User Database". Instead, we guide users
            toward purpose-built tools like SPIFFE/SPIRE or Open Policy Agent.

      * **How do you recommend users alter security defaults in order to "loosen" the security of the project? Please
        link to any documentation the project has written concerning these use cases.**

        While we do not recommend that security configurations be loosened in any production environment, we do
        recognize several cases where users frequently ask for restrictions to be limited. Primarily, in test
        environments and localhost communication.

        In these cases, some users find it burdensome to set up the infrastructure for TLS certificates. But certain
        functionality _is not allowed_ on insecure channels, such as call credentials (e.g. tokens). Rather than falling
        all the way back to the weakest security posture and allowing tokens to be sent over insecure channels, we
        instead introduced the concept of _local_ channel credentials ([Python
        example](https://grpc.github.io/grpc/python/grpc.html#grpc.local_channel_credentials)). These are channel
        credentials which require no cryptographic configuration, but are only allowed in contexts which have _some_
        known security guarantees due to the operating system: Unix Domains Socket communication and localhost TCP
        communication. This middle ground allows both ease of use _and_ satisfactory security guarantees.

  * **Security Hygiene**

      * **Please describe the frameworks, practices and procedures the project uses to maintain the basic health and
        security of the project.**

        * Vulnerability Management: The gRPC CVE reporting process is defined in [gRFC
          P4](https://github.com/grpc/proposal/blob/2405799e80581b3a3add27663116db6eada59154/P4-grpc-cve-process.md).
          CVEs are privately disclosed to the email address grpc-security@googlegroups.com. The reporter is kept
          apprised of the progress of the CVE from root cause analysis, fix implementation, fix rollout, and
          retrospective.
        * Public Disclosure: Also as outlined in gRFC P4, all CVEs will be announced on [the grpc-io mailing
          list](https://groups.google.com/g/grpc-io) once resolved to inform gRPC users.
        * Fuzzing & Sanitizers: We employ continuous fuzzing tests and use sanitizers (asan, tsan, etc.) in our CI to
          detect memory-related bugs, data races, and integer overflows.
        * Code Review: All code changes are submitted via Pull Request and must be reviewed and approved by project
          Maintainers. Pull Requests touching on security must be approved by our security domain experts.

      * **Describe how the project has evaluated which features will be a security risk to users if they are not
        maintained by the project?**

        The decision to implement specific functionality within the gRPC library in-house or to incur a dependency on a
        third-party library is a function of:

        * The existence of third-party, high-fidelity implementations in the target language
        * The level of expertise in the subject area among gRPC project contributors
        * The criticality of faults within the implementation

        For example, _no_ gRPC language implements its own cryptography primitives such as TLS. Instead, external
        libraries such as BoringSSL are used instead.  On the other hand, HTTP/2 lies very much within the domain of
        expertise of the gRPC project contributors. So in the case of C++/Core where no high-fidelity HTTP/2
        implementation was available at the time of release, gRPC implemented its own HTTP/2 stack. In Java and Go
        however, high-fidelity third-party implementations _did_ exist (Netty and `x/net/http2` respectively), so Java
        and Go _did not_ implement their own HTTP/2 stack.

  * **Cloud Native Threat Modeling**

      * **Explain the least minimal privileges required by the project and reasons for additional privileges.**

          gRPC is a userspace networking library. As a result, the vast majority of gRPC applications will require _no_
          special privileges to operate. There are a few exceptions to this rule of thumb, all of which follow directly
          from the requirements of the Linux kernel.

          * Low ports: While most gRPC server applications bind to a high port (50051 is customary), binding to port
            1023 or lower will require
            [`CAP_NET_BIND_SERVICE`](https://man7.org/linux/man-pages/man7/capabilities.7.html#:~:text=CAP_NET_BIND_SERVICE).
          * Unix Domain Sockets: When communicating over a Unix Domain Socket, the user under which the gRPC application
            is running must have write permissions to the chosen UDS. This does _not_ mean that root is required and we
            do _not_ recommend using the root user to solve this issue.
          * Private keys: In order to run a server over TLS, the user under which the gRPC server is running must have
            read permissions to the private key file.

      * **Describe how the project is handling certificate rotation and mitigates any issues with certificates.**

        gRPC provides the necessary hooks to implement TLS certificate rotation. The TlsCredentials APIs support dynamic
        reloading of certificate and key material from disk or memory. This allows an adopter's application to pick up
        new certificates (e.g., when a Kubernetes secret is updated) without a restart, enabling seamless, zero-downtime
        certificate rotation. This is outlined in detail in [gRFC
        A69](https://github.com/grpc/proposal/blob/2405799e80581b3a3add27663116db6eada59154/A69-crl-enhancements.md).

      * **Describe how the project is following and implementing [secure software supply chain best
        practices](https://project.linuxfoundation.org/hubfs/CNCF_SSCP_v1.pdf)**

        The gRPC project grants elevated permissions to only the most senior contributors (generally
        [Maintainers](https://github.com/grpc/grpc-community/blob/3711725571c0b94a9f6209e7304c36b4dc9f53da/contributors/maintainers.md)).
        While each repo varies in its precise CI setup based on their individual needs, the following is true across all
        repos:

        * Every PR must be reviewed by at least 1 contributor before eligibility for merge (the four eyes principle)
        * A comprehensive suite of pre-submit tests must pass on the change before eligibility for merge, often
          including fuzz testing
        * The `main` / `master` branches and all active release branches are branch-protected
        * All third-party runtime dependencies are verified at least by hash (though gRPC's dependency list is
          intentionally kept quite small)


## Day 1 - Installation and Deployment Phase

### Project Installation and Configuration

  * **Describe what project installation and configuration look like.**
      * **Installation (Day 0):**

        The gRPC library is installed into an application's codebase using language-specific package managers (e.g.,
        `pip`, `go get`, `mvn`).

      * **Configuration (Day 1 - Deployment):**

        The gRPC library is primarily configured _programmatically_, i.e. via functions/methods called on the library.
        For example, the listening port of a server is determined based on a string programmatically passed to a Server
        constructor/method. Some applications using gRPC will choose to make that configuration available
        non-programmatically (e.g. via environment variable or config file), but that is left to the user. There are a
        few notable exceptions where the gRPC library _does_ read configuration from files:

        * Private Keys: When running a server with TLS encryption, the private key will be read from file.
        * Auth tokens: Some cloud-provider-specific kinds of auth token, such as [Google Default
          Credentials](https://github.com/grpc/grpc-community/blob/3711725571c0b94a9f6209e7304c36b4dc9f53da/contributors/maintainers.md)
          will read the user's identity from the file system.
        * xDS Service Mesh: When service mesh functionality is configured, an xDS bootstrap file (originally introduced
          in [gRFC
          A27](https://github.com/grpc/proposal/blob/2405799e80581b3a3add27663116db6eada59154/A27-xds-global-load-balancing.md))
          is read from disk to containing the address of the control plane, the credentials to use, and other bootstrap
          xDS information.

### Project Enablement and Rollback

  * **How can this project be enabled or disabled in a live cluster? Please describe any downtime required of the
    control plane or nodes.**

      **N/A.** As a library, gRPC itself cannot be directly enabled or disabled. An application _using_ gRPC can of
      course be enabled/disabled. The procedure for this depends on the particular application.

  * **Describe how enabling the project changes any default behavior of the cluster or running workloads.**

      **N/A.** As a library, gRPC does not change any cluster behavior. Without user extensions, gRPC will not make
      requests to the Kubernetes API server at all.

  * **Describe how the project tests enablement and disablement.**

      **N/A.** As a library, gRPC itself cannot be directly enabled or disabled. An application _using_ gRPC can of
      course be enabled/disabled. The procedure for this depends on the particular application.

  * **How does the project clean up any resources created, including CRDs?**

      **N/A.** Since gRPC does not interact with Kubernetes resources, there is no clean-up required. In terms of
      _operating system_-level resources, allocated memory, sockets, FDs, etc. will be cleaned up when the application
      process exits (i.e. when the pod is terminated) or when the application explicitly destructs the owning objects.

### Rollout, Upgrade and Rollback Planning

  * **How does the project intend to provide and maintain compatibility with infrastructure and orchestration management
    tools like Kubernetes and with what frequency?**

    The gRPC library itself does not integrate with the control plane of Kubernetes or other container orchestrators.
    Instead, the gRPC library builds on top of solid, stable _operating system_ APIs (e.g. Linux syscalls). Access to
    these APIs forms the primary basis for interaction with a Kubernetes cluster. As a result of this "lowest common
    denominator" approach to integration with container orchestrators, maintenance of integration with Kubernetes is
    low-cost and equivalent to maintaining the core functionality of the library (which we accomplish through the many
    quality control techniques laid out in this document).

  * **Describe how the project handles rollback procedures.**

    gRPC _supports_ applications in their rollback procedures. gRPC maintains an extensive cross-version, cross-language
    compatibility matrix verified by [interop
    tests](https://grpc.github.io/grpc/core/md_doc_interop-test-descriptions.html). These tests ensure that applications
    are able to roll back their applications to previous versions of gRPC and they will continue to be compatible with
    their gRPC peers.

  * **How can a rollout or rollback fail? Describe any impact to already running workloads.**

    In the case that the user adopts the default IDL (protocol buffers), a rollout/rollback of a gRPC-based service will
    fail if the adopter violates the compatibility rules, specifically by making a breaking change in their `.proto`
    files. Any clients to the service being rolled out would begin observing failed RPCs with status `INVALID_ARGUMENT`,
    `UNIMPLEMENTED`, or `INTERNAL`.

  * **Describe any specific metrics that should inform a rollback.**

    Adopters should have dashboards and alerts monitoring their gRPC status codes, which we provide via OpenTelemetry
    metrics. A sudden spike in the following server-side metrics (SLIs) immediately following a deployment is a clear
    signal to initiate a rollback:

      * `sum(increase(grpc_server_call_duration_count{grpc_status="UNIMPLEMENTED"}[1h]))`
      * `sum(increase(grpc_server_call_duration_count{grpc_status="INVALID_ARGUMENT"}[1h]))`
      * `sum(increase(grpc_server_call_duration_count{grpc_status="UNKNOWN"}[1h]))`

  * **Explain how upgrades and rollbacks were tested and how the upgrade-\>downgrade-\>upgrade path was tested.**

    As mentioned above, the gRPC project tests cross-version interoperability continuously via our interop test suite.

  * **Explain how the project informs users of deprecations and removals of features and APIs.**

      Deprecations and removals of stable APIs require a major version bump of the library. To this point, only C# has
      undergone a major version bump due to a standard library type in its public API which was deprecated and removed.
      Any deprecation and removal of a stable API is announced well ahead of time on [the grpc-io mailing
      list](https://groups.google.com/g/grpc-io) and in release notes on Github.

  * **Explain how the project permits utilization of alpha and beta capabilities as part of a rollout.**

    We designate "alpha/beta" capabilities by marking functions, methods, classes, etc. "experimental." The exact method
    of marking an API experimental differs from language to language. But at a minimum, the documentation for that API
    will indicate that it is experimental. In cases where there are native language-level mechanisms for marking an API
    experimental ([as is the case in Java](https://docs.oracle.com/javase/8/docs/api/java/lang/Deprecated.html)), that
    mechanism is used.

## Day 2 - Day-to-Day Operations Phase

### Scalability/Reliability

  * **Describe how the project increases the size or count of existing API objects.**

      **N/A.** gRPC does not use CRDs or have its own "API objects" in the Kubernetes sense.

  * **Describe how the project defines Service Level Objectives (SLOs) and Service Level Indicators (SLIs).**

    gRPC does not define SLOs for adopters. Instead, our adopters define SLOs for their services (e.g., "99.9% of
    `CreateUser` RPCs must be successful") _using_ Service Level Indicators (SLIs) which _we_ provide.

    Through our [OpenTelemetry integration](https://grpc.io/docs/guides/opentelemetry-metrics/), we export the three
    most critical SLIs: availability, latency, and throughput.

  * **Describe any operations that will increase in time covered by existing SLIs/SLOs.**

      **None.** Because gRPC does not hook into the Kubernetes lifecycle via (e.g.) webhooks, no existing SLIs/SLOs will
      be inflated by deploying an application using gRPC to an cluster.

  * **Describe the increase in resource usage in any components as a result of enabling this project, to include CPU, Memory, Storage, Throughput.**
    
    Because gRPC does not communicate with the Kubernetes control plane, gRPC does not cause increased load on the usual
    components (the Kubernetes API server, etcd, etc.). In fact, relative to a traditional HTTP+JSON REST application,
    applications migrating to gRPC should see _decreases_ in CPU usage, network bandwidth usage, _and_ latency because
    of the byte-efficiency of the payload and simplicity of the "marshalling / unmarshalling" process.

  * **Describe which conditions enabling / using this project would result in resource exhaustion of some node resources (PIDs, sockets, inodes, etc.)**

      gRPC clients use long-lived TCP connections. A misconfigured client (e.g., creating a new Channel for every single
      request instead of re-using a Channel) _could_ exhaust the available file descriptors/sockets on its node. But
      even in this case, this should be no worse than a typical REST application, as the majority of HTTP libraries _do
      not_ reuse TCP connections across requests.

  * **Describe the load testing that has been performed on the project and the results.**

      gRPC's benchmarking framework is described in depth [here](https://grpc.io/docs/guides/benchmarking/). Its
      [dashboard](https://grafana-dot-grpc-testing.appspot.com/?orgId=1) and [code](https://github.com/grpc/test-infra)
      are publicly available. 

  * **Describe the recommended limits of users, requests, system resources, etc. and how they were obtained.**

      **N/A.** As a library, gRPC has no prescribed limits. The limits are the physical limits of the underlying system
      (OS, hardware, network). We are designed to scale to millions of RPCs per second, as demonstrated by the scale of
      usage of our adopters.

  * **Describe which resilience pattern the project uses and how, including the circuit breaker pattern.**

    gRPC is a comprehensive resilience _framework_ providing applications the building blocks they need to produce
    resilient _systems_. The following table outlines the key pieces of resilence functionality gRPC provides.

| Pattern | Specification / Documentation | Purpose & Function |
| :--- | :--- | :--- |
| Client-Side Load Balancing | [Documentation](https://github.com/grpc/grpc/blob/master/doc/load-balancing.md) | The client maintains a list of healthy backends and chooses one for each RPC. This allows the client to actively route around a failed server instance, enabling high availability. |
| Circuit Breaking | [gRFC A32](https://github.com/grpc/proposal/blob/2405799e80581b3a3add27663116db6eada59154/A32-xds-circuit-breaking.md) | When enabled, clients will automatically and immediately fail outgoing requests to any xDS Cluster (roughly equivalent to Kubernetes Service) to which current in-flight requests exceed a configurable threshhold, helping to alleviate overload issues. |
| Deadlines | [Documentation](https://grpc.io/docs/guides/deadlines/) | A client sets a timeout for an RPC. This deadline is propagated to downstream servers. If the deadline is reached, the call is cancelled, freeing resources on all servers. This limits the amount of compute potentially wasted within a system. |
| Retries | [Documentation](https://grpc.io/docs/guides/retry/) | A client channel can be configured with a retry policy (e.g., retry on `UNAVAILABLE`) with exponential backoff. This handles transient network and server errors automatically. |
| Health Checking | [Documentation](https://grpc.io/docs/guides/health-checking/) | Servers can expose the `grpc.health.v1` service. gRPC clients, orchestrators (such as Kubernetes), and service mesh control planes may all use this signal to determine if a backend is healthy and ready to receive traffic. |
| Fault Injection | [gRFC A33](https://github.com/grpc/proposal/blob/master/A33-Fault-Injection.md) | Allows the user to deliberately inject faults into a fraction RPCs in order to observe how the service behaves, improving the robustness of the overall system. |
| Outlier Detection | [gRFC A50](https://github.com/grpc/proposal/blob/master/A50-xds-outlier-detection.md) | Allows the user to configure heuristics that clients can use to decide to stop sending requests to a specific backend because it appears unhealthy. |

Many of these features are implemented as part of gRPC's xDS support and therefore the behavior and API align with those
of Envoy. A full listing of xDS functionality is available
[here](https://grpc.github.io/grpc/core/md_doc_grpc_xds_features.html).

### Observability Requirements

  * **Describe the signals the project is using or producing, including logs, metrics, profiles and traces. Please
    include supported formats, recommended configurations and data storage.**

    gRPC exports a rich, standards-based set of observability [metrics via
    OpenTelemetry](https://grpc.io/docs/guides/opentelemetry-metrics/). We export [the "golden
    signals"](https://sre.google/sre-book/monitoring-distributed-systems/#xref_monitoring_golden-signals) (latency,
    traffic, errors, saturation) in addition to many other more granular metrics in the vendor-neutral [OTLP
    format](https://opentelemetry.io/docs/specs/otlp/). gRPC's OpenTelemetry Metric integration is specified in [gRFC
    A66](https://github.com/grpc/proposal/blob/master/A66-otel-stats.md).

    gRPC produces diagnostic logs related to the transport, event loop, and internal state of the library. These
    integrate with each language's logging framework and the user is free to set the verbosity level of the library.
    These are primarily useful for troubleshooting issues with gRPC or with the network configuration.

  * **Describe how the project captures audit logging.**

    Several gRPC language implementations support audit logging (specified in [gRFC
    A59](https://github.com/grpc/proposal/blob/2405799e80581b3a3add27663116db6eada59154/A59-audit-logging.md)). This may
    be configured either by an xDS control plane or programmatically using the gRPC Authorization API ([gRFC
    A43](https://github.com/grpc/proposal/blob/master/A43-grpc-authorization-api.md)). When enabled, this feature will
    log the full RPC method name and the principal (identity) of the caller.

    A `STDOUT` logger is provided by default so that applications integrating with this feature can retrieve their logs
    from their standard log stream. But an API is provided so that an indvidual application may export these audit logs
    however they desire.

  * **Describe any dashboards the project uses or implements as well as any dashboard requirements.**

    We do not ship or implement dashboards. However, our standardized OpenTelemetry metrics are *designed* to populate
    standard dashboards. Any out-of-the-box Grafana or Datadog dashboard for gRPC will work with our metrics, as we
    follow the OTel semantic conventions.

  * **Describe how the project surfaces project resource requirements for adopters to monitor cloud and infrastructure costs, e.g. FinOps**

    gRPC facilitates FinOps through high-cardinality tagging in its OpenTelemetry metrics. The `grpc.target` and
    `grpc.method` attributes allow operators to attribute resource consumption (CPU, Network I/O) to specific backend
    services or API methods. This granular visibility enables accurate cost allocation in clusters.

  * **Which parameters is the project covering to ensure the health of the application/service and its workloads?**

    gRPC provides a [standardized health-checking protocol](https://grpc.io/docs/guides/health-checking/) to monitor the
    health of individual backends. Since version `1.27`, Kubernetes supports configuring [pod liveness probes using this
    protocol](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/#define-a-grpc-liveness-probe).
    Various xDS control planes also integrate with this protocol.

  * **How can an operator determine if the project is in use by workloads?**

    If a gRPC application has enabled OpenTelemetry integration, then the presence of `grpc.*` OpenTelemetry metrics
    will indicate that gRPC traffic is running within the cluster.

  * **How can someone using this project know that it is working for their instance?**

      * Client-side: The client RPC call returns with a status code of `OK`.
      * Server-side: The server's health check probe  returns `SERVING`.
      * OpenTelemetry: The `grpc.server.call.started` metric increments, and the
        `grpc.server.call.duration{status="OK"}` metric increments.

  * **Describe the SLOs (Service Level Objectives) for this project.**

      **N/A.** We are a library and do not have SLOs. We provide the *SLIs* for our *adopters* to build SLOs, as
      described above.

  * **What are the SLIs (Service Level Indicators) an operator can use to determine the health of the service?**

    We provide the "Golden Signals" as SLIs directly via our OpenTelemetry integration. (latency, traffic, errors,
    saturation) in addition to many other more granular metrics which may be used as SLIs by adopters depending on
    their needs.

| Signal (SLI) | gRPC OpenTelemetry Metric | Description |
| :--- | :--- | :--- |
| **Latency** | `grpc.server.call.duration` (Histogram) | Measures the end-to-end time for the server to handle an RPC. Allows for calculating p50, p90, p99 latencies. |
| **Traffic / Request Rate** | `grpc.server.call.started` (Counter) | A counter of all RPCs started. Used to measure QPS (Queries Per Second). |
| **Error Rate** | `grpc.server.call.duration` (Histogram) | The total count of requests, which can be filtered by the `grpc.status` label. Error Rate = `count{status!="OK"} / count{all}`. |
| **Saturation / Data Volume** | `grpc.server.call.rcvd_total_compressed_message_size` (Histogram) | Measures the size of request messages. Used to track bandwidth consumption and potential memory pressure. |

### Dependencies

  * **Describe the specific running services the project depends on in the cluster.**

      **N/A.** gRPC has no *runtime service dependencies*. It may *optionally* connect to DNS for discovery or an xDS
      control plane for service mesh integration.

  * **Describe the project’s dependency lifecycle policy.**

    gRPC is a multi-language project with somewhat different dependency sets per language.

      * Dependencies by implementation:
          * Cross-Language: While not a hard dependency, most gRPC adopters use `protobuf` as their IDL. We track their
            release cycle closely and continuously test with their newest versions to ensure compatibility.
          * `grpc-java`: Depends on Netty for high-performance I/O.
          * Core (for C++, Python, Ruby, etc.): Depends on foundational libraries like BoringSSL and zlib.
          * `grpc-go`, `grpc-dotnet`:** These are "pure" implementations with minimal external dependencies. This is a
            strategic direction to reduce complexity and security surface area.

    We set a high bar for the quality and maintenance of any external dependency that we incur. Once we have added a
    dependency, updating that dependency to the newest version becomes part of our roughly 6-week release cadence. In
    case of CVEs, a new patch release on a release branch is cut to ensure the safety of user workloads.

  * **How does the project incorporate and consider source composition analysis as part of its development and security
    hygiene? Describe how this source composition analysis (SCA) is tracked.**

    SCA is a standard, automated part of our CI and security hygiene.

    Tools like Dependabot automatically scan our pom.xml, go.mod, package.json, etc., and file issues or pull requests
    when a new vulnerability is found in a dependency.

    Patches are applied and released, either in a regular 6-week minor release or an emergency patch release. A prime
    example is [our response to the HTTP/2 Rapid Reset
    vulnerability](https://github.com/grpc/grpc-go/security/advisories/GHSA-m425-mq94-257g).

  * **Describe how the project implements changes based on source composition analysis (SCA) and the timescale.**

    The timescale for changes is tiered based on severity, governed by our [CVE
    Process](https://github.com/grpc/proposal/blob/2405799e80581b3a3add27663116db6eada59154/P4-grpc-cve-process.md).

### Troubleshooting

  * **How does this project recover if a key component or feature becomes unavailable? e.g Kubernetes API server, etcd, database, leader node, etc.**

    gRPC applications recover via:

      * Retries & Exponential Backoff: Clients automatically retry failed RPCs with `UNAVAILABLE` status using a
        configured backoff strategy (jittered exponential) to prevent thundering herds.
      * Name Resolution Refresh: If the backend set changes (e.g., when leader election happens), the name gRPC name
        resolver will update the set of backend endpoints.
      * Outlier Detection: When enabled, clients will actively eject unhealthy backends from their load balancing pool.

  * **Describe the known failure modes.**

    The failure modes of gRPC applications are explicitly enumerated and defined by the [gRPC standard status
    codes](https://grpc.io/docs/guides/status-codes/).

### Compliance

  * **What steps does the project take to ensure that all third-party code and components have correct and complete
    attribution and license notices?**

    The gRPC project is licensed under the Apache License 2.0, which has explicit requirements for attribution. We
    follow these requirements strictly.

    Our repositories contain a `LICENSE` file and, where required by our dependencies, [a `NOTICE.txt`
    file](https://github.com/grpc/grpc/blob/bdc7762a151abd64b49dc9d430cdf37a51ee7ccc/NOTICE.txt). Since adding a new
    dependency to gRPC has such a high bar, the need for new attribution is determined during the thorough review
    process for the new dependency.

  * **Describe how the project ensures alignment with CNCF
    [recommendations](https://github.com/cncf/foundation/blob/main/policies-guidance/recommendations-for-attribution.md)
    for attribution notices.**

      * **How are notices managed for third-party code incorporated directly into the project's source files?**

        Code incorporated directly (vendored) is placed in the
        [`third_party`](https://github.com/grpc/grpc/tree/master/third_party) directory. Each entry there includes its
        own `LICENSE` file.

      * **How are notices retained for unmodified third-party components included within the project's repository?**

        We do this for key third-party components. One example is the
        [`roots.pem`](https://github.com/grpc/grpc/blob/bdc7762a151abd64b49dc9d430cdf37a51ee7ccc/etc/roots.pem) file
        included in our repository, which is sourced from Mozilla and is licensed under the MPL v2.0. [Our main
        `LICENSE` file explicitly includes the full text of the MPLv2
        license](https://github.com/grpc/grpc/blob/bdc7762a151abd64b49dc9d430cdf37a51ee7ccc/LICENSE#L238) to comply with
        its distribution requirements.

      * **How are notices for all dependencies obtained at build time included in the project's distributed build
        artifacts (e.g. compiled binaries, container images)?**

        For our language-specific implementations, we maintain a `NOTICE.txt` file. This file contains the required
        attribution notices for our build-time dependencies. This `NOTICE.txt` file is distributed alongside our source
        code and build artifacts.

### Security

  * **Security Hygiene**

      * **How is the project executing access control?**
        
        Access to our [source code repositories](https://github.com/grpc) is controlled by GitHub permissions. Per-repo
        admin permissions are granted to [the
        Maintainers](https://github.com/grpc/grpc-community/blob/3711725571c0b94a9f6209e7304c36b4dc9f53da/contributors/maintainers.md)
        for each repo and further permissions are delegated to other contributors by the Maintainers based on merit and
            necessity.

        Branch protection rules prevent direct pushes to `main` / `master` and release branches.

  * **Cloud Native Threat Modeling**

      * **How does the project ensure its security reporting and response team is representative of its community
        diversity (organizational and individual)?**

        Membership on the security team is a function of our [Contributor
        Ladder](https://github.com/grpc/grpc-community/blob/3711725571c0b94a9f6209e7304c36b4dc9f53da/contributor_ladder.md)
        and
        [Governance](https://github.com/grpc/grpc-community/blob/3711725571c0b94a9f6209e7304c36b4dc9f53da/governance.md).
        As contributors advance to become [Core Contributors within the Security
        domain](https://github.com/grpc/grpc-community/blob/3711725571c0b94a9f6209e7304c36b4dc9f53da/contributors/core_contributors.md),
        they take on the responsibilities of the Security team, potentially including being part of the security
        response team. These responsibilities scale as they move from Core Contributor to
        [Maintainer](https://github.com/grpc/grpc-community/blob/3711725571c0b94a9f6209e7304c36b4dc9f53da/contributors/maintainers.md).

        Since the pathway is open for contributors from any organization, the security team should over time reflect the
        diversity of contributors to the gRPC project.

      * **How does the project invite and rotate security reporting team members?**

        See the previous response.
