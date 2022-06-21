from typing import Generator, Literal, Tuple, Union

from kat.harness import Query

from abstract_tests import AmbassadorTest, HTTP, ServiceType, RLSGRPC, Node
from tests.selfsigned import TLSCerts

from ambassador import Config


class RateLimitV0Test(AmbassadorTest):
    # debug = True
    target: ServiceType
    rls: ServiceType

    def init(self):
        self.target = HTTP()
        self.rls = RLSGRPC()

    def config(self) -> Generator[Union[str, Tuple[Node, str]], None, None]:
        # Use self.target here, because we want this mapping to be annotated
        # on the service, not the Ambassador.
        # ambassador_id: [ {self.with_tracing.ambassador_id}, {self.no_tracing.ambassador_id} ]
        yield self.target, self.format("""
---
apiVersion: getambassador.io/v3alpha1
kind: Mapping
name:  ratelimit_target_mapping
hostname: "*"
prefix: /target/
service: {self.target.path.fqdn}
labels:
  ambassador:
    - request_label_group:
      - request_headers:
          key: x-ambassador-test-allow
          header_name: "x-ambassador-test-allow"
          omit_if_not_present: true
      - request_headers:
          key: x-ambassador-test-headers-append
          header_name: "x-ambassador-test-headers-append"
          omit_if_not_present: true
---
apiVersion: getambassador.io/v3alpha1
kind: Mapping
name:  ratelimit_label_mapping
hostname: "*"
prefix: /labels/
service: {self.target.path.fqdn}
labels:
  ambassador:
    - host_and_user:
      - request_headers:
          key: custom-label
          header_name: ":authority"
          omit_if_not_present: true
      - request_headers:
          key: user
          header_name: "x-user"
          omit_if_not_present: true

    - omg_header:
      - request_headers:
          key: custom-label
          header_name: "x-omg"
          default: "OMFG!"
""")

        # For self.with_tracing, we want to configure the TracingService.
        yield self, self.format("""
---
apiVersion: getambassador.io/v3alpha1
kind: RateLimitService
name: {self.rls.path.k8s}
service: "{self.rls.path.fqdn}"
timeout_ms: 500
protocol_version: "v3"
""")

    def queries(self):
        # Speak through each Ambassador to the traced service...
        # yield Query(self.with_tracing.url("target/"))
        # yield Query(self.no_tracing.url("target/"))

        # [0]
        # No matching headers, won't even go through ratelimit-service filter
        yield Query(self.url("target/"))

        # [1]
        # Header instructing dummy ratelimit-service to allow request
        yield Query(self.url("target/"), expected=200, headers={
            'x-ambassador-test-allow': 'true',
            'x-ambassador-test-headers-append': 'no header',
        })

        # [2]
        # Header instructing dummy ratelimit-service to reject request with
        # a custom response body
        yield Query(self.url("target/"), expected=429, headers={
            'x-ambassador-test-allow': 'over my dead body',
            'x-ambassador-test-headers-append': 'Hello=Foo; Hi=Baz',
        })

    def check(self):
        # [2] Verifies the 429 response and the proper content-type.
        # The kat-server gRPC ratelimit implementation explicitly overrides
        # the content-type to json, because the response is in fact json
        # and we need to verify that this override is possible/correct.
        assert self.results[2].headers["Hello"] == [ "Foo" ]
        assert self.results[2].headers["Hi"] == [ "Baz" ]
        assert self.results[2].headers["Content-Type"] == [ "application/json" ]
        assert self.results[2].headers["X-Grpc-Service-Protocol-Version"] == [ "v3" ]

