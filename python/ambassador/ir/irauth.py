from typing import TYPE_CHECKING, Literal, Optional
from typing import cast as typecast

from ..config import Config
from ..resource import Resource
from ..utils import RichStatus
from .ircluster import IRCluster
from .irfilter import IRFilter

if TYPE_CHECKING:
    from .ir import IR  # pragma: no cover


class IRAuth(IRFilter):
    cluster: Optional[IRCluster]
    protocol_version: Literal["v2", "v3"]

    def __init__(
        self,
        ir: "IR",
        aconf: Config,
        rkey: str = "ir.auth",
        kind: str = "IRAuth",
        name: str = "extauth",
        namespace: Optional[str] = None,
        type: Optional[str] = "decoder",
        **kwargs,
    ) -> None:

        super().__init__(
            ir=ir,
            aconf=aconf,
            rkey=rkey,
            kind=kind,
            name=name,
            namespace=namespace,
            cluster=None,
            timeout_ms=None,
            connect_timeout_ms=3000,
            path_prefix=None,
            api_version=None,
            allowed_headers=[],
            allowed_request_headers=[],
            allowed_authorization_headers=[],
            hosts={},
            type=type,
            **kwargs,
        )

    def setup(self, ir: "IR", aconf: Config) -> bool:
        module_info = aconf.get_module("authentication")

        if module_info:
            self._load_auth(module_info, ir)

        config_info = aconf.get_config("auth_configs")

        if config_info:
            for config in config_info.values():
                self._load_auth(config, ir)

        if not self.hosts:
            self.logger.debug("IRAuth: no AuthServices, going inactive")
            return False

        self.logger.debug("IRAuth: going active")

        return True

    def add_mappings(self, ir: "IR", aconf: Config):
        cluster_hosts = self.get("hosts", {"127.0.0.1:5000": (100, None, "-internal-")})

        self.cluster = None
        cluster_good = False

        for service, params in cluster_hosts.items():
            weight, grpc, ctx_name, location = params

            self.logger.debug(
                "IRAuth: svc %s, weight %s, grpc %s, ctx_name %s, location %s"
                % (service, weight, grpc, ctx_name, location)
            )

            cluster = IRCluster(
                ir=ir,
                aconf=aconf,
                parent_ir_resource=self,
                location=location,
                service=service,
                host_rewrite=self.get("host_rewrite", False),
                ctx_name=ctx_name,
                grpc=grpc,
                marker="extauth",
                stats_name=self.get("stats_name", None),
                circuit_breakers=self.get("circuit_breakers", None),
            )

            cluster.referenced_by(self)

            cluster_good = True

            if self.cluster:
                if not self.cluster.merge(cluster):
                    self.post_error(
                        RichStatus.fromError(
                            "auth canary %s can only change service!" % cluster.name
                        )
                    )
                    cluster_good = False
            else:
                self.cluster = cluster

        if cluster_good:
            ir.add_cluster(typecast(IRCluster, self.cluster))
            self.referenced_by(typecast(IRCluster, self.cluster))

    def _load_auth(self, module: Resource, ir: "IR"):
        self.namespace = module.get("namespace", self.namespace)
        if self.location == "--internal--":
            self.sourced_by(module)

        for key in ["path_prefix", "timeout_ms", "cluster", "allow_request_body", "proto"]:
            value = module.get(key, None)

            if value:
                previous = self.get(key, None)

                if previous and (previous != value):
                    # Don't use self.post_error() here, since we need to explicitly override the
                    # resource. And don't use self.ir.post_error, since our module isn't an IRResource.
                    self.ir.aconf.post_error(
                        "AuthService cannot support multiple %s values; using %s" % (key, previous),
                        resource=module,
                    )
                else:
                    self[key] = value

            self.referenced_by(module)

        if module.get("add_linkerd_headers"):
            self["add_linkerd_headers"] = module.get("add_linkerd_headers")
        else:
            add_linkerd_headers = module.get("add_linkerd_headers", None)
            if add_linkerd_headers is None:
                self["add_linkerd_headers"] = ir.ambassador_module.get("add_linkerd_headers", False)

        if module.get("circuit_breakers", None):
            self["circuit_breakers"] = module.get("circuit_breakers")
        else:
            cb = ir.ambassador_module.get("circuit_breakers")

            if cb:
                self["circuit_breakers"] = cb

        self["allow_request_body"] = module.get("allow_request_body", False)
        self["include_body"] = module.get("include_body", None)
        self["api_version"] = module.get("apiVersion", None)
        self["proto"] = module.get("proto", "http")
        self["timeout_ms"] = module.get("timeout_ms", 5000)
        self["connect_timeout_ms"] = module.get("connect_timeout_ms", 3000)
        self["cluster_idle_timeout_ms"] = module.get("cluster_idle_timeout_ms", None)
        self["cluster_max_connection_lifetime_ms"] = module.get(
            "cluster_max_connection_lifetime_ms", None
        )
        self["add_auth_headers"] = module.get("add_auth_headers", {})
        self.__to_header_list("allowed_headers", module)
        self.__to_header_list("allowed_request_headers", module)
        self.__to_header_list("allowed_authorization_headers", module)

        if self["proto"] not in ["grpc", "http"]:
            self.post_error(
                f'AuthService: proto_version {self["proto"]} is unsupported, proto must be "grpc" or "http"'
            )

        self.protocol_version = module.get("protocol_version", "v2")
        if self["proto"] == "grpc" and self.protocol_version not in ["v3"]:
            self.post_error(
                f'AuthService: protocol_version {self.protocol_version} is unsupported, protocol_version must be "v3"'
            )

        self["stats_name"] = module.get("stats_name", None)

        status_on_error = module.get("status_on_error", None)
        if status_on_error:
            self["status_on_error"] = status_on_error

        failure_mode_allow = module.get("failure_mode_allow", None)
        if failure_mode_allow:
            self["failure_mode_allow"] = failure_mode_allow

        # Required fields check.
        if self["api_version"] == None:
            self.post_error(RichStatus.fromError("AuthService config requires apiVersion field"))

        if self["proto"] == None:
            self.post_error(RichStatus.fromError("AuthService requires proto field."))

        if self.get("include_body") and self.get("allow_request_body"):
            self.post_error("AuthService ignoring allow_request_body since include_body is present")
            del self["allow_request_body"]

        auth_service = module.get("auth_service", None)
        weight = 100  # Can't support arbitrary weights right now.

        if auth_service:
            is_grpc = True if self["proto"] == "grpc" else False
            self.hosts[auth_service] = (weight, is_grpc, module.get("tls", None), module.location)

    def __to_header_list(self, list_name, module):
        headers = module.get(list_name, None)

        if headers:
            allowed_headers = self.get(list_name, [])

            for hdr in sorted(headers):
                if hdr.lower() not in allowed_headers:
                    allowed_headers.append(hdr.lower())

            self[list_name] = allowed_headers
