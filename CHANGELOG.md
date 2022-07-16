# Release (2022-07-16)

:warning: DEPRECATED :warning:

This module is deprecated and replaced with [github.com/patrickcping/pingone-go-sdk-v2](https://github.com/patrickcping/pingone-go-sdk-v2/)

Conversion from this module to `github.com/patrickcping/pingone-go-sdk-v2` [Release 2022-07-16](https://github.com/patrickcping/pingone-go-sdk-v2/releases/tag/release-2022-07-16) is direct, there is no functional change.  The deprecation is to rebase the version history to better reflect the general stability of the module.

# Release (2022-07-15)

* `github.com/patrickcping/pingone-go/management` : [v1.11.0](./management/CHANGELOG.md)
    * **Feature:** Expansion of P1Error model
    * **Feature:** Expansion of PasswordPolicy model + submodels
    * **Feature:** Cleanup of the Password Policy API
    * **Enhancement:** Change popID parameter name to populationID for readability
    * **Feature:** Cleanup of the Sign-on Policies API
    * **Feature:** Cleanup of the Application Sign-on policies assignment API
    * **Enhancement:** Added common error codes
* `github.com/patrickcping/pingone-go/mfa` : [v1.11.0](./mfa/CHANGELOG.md)
    * **Feature:** Expansion of P1Error model
    * **Enhancement:** Added common error codes
* `github.com/patrickcping/pingone-go/risk` : [v1.11.0](./risk/CHANGELOG.md)
    * **Feature:** Expansion of P1Error model
    * **Enhancement:** Added common error codes

# Release (2022-07-13)

**Redrafted release process**

This GO module will now follow a release naming convention based on date, and the submodules will be based on GO version naming conventions.

The following repository releases are deprecated:

* [v1.9.1](https://github.com/patrickcping/pingone-go/releases/tag/v1.9.1)
* [v1.8.1](https://github.com/patrickcping/pingone-go/releases/tag/v1.8.1)
* [v1.8.0](https://github.com/patrickcping/pingone-go/releases/tag/v1.8.0)

## Module Versions

As modules were released under the previous naming convention, the following provides the starting release number for each

* `github.com/patrickcping/pingone-go/management` : [v1.10.0](./management/CHANGELOG.md)
* `github.com/patrickcping/pingone-go/mfa` : [v1.10.0](./mfa/CHANGELOG.md)
* `github.com/patrickcping/pingone-go/risk` : [v1.10.0](./risk/CHANGELOG.md)