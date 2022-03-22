## KubePlus - Kubernetes SaaS Operator to deliver Helm charts as-a-service

KubePlus is a turn-key solution to transform any containerized application into a SaaS. It takes an application Helm chart and delivers it as-a-service by automating multi-tenancy management and day2 operations such as monitoring, troubleshooting and application upgrades. KubePlus consists of a CRD that enables creating new Kubernetes APIs (CRDs) to realize such services. The new CRDs enable creation of a Helm release per tenant with tenant level isolation, monitoring and consumption tracking.


KubePlus offers following benefits towards deploying a Kubernetes-native application (Helm chart) in SaaS form:
- Seamless support for Namespace-based multi-tenancy where each application instance (Helm release) is created in a separate namespace.
- Application-specific provider and consumer APIs for role based access to the clusters.
- Monitoring and governance of application instances.
- Tracking consumption metrics (cpu, memory, storage and network) at Helm release level. Application providers can use these metrics to define consumption-based chargeback models.

<p align="center">
<img src="./docs/application-stacks-1.png" width="700" height="150" class="center">
</p>


## Overview

The typical requirements in a service-based delivery model of Kubernetes applications are as follows:
- From cluster admin's perspective it is important to isolate different application instances from one another on the cluster.
- Application consumers need a self-service model to provision application instances.
- Application providers need to be able to troubleshoot application instances, monitor them, and track their resource consumption.

KubePlus achieves these goals as follows. KubePlus defines a ```provider API``` to create application-specific ```consumer APIs```.
The ```provider API``` is a KubePlus CRD (Custom Resource Definition) named ``ResourceComposition`` that enables registering an application Helm chart in the cluster by defining a new Kubernetes API (CRD) representing the chart. The new CRD is essentially the ```consumer API``` which the application consumers use to instantiate the registered Helm chart in a self-service manner. Through ``ResourceComposition``application providers can define application-level policies which KubePlus applies when instantiating the registered chart as part of handling the consumer APIs.


<p align="center">
<img src="./docs/provider-consumer.png" width="600" height="200" class="center">
</p>

KubePlus offers following functions:
- Create: Create a Kubernetes-native API to represent an application packaged as a Helm chart.
- Govern: Tenant level policies for isolation and resource utilization per application instance.
- Monitor: Tenant level consumption metrics tracking for cpu, memory, storage, network.
- Troubleshoot: Application-level insights through fine-grained Kubernetes resource relationship graphs.


## Example

To understand the working of KubePlus, let us see how a Wordpress provider can offer a multi-tenant Wordpress service.


### Cluster admin actions

*1. Install KubePlus*

Cluster administrator installs KubePlus on their cluster.

```
$ KUBEPLUS_NS=default
$ helm install kubeplus "https://github.com/cloud-ark/operatorcharts/blob/master/kubeplus-chart-2.0.8.tgz?raw=true" -n $KUBEPLUS_NS
```

*2. Retrieve Provider kubeconfig file*

KubePlus creates provider and consumer kubeconfig files with appropriately scoped
RBAC policies. Cluster admin needs to distribute them to application providers and consumers. KubePlus comes with kubectl plugins to retrieve these files.
The provider kubeconfig file has permissions to register application helm charts under consumer APIs in the cluster. The consumer kubeconfig file has permissions to perform CRUD operations on the registered consumer APIs.

```
$ kubectl retrieve kubeconfig provider $KUBEPLUS_NS > provider.conf
```

### Provider actions

*1. Create consumer API*

The provider team defines the consumer API named ```WordpressService``` using the ```ResourceComposition``` CRD (the provider API). The Wordpress Helm chart that underlies this service is created by the provider team. The spec properties of the ```WordpressService Custom Resource``` are the attributes defined in the Wordpress Helm chart's values.yaml.

As part of registering the consumer API, the provider team can define policies such as the cpu and memory that should be allocated to each Wordpress stack, or the specific worker node on which to deploy a Wordpress stack, etc. KubePlus will apply these policies to the Helm releases when instantiating the underlying Helm chart.

<p align="center">
<img src="./docs/wordpress-service-crd.png" width="650" height="250" class="center">
</p> 