class RateLimitV1Test(AmbassadorTest):
    # debug = True
    target: ServiceType

    def init(self):
        self.target = HTTP()
        self.rls = RLSGRPC()

    def config(self) -> Generator[Union[str, Tuple[Node, str]], None, None]:
        # Use self.target here, because we want this mapping to be annotated
        # on the service, not the Ambassador.
        yield self.target, self.format("""
---
apiVersion: getambassador.io/v3alpha1
kind: Mapping
name:  ratelimit_target_mapping
hostname: "*"
prefix: /target/
service: {self.target.path.fqdn}
labels:
  ambassador:
    - request_label_group:
      - request_headers:
          key: x-ambassador-test-allow
          header_name: "x-ambassador-test-allow"
          omit_if_not_present: true
      - request_headers:
          key: x-ambassador-test-headers-append
          header_name: "x-ambassador-test-headers-append"
          omit_if_not_present: true
""")

        yield self, self.format("""
---
apiVersion: getambassador.io/v3alpha1
kind: RateLimitService
name: {self.rls.path.k8s}
service: "{self.rls.path.fqdn}"
timeout_ms: 500
protocol_version: "v3"
""")

    def queries(self):
        # [0]
        # No matching headers, won't even go through ratelimit-service filter
        yield Query(self.url("target/"))

        # [1]
        # Header instructing dummy ratelimit-service to allow request
        yield Query(self.url("target/"), expected=200, headers={
            'x-ambassador-test-allow': 'true',
            'x-ambassador-test-headers-append': 'no header',
        })

        # [2]
        # Header instructing dummy ratelimit-service to reject request
        yield Query(self.url("target/"), expected=429, headers={
            'x-ambassador-test-allow': 'over my dead body',
            'x-ambassador-test-headers-append': 'Hello=Foo; Hi=Baz',
        })

    def check(self):
        # [2] Verifies the 429 response and the proper content-type.
        # The kat-server gRPC ratelimit implementation explicitly overrides
        # the content-type to json, because the response is in fact json
        # and we need to verify that this override is possible/correct.
        assert self.results[2].headers["Hello"] == [ "Foo" ]
        assert self.results[2].headers["Hi"] == [ "Baz" ]
        assert self.results[2].headers["Content-Type"] == [ "application/json" ]
        assert self.results[2].headers["X-Grpc-Service-Protocol-Version"] == [ "v3" ]

class RateLimitV1WithTLSTest(AmbassadorTest):
    # debug = True
    target: ServiceType

    def init(self):
        self.target = HTTP()
        self.rls = RLSGRPC()

    def manifests(self) -> str:
        return f"""
---
apiVersion: v1
data:
  tls.crt: {TLSCerts["ratelimit.datawire.io"].k8s_crt}
  tls.key: {TLSCerts["ratelimit.datawire.io"].k8s_key}
kind: Secret
metadata:
  name: ratelimit-tls-secret
type: kubernetes.io/tls
""" + super().manifests()

    def config(self) -> Generator[Union[str, Tuple[Node, str]], None, None]:
        # Use self.target here, because we want this mapping to be annotated
        # on the service, not the Ambassador.
        yield self.target, self.format("""
---
apiVersion: getambassador.io/v3alpha1
kind: TLSContext
name: ratelimit-tls-context
secret: ratelimit-tls-secret
alpn_protocols: h2
---
apiVersion: getambassador.io/v3alpha1
kind: Mapping
name:  ratelimit_target_mapping
hostname: "*"
prefix: /target/
service: {self.target.path.fqdn}
labels:
  ambassador:
    - request_label_group:
      - request_headers:
          key: x-ambassador-test-allow
          header_name: "x-ambassador-test-allow"
          omit_if_not_present: true
      - request_headers:
          key: x-ambassador-test-headers-append
          header_name: "x-ambassador-test-headers-append"
          omit_if_not_present: true
""")

        yield self, self.format("""
---
apiVersion: getambassador.io/v3alpha1
kind: RateLimitService
name: {self.rls.path.k8s}
service: "{self.rls.path.fqdn}"
timeout_ms: 500
tls: ratelimit-tls-context
protocol_version: "v3"
""")

    def queries(self):
        # No matching headers, won't even go through ratelimit-service filter
        yield Query(self.url("target/"))

        # Header instructing dummy ratelimit-service to allow request
        yield Query(self.url("target/"), expected=200, headers={
            'x-ambassador-test-allow': 'true'
        })

        # Header instructing dummy ratelimit-service to reject request
        yield Query(self.url("target/"), expected=429, headers={
            'x-ambassador-test-allow': 'nope',
            'x-ambassador-test-headers-append': 'Hello=Foo; Hi=Baz'
        })

    def check(self):
        # [2] Verifies the 429 response and the proper content-type.
        # The kat-server gRPC ratelimit implementation explicitly overrides
        # the content-type to json, because the response is in fact json
        # and we need to verify that this override is possible/correct.
        assert self.results[2].headers["Hello"] == [ "Foo" ]
        assert self.results[2].headers["Hi"] == [ "Baz" ]
        assert self.results[2].headers["Content-Type"] == [ "application/json" ]
        assert self.results[2].headers["X-Grpc-Service-Protocol-Version"] == [ "v3" ]


