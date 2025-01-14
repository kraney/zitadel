---
title: Service description for ZITADEL Cloud and ZITADEL Enterprise
sidebar_label: Service description
custom_edit_url: null
--- 

Last updated on November 15, 2023

This annex of the [Framework Agreement](../terms-of-service) describes the services offered by us.

## Services offered

### ZITADEL Cloud

[ZITADEL Cloud](https://zitadel.com) is a fully managed cloud service of the [ZITADEL software](https://github.com/zitadel).

You will benefit from the same software as the open-source project, but we take care of the hosting, maintenance, backup, scaling, and operational security. The cloud service is managed and maintained by the team that also develops the software.

When creating a new instance, you are able to choose a [data location](#data-location).
We follow a single-provider strategy by minimizing the involved [sub-processors](../subprocessors.md) to increase security, compliance, and performance of our services. [Billing](./billing.md) is based on effective usage of our services.

### Enterprise license / self-hosted

The ZITADEL Enterprise license allows you to use the [ZITADEL software](https://github.com/zitadel) on your own data center or private cloud.

You will benefit from the transparency of the open source and the hyper-scalability of the same software that is being used to operate [ZITADEL Cloud](#zitadel-cloud).

#### Benefits over using open source / community license

- [Enterprise supported features](support-services) are only supported under an Enterprise license
- Individual [onboarding support](./support-services#onboarding-support) tailored to your needs and team
- Get access to our support with a [Service Level Agreement](support-services#service-level-agreement) that is tailored to your needs
- Benefit from personal [technical account management](support-services#technical-account-manager) provided by our engineers to help you with architecture, integration, migration, and operational improvements of your setup

#### Benefits over ZITADEL Cloud

You can reduce your supply-chain risks by removing us as sub-processor of personal information about your users.
Support staff will have no access to your infrastructure and will only provide technical support.
Operation and direct maintenance of ZITADEL will be done by you.

You can freely choose the infrastructure and location to host ZITADEL.

## Data location

Data location refers to a region, consisting of one or many countries or territories, where the customer's data is being hosted.

We can not guarantee that during transit the data will only remain within this region. We take measures, as outlined in our [privacy policy](../policies/privacy-policy), to protect your data in transit and in rest.

The following regions will be available when using our cloud service. This list is for informational purposes and will be updated in due course, please refer to our website for all available regions at this time.

- **Global**: All available cloud regions offered by our cloud provider
- **Switzerland**: Exclusively on Swiss region
- **GDPR safe countries**: Hosting location is within any of the EU member states and [Adequate Countries](https://ec.europa.eu/info/law/law-topic/data-protection/international-dimension-data-protection/adequacy-decisions_en) as recognized by the European Commission under the GDPR

## Backup

Our backup strategy executes daily full backups and differential backups on much higher frequency.
In a disaster recovery scenario, our goal is to guarantee a recovery point objective (RPO) of 1h, and a higher but similar recovery time objective (RTO).
Under normal operations, RPO and RTO goals are below 1 minute.

If you you have different requirements we provide you with a flexible approach to backup, restore, and transfer data (f.e. to a self-hosted setup) through our APIs.
Please consult the [migration guides](/docs/guides/migrate/introduction.md) for more information.