[Here](https://raw.githubusercontent.com/cloud-ark/kubeplus/master/examples/multitenancy/wordpress-mysqlcluster-stack/wordpress-service-composition.yaml) is the ResourceComposition definition for the WordpressService.

*2. Grant permission to the consumer to create instances of WordpressService*

Before consumers can instantiate WordpressService resources, the provider needs to grant permission to the consumer for that resource. KubePlus comes with a kubectl plugin for this purpose:

```
kubectl grantpermission consumer wordpressservices provider.conf $KUBEPLUS_NS
```

*3. Provider team uses kubeplus kubectl plugins to troubleshoot and monitor WordpressService instances*.

Once a consumer has instantiated a WordpressService instance, the provider can
monitor and troubleshoot it using the various kubectl plugins that KubePlus provides.

With the ``kubectl connections`` plugin provider can check whether all Kubernetes resources have been created as expected. The graphical output makes it easy to check the connectivity between different resources.

```
kubectl connections WordpressService tenant1 default -k provider.conf -o png -i Namespace:default,ServiceAccount:default -n label,specproperty,envvariable,annotation 
```
<p align="center">
<img src="./examples/multitenancy/wordpress-mysqlcluster-stack/wp-tenant1.png" class="center">
</p>

Using ```kubectl metrics``` plugin, provider can check cpu, memory, storage, network ingress/egress for a WordpressService instance. The metrics output is available in pretty, json and Prometheus formats.

```
kubectl metrics WordpressService tenant1 default -o pretty -k provider.conf 
```

<p align="center">
<img src="./examples/multitenancy/wordpress-mysqlcluster-stack/wp-tenant1-metrics-pretty.png" class="center">
</p>

### Consumer action

The consumer uses WordpressService Custom Resource (the consumer API) to provision an instance of Wordpress stack. The instances can be created using ``kubectl`` or through a web portal. The portal is part of KubePlus Operator and runs on the cluster. It is accessible through local proxy. Here is consumer portal for WordpressService showing the created ```tenant1``` instance.

<p align="center">
<img src="./examples/multitenancy/wordpress-mysqlcluster-stack/wp-tenant1-consumerui.png" class="center">
</p>


Our [KubePlus SaaS Manager product](https://cloudark.io/kubeplus-saas-manager) offers enterprise-ready control center with embedded Prometheus integration for providers to manage their SaaS across multiple Kubernetes clusters.


## Components

KubePlus consists of an Operator and kubectl plugins.

### 1. KubePlus Operator

The KubePlus Operator consists of a custom controller, a mutating webhook and the helmer module. Here is a brief summary of these components. Details about them are available [here](https://cloud-ark.github.io/kubeplus/docs/html/html/kubeplus-components.html).

<p align="center">
<img src="./docs/crd-for-crds-2.jpg" width="700" height="300" class="center">
</p>

The custom controller handles the ```ResourceComposition```. It is used to:
- Define new CRDs representing Helm charts (consumer APIs).
- Define policies (e.g. cpu/memory limits, node selection, etc.) for service instances.

The mutating webook and helmer modules support the custom controller in delivering the KubePlus experience.


### 2. KubePlus kubectl plugins

KubePlus kubectl plugins enable providers to discover, monitor and troubleshoot application instances. The primary plugin is: ```kubectl connections```. It tracks resource relationships through owner references, labels, annotations, and spec properties. These relationships enable providers to gain fine grained visibility into running application instances through resource relationship graphs. Additional plugins offer the ability to get aggregated consumption metrics (for cpu, memory, storage, network), and logs at the application instance level.


## Try

- Use Kubernetes version <= 1.20 and Helm version 3+. With minikube, you can create a cluster with a specific version like so:
```
    $ minikube start --kubernetes-version=v1.20.0
```

- Install KubePlus Operator.

```
   $ KUBEPLUS_NS=default (or any namespace in which you want to install KubePlus)
   $ helm install kubeplus "https://github.com/cloud-ark/operatorcharts/blob/master/kubeplus-chart-2.0.8.tgz?raw=true" -n $KUBEPLUS_NS
```

- Install KubePlus kubectl plugins (see below)

- Try following examples:
  - [Hello World service](./examples/multitenancy/hello-world/steps.txt)
  - [Wordpress service](./examples/multitenancy/wordpress-mysqlcluster-stack/steps.txt)
  - [Mysql service](./examples/multitenancy/stacks/steps.txt)
  - [MongoDB service](./examples/multitenancy/mongodb-as-a-service/steps.md)
  - [Multiple teams](./examples/multitenancy/team/steps.txt) with applications deployed later

- Debug:
  ```
  - kubectl logs kubeplus $KUBEPLUS_NS -c crd-hook
  - kubectl logs kubeplus $KUBEPLUS_NS -c helmer
  - kubectl logs kubeplus $KUBEPLUS_NS -c platform-operator
  - kubectl logs kubeplus $KUBEPLUS_NS -c webhook-cert-setup
  - kubectl logs kubeplus $KUBEPLUS_NS -c consumerui
  ```

- Cleanup:
  ```
  - helm delete kubeplus -n $KUBEPLUS_NS
  - wget https://github.com/cloud-ark/kubeplus/raw/master/deploy/delete-kubeplus-components.sh
  - ./delete-kubeplus-components.sh
  ```

## Kubectl plugins for discovery, monitoring and troubleshooting

KubePlus kubectl plugins enable discovery, monitoring and troubleshooting of Kubernetes applications. You can install them following these steps:

```
   $ wget https://github.com/cloud-ark/kubeplus/raw/master/kubeplus-kubectl-plugins.tar.gz
   $ gunzip kubeplus-kubectl-plugins.tar.gz
   $ tar -xvf kubeplus-kubectl-plugins.tar
   $ export KUBEPLUS_HOME=`pwd`
   $ export PATH=$KUBEPLUS_HOME/plugins/:$PATH
   $ kubectl kubeplus commands
```

KubePlus's ``kubectl connections`` plugin enables discovering Kubernetes application topologies. You can use it with any Kubernetes resource (built-in resources like Pod, Deployment, or custom resources like MysqlCluster, Jenkins, etc.).
Here is how you can use the ``kubectl connections`` plugin:

```
NAME
        kubectl connections

SYNOPSIS
        kubectl connections <Kind> <Instance> <Namespace> [-k <Absolute path to kubeconfig>] [-o json|png|flat|html] [-i <Kind1:Instance1,Kind1:Instance1>] [-n <label|specproperty|envvariable|annotation>]

DESCRIPTION
        kubectl connections shows how the input resource is connected to other Kubernetes resources through one of the following 
        types of relationships: ownerReference, labels, annotations, spec property.
OPTIONS
        kubectl connections takes following optional flags as input.
        -k <Absolute path to kubeconfig file>
        -o <json|png|flat>
            This flag controls what type of output to generate.
        -i <Kind1:Instance1,Kind2:Instance2>
            This flag defines which Kinds and instances to ignore when traversing the resource graph.
            kubectl connections will not discover the sub-graphs starting at such nodes.
        -n <label|specproperty|envvariable|annotation>
            This flag defines the relationship types whose details should not be displayed in the graphical output (png).
            You can specify multiple values as comma separated list.
```

[Here](./examples/graphs) are some resource relationship graphs generated using the connections plugin.


## CNCF Landscape

KubePlus is part of CNCF landscape's [Application Definition section](https://landscape.cncf.io/card-mode?category=application-definition-image-build&grouping=category).


## Operator Maturity Model

As enterprise teams build their custom Kubernetes platforms using community or in house developed Operators, they need a set of guidelines for Operator readiness in multi-Operator and multi-tenant environments. We have developed the [Operator Maturity Model](https://github.com/cloud-ark/kubeplus/blob/master/Guidelines.md) for this purpose. Operator developers are using this model today to ensure that their Operator is a good citizen of the multi-Operator world and ready to serve multi-tenant workloads. It is also being used by Kubernetes cluster administrators for curating community Operators towards building their custom platforms.


## Presentations

1. [DevOps.com Webinar: Deliver your Kubernetes Applications as-a-Service](https://webinars.devops.com/deliver-your-kubernetes-applications-as-a-service)

2. [Being a good citizen of the Multi-Operator world, Kubecon NA 2020](https://www.youtube.com/watch?v=NEGs0GMJbCw&t=2s)

3. [Operators and Helm: It takes two to Tango, Helm Summit 2019](https://youtu.be/F_Dgz1V5Q2g)

4. [KubePlus presentation at community meetings (CNCF sig-app-delivery, Kubernetes sig-apps, Helm)](https://github.com/cloud-ark/kubeplus/blob/master/KubePlus-presentation.pdf)


## Contact

Submit issues on this repository or reach out to our team on [Slack](https://join.slack.com/t/cloudark/shared_invite/zt-2yp5o32u-sOq4ub21TvO_kYgY9ZfFfw).