class RateLimitVerTest(AmbassadorTest):
    # debug = True
    target: ServiceType
    specified_protocol_version: Literal['v2', 'v3', 'default']
    expected_protocol_version: Literal['v3']
    rls: ServiceType

    @classmethod
    def variants(cls) -> Generator[Node, None, None]:
        for protocol_version in ['v2', 'v3', 'default']:
            yield cls(protocol_version, name="{self.specified_protocol_version}")

    def init(self, protocol_version: Literal['v2', 'v3', 'default']):
        self.target = HTTP()
        self.specified_protocol_version = protocol_version
        self.expected_protocol_version = "v3"
        self.rls = RLSGRPC(protocol_version=self.expected_protocol_version)

    def config(self) -> Generator[Union[str, Tuple[Node, str]], None, None]:
        # Use self.target here, because we want this mapping to be annotated
        # on the service, not the Ambassador.
        yield self.target, self.format("""
---
apiVersion: getambassador.io/v3alpha1
kind: Mapping
name:  ratelimit_target_mapping
hostname: "*"
prefix: /target/
service: {self.target.path.fqdn}
labels:
  ambassador:
    - request_label_group:
      - request_headers:
          key: x-ambassador-test-allow
          header_name: "x-ambassador-test-allow"
          omit_if_not_present: true
      - request_headers:
          key: x-ambassador-test-headers-append
          header_name: "x-ambassador-test-headers-append"
          omit_if_not_present: true
""")

        yield self, self.format("""
---
apiVersion: getambassador.io/v3alpha1
kind: RateLimitService
name: {self.rls.path.k8s}
service: "{self.rls.path.fqdn}"
timeout_ms: 500
""") + ("" if self.specified_protocol_version == "default" else f"protocol_version: '{self.specified_protocol_version}'")

    def queries(self):
        # [0]
        # No matching headers, won't even go through ratelimit-service filter
        yield Query(self.url("target/"))

        # [1]
        # Header instructing dummy ratelimit-service to allow request
        yield Query(self.url("target/"), expected=200, headers={
            'x-ambassador-test-allow': 'true',
            'x-ambassador-test-headers-append': 'no header',
        })

        # [2]
        # Header instructing dummy ratelimit-service to reject request
        # if v2 or default then RateLimit Filter is dropped and request will get through
        exp_query2 = 200 if self.specified_protocol_version in ["v2", "default"] else 429
        yield Query(self.url("target/"), expected=exp_query2, headers={
            'x-ambassador-test-allow': 'over my dead body',
            'x-ambassador-test-headers-append': 'Hello=Foo; Hi=Baz',
        })

    def check(self):
        if self.specified_protocol_version in ["v2", "default"]:
            # all queries should succeed because the rate-limit filter was dropped, due to bad protocol
            assert self.results[0].status == 200
            assert self.results[1].status == 200
            assert self.results[2].status == 200
        else:
            # [2] Verifies the 429 response and the proper content-type.
            # The kat-server gRPC ratelimit implementation explicitly overrides
            # the content-type to json, because the response is in fact json
            # and we need to verify that this override is possible/correct.
            assert self.results[2].headers["Hello"] == [ "Foo" ]
            assert self.results[2].headers["Hi"] == [ "Baz" ]
            assert self.results[2].headers["Content-Type"] == [ "application/json" ]
            assert self.results[2].headers["X-Grpc-Service-Protocol-Version"] == [ self.expected_protocol_version ]
